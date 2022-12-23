package main

import (
	"context"
	"encoding/pem"
	"errors"
	"expvar"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ardanlabs/conf/v3"
	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	"github.com/ardanlabs/liarsdice/app/services/engine/handlers"
	scbank "github.com/ardanlabs/liarsdice/business/contract/go/bank"
	"github.com/ardanlabs/liarsdice/business/core/bank"
	"github.com/ardanlabs/liarsdice/business/web/auth"
	"github.com/ardanlabs/liarsdice/foundation/events"
	"github.com/ardanlabs/liarsdice/foundation/logger"
	"github.com/ardanlabs/liarsdice/foundation/vault"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"
)

/*
	-- Game Engine
	Once Liar is called, the status needs to share the dice for all players.
	Add in-game chat support.
	Add a Drain function to the smart contract.
	Add an account fix function to adjust balances.
	Have engine sign all transactions to the smart contract.
	Add multi-table with max of 5 players.

	-- Browser UI
	Use Phaser as a new UI
*/

// build is the git version of this program. It is set using build flags in the makefile.
var build = "develop"

func main() {

	// Construct the application logger.
	log, err := logger.New("ENGINE")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer log.Sync()

	// Perform the startup and shutdown sequence.
	if err := run(log); err != nil {
		log.Errorw("startup", "ERROR", err)
		log.Sync()
		os.Exit(1)
	}
}

func run(log *zap.SugaredLogger) error {

	// =========================================================================
	// GOMAXPROCS

	// Want to see what maxprocs reports.
	opt := maxprocs.Logger(log.Infof)

	// Set the correct number of threads for the service
	// based on what is available either by the machine or quotas.
	if _, err := maxprocs.Set(opt); err != nil {
		return fmt.Errorf("maxprocs: %w", err)
	}
	log.Infow("startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	// =========================================================================
	// Configuration

	cfg := struct {
		conf.Version
		Web struct {
			ReadTimeout     time.Duration `conf:"default:5s"`
			WriteTimeout    time.Duration `conf:"default:10s"`
			IdleTimeout     time.Duration `conf:"default:120s"`
			ShutdownTimeout time.Duration `conf:"default:20s"`
			APIHost         string        `conf:"default:0.0.0.0:3000"`
			DebugHost       string        `conf:"default:0.0.0.0:4000"`
		}
		Vault struct {
			Address   string `conf:"default:http://vault-service.liars-system.svc.cluster.local:8200"`
			MountPath string `conf:"default:secret"`
			Token     string `conf:"default:mytoken,mask"`
		}
		Auth struct {
			ActiveKID string `conf:"default:54bb2165-71e1-41a6-af3e-7da4a0e1e2c1"`
		}
		Game struct {
			ContractID     string        `conf:"default:0x0"`
			AnteUSD        float64       `conf:"default:5"`
			ConnectTimeout time.Duration `conf:"default:60s"`
		}
		Bank struct {
			KeyID            string        `conf:"default:6327a38415c53ffb36c11db55ea74cc9cb4976fd"`
			Network          string        `conf:"default:http://geth-service.liars-system.svc.cluster.local:8545"`
			Timeout          time.Duration `conf:"default:10s"`
			CoinMarketCapKey string        `conf:"default:a8cd12fb-d056-423f-877b-659046af0aa5"`
		}
	}{
		Version: conf.Version{
			Build: build,
			Desc:  "copyright information here",
		},
	}

	const prefix = ""
	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}
		return fmt.Errorf("parsing config: %w", err)
	}

	// =========================================================================
	// App Starting

	log.Infow("starting service", "version", build)
	defer log.Infow("shutdown complete")

	out, err := conf.String(&cfg)
	if err != nil {
		return fmt.Errorf("generating config for output: %w", err)
	}
	log.Infow("startup", "config", out)

	expvar.NewString("build").Set(build)

	// =========================================================================
	// Initialize authentication support

	log.Infow("startup", "status", "initializing authentication support")

	vaultClient, err := vault.New(vault.Config{
		Address:   cfg.Vault.Address,
		Token:     cfg.Vault.Token,
		MountPath: cfg.Vault.MountPath,
	})
	if err != nil {
		return fmt.Errorf("constructing vaultConfig: %w", err)
	}

	authCfg := auth.Config{
		Log:       log,
		KeyLookup: vaultClient,
	}

	authClient, err := auth.New(authCfg)
	if err != nil {
		return fmt.Errorf("constructing authClient: %w", err)
	}

	// =========================================================================
	// Create the currency converter and bankClient needed for the game

	if cfg.Game.ContractID == "0x0" {
		return errors.New("smart contract id not provided")
	}

	converter, err := currency.NewConverter(scbank.BankMetaData.ABI, cfg.Bank.CoinMarketCapKey)
	if err != nil {
		log.Infow("unable to create converter, using default", "ERROR", err)
		converter = currency.NewDefaultConverter(scbank.BankMetaData.ABI)
	}

	oneETHToUSD, oneUSDToETH := converter.Values()
	log.Infow("currency values", "oneETHToUSD", oneETHToUSD, "oneUSDToETH", oneUSDToETH)

	evts := events.New()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	backend, err := ethereum.CreateDialedBackend(ctx, cfg.Bank.Network)
	if err != nil {
		return fmt.Errorf("ethereum backend: %w", err)
	}
	defer backend.Close()

	privateKey, err := vaultClient.PrivateKeyPEM(cfg.Bank.KeyID)
	if err != nil {
		return fmt.Errorf("capture private key: %w", err)
	}

	block, _ := pem.Decode([]byte(privateKey))
	ecdsaKey, err := crypto.ToECDSA(block.Bytes)
	if err != nil {
		return fmt.Errorf("error converting PEM to ECDSA: %w", err)
	}

	bankClient, err := bank.New(ctx, log, backend, ecdsaKey, common.HexToAddress(cfg.Game.ContractID))
	if err != nil {
		return fmt.Errorf("connecting to bankClient: %w", err)
	}

	// =========================================================================
	// Start Debug Service

	log.Infow("startup", "status", "debug v1 router started", "host", cfg.Web.DebugHost)

	// The Debug function returns a mux to listen and serve on for all the debug
	// related endpoints. This includes the standard library endpoints.

	// Construct the mux for the debug calls.
	debugMux := handlers.DebugMux(build, log)

	// Start the service listening for debug requests.
	// Not concerned with shutting this down with load shedding.
	go func() {
		if err := http.ListenAndServe(cfg.Web.DebugHost, debugMux); err != nil {
			log.Errorw("shutdown", "status", "debug v1 router closed", "host", cfg.Web.DebugHost, "ERROR", err)
		}
	}()

	// =========================================================================
	// Start API Service

	log.Infow("startup", "status", "initializing V1 API support")

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	// Construct the mux for the API calls.
	apiMux := handlers.APIMux(handlers.APIMuxConfig{
		Shutdown:       shutdown,
		Log:            log,
		Auth:           authClient,
		Converter:      converter,
		Bank:           bankClient,
		Evts:           evts,
		AnteUSD:        cfg.Game.AnteUSD,
		ActiveKID:      cfg.Auth.ActiveKID,
		BankTimeout:    cfg.Bank.Timeout,
		ConnectTimeout: cfg.Game.ConnectTimeout,
	}, handlers.WithCORS("*"))

	// Construct a server to service the requests against the mux.
	api := http.Server{
		Addr:         cfg.Web.APIHost,
		Handler:      apiMux,
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
		IdleTimeout:  cfg.Web.IdleTimeout,
		ErrorLog:     zap.NewStdLog(log.Desugar()),
	}

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	// Start the service listening for api requests.
	go func() {
		log.Infow("startup", "status", "api router started", "host", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	// =========================================================================
	// Shutdown

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		log.Infow("shutdown", "status", "shutdown started", "signal", sig)
		defer log.Infow("shutdown", "status", "shutdown complete", "signal", sig)

		// Release any web sockets that are currently active.
		log.Infow("shutdown", "status", "shutdown web socket channels")
		evts.Shutdown()

		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(context.Background(), cfg.Web.ShutdownTimeout)
		defer cancel()

		// Asking listener to shut down and shed load.
		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}
