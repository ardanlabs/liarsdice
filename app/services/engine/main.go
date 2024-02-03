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

	"github.com/ardanlabs/conf/v3"
	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	"github.com/ardanlabs/liarsdice/app/services/engine/v1/build/all"
	scbank "github.com/ardanlabs/liarsdice/business/contract/go/bank"
	"github.com/ardanlabs/liarsdice/business/core/bank"
	"github.com/ardanlabs/liarsdice/business/data/sqldb"
	"github.com/ardanlabs/liarsdice/business/web/v1/auth"
	"github.com/ardanlabs/liarsdice/business/web/v1/debug"
	"github.com/ardanlabs/liarsdice/business/web/v1/mux"
	"github.com/ardanlabs/liarsdice/foundation/keystore"
	"github.com/ardanlabs/liarsdice/foundation/logger"
	"github.com/ardanlabs/liarsdice/foundation/web"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

/*
	-- Game Engine
	Deprecated: please use simulated.Backend from package : github.com/ethereum/go-ethereum/ethclient/simulated instead.
	Fixing logging with GameID
	Fix CheckDatabase function to test more fields
	Add flags to migrate command
	Finish dbGame package to handle all tables on to functions.

	Now that a player can be in multiple games, we need to better control their available balance.
	Once Liar is called, the status needs to share the dice for all players.
	Add in-game chat support.
	Add a Drain function to the smart contract.
	Add an account fix function to adjust balances.
	Have engine sign all transactions to the smart contract.

	-- Browser UI
	Use Phaser as a new UI
*/

// build is the git version of this program. It is set using build flags in the makefile.
var build = "develop"

func main() {
	var log *logger.Logger

	events := logger.Events{
		Error: func(ctx context.Context, r logger.Record) {
			log.Info(ctx, "******* SEND ALERT ******")
		},
	}

	traceIDFunc := func(ctx context.Context) string {
		return web.GetTraceID(ctx)
	}

	log = logger.NewWithEvents(os.Stdout, logger.LevelInfo, "SALES-API", traceIDFunc, events)

	// -------------------------------------------------------------------------

	ctx := context.Background()

	if err := run(ctx, log); err != nil {
		log.Error(ctx, "startup", "msg", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, log *logger.Logger) error {

	// -------------------------------------------------------------------------
	// GOMAXPROCS

	log.Info(ctx, "startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	// -------------------------------------------------------------------------
	// Configuration

	cfg := struct {
		conf.Version
		Web struct {
			ReadTimeout        time.Duration `conf:"default:5s"`
			WriteTimeout       time.Duration `conf:"default:10s"`
			IdleTimeout        time.Duration `conf:"default:120s"`
			ShutdownTimeout    time.Duration `conf:"default:20s"`
			APIHost            string        `conf:"default:0.0.0.0:3000"`
			DebugHost          string        `conf:"default:0.0.0.0:4000"`
			CORSAllowedOrigins []string      `conf:"default:*"`
		}
		Vault struct {
			Address   string `conf:"default:http://vault-service.liars-system.svc.cluster.local:8200"`
			MountPath string `conf:"default:secret"`
			Token     string `conf:"default:mytoken,mask"`
		}
		Auth struct {
			KeysFolder string `conf:"default:zarf/keys/"`
			ActiveKID  string `conf:"default:54bb2165-71e1-41a6-af3e-7da4a0e1e2c1"`
		}
		DB struct {
			User         string `conf:"default:postgres"`
			Password     string `conf:"default:postgres,mask"`
			HostPort     string `conf:"default:database-service.liars-system.svc.cluster.local"`
			Name         string `conf:"default:postgres"`
			MaxIdleConns int    `conf:"default:2"`
			MaxOpenConns int    `conf:"default:0"`
			DisableTLS   bool   `conf:"default:true"`
		}
		Game struct {
			ContractID     string        `conf:"default:0x0"`
			AnteUSD        float64       `conf:"default:5"`
			ConnectTimeout time.Duration `conf:"default:60s"`
		}
		Bank struct {
			KeysFolder       string        `conf:"default:zarf/ethereum/keystore/"`
			PassPhrase       string        `conf:"default:123,noprint"`
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

	// -------------------------------------------------------------------------
	// App Starting

	log.Info(ctx, "starting service", "version", build)
	defer log.Info(ctx, "shutdown complete")

	out, err := conf.String(&cfg)
	if err != nil {
		return fmt.Errorf("generating config for output: %w", err)
	}
	log.Info(ctx, "startup", "config", out)

	expvar.NewString("build").Set(build)

	// -------------------------------------------------------------------------
	// Initialize keystore

	log.Info(ctx, "startup", "status", "initializing keystore")

	ks := keystore.New(log)

	if err := ks.LoadAuthKeys(cfg.Auth.KeysFolder); err != nil {
		return fmt.Errorf("reading keys: %w", err)
	}

	if err := ks.LoadBankKeys(cfg.Bank.KeysFolder, cfg.Bank.PassPhrase); err != nil {
		return fmt.Errorf("reading keys: %w", err)
	}

	// -------------------------------------------------------------------------
	// Database Support

	log.Info(ctx, "startup", "status", "initializing database support", "hostport", cfg.DB.HostPort)

	db, err := sqldb.Open(sqldb.Config{
		User:         cfg.DB.User,
		Password:     cfg.DB.Password,
		HostPort:     cfg.DB.HostPort,
		Name:         cfg.DB.Name,
		MaxIdleConns: cfg.DB.MaxIdleConns,
		MaxOpenConns: cfg.DB.MaxOpenConns,
		DisableTLS:   cfg.DB.DisableTLS,
	})
	if err != nil {
		return fmt.Errorf("connecting to db: %w", err)
	}
	defer func() {
		log.Info(ctx, "shutdown", "status", "stopping database support", "hostport", cfg.DB.HostPort)
		db.Close()
	}()

	// -------------------------------------------------------------------------
	// Initialize authentication support

	log.Info(ctx, "startup", "status", "initializing authentication support")

	authCfg := auth.Config{
		KeyLookup: ks,
	}

	authClient, err := auth.New(authCfg)
	if err != nil {
		return fmt.Errorf("constructing authClient: %w", err)
	}

	// -------------------------------------------------------------------------
	// Create the currency converter and bankClient needed for the game

	if cfg.Game.ContractID == "0x0" {
		return errors.New("smart contract id not provided")
	}

	converter, err := currency.NewConverter(scbank.BankMetaData.ABI, cfg.Bank.CoinMarketCapKey)
	if err != nil {
		log.Info(ctx, "unable to create converter, using default", "ERROR", err)
		converter = currency.NewDefaultConverter(scbank.BankMetaData.ABI)
	}

	oneETHToUSD, oneUSDToETH := converter.Values()
	log.Info(ctx, "currency values", "oneETHToUSD", oneETHToUSD, "oneUSDToETH", oneUSDToETH)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	backend, err := ethereum.CreateDialedBackend(ctx, cfg.Bank.Network)
	if err != nil {
		return fmt.Errorf("ethereum backend: %w", err)
	}
	defer backend.Close()

	privateKeyPEM, err := ks.PrivateKey(cfg.Bank.KeyID)
	if err != nil {
		return fmt.Errorf("capture private key: %w", err)
	}

	block, _ := pem.Decode([]byte(privateKeyPEM))
	ecdsaKey, err := crypto.ToECDSA(block.Bytes)
	if err != nil {
		return fmt.Errorf("error converting PEM to ECDSA: %w", err)
	}

	bankClient, err := bank.New(ctx, log, backend, ecdsaKey, common.HexToAddress(cfg.Game.ContractID))
	if err != nil {
		return fmt.Errorf("connecting to bankClient: %w", err)
	}

	// -------------------------------------------------------------------------
	// Start Debug Service

	log.Info(ctx, "startup", "status", "debug v1 router started", "host", cfg.Web.DebugHost)

	go func() {
		if err := http.ListenAndServe(cfg.Web.DebugHost, debug.Mux()); err != nil {
			log.Error(ctx, "shutdown", "status", "debug v1 router closed", "host", cfg.Web.DebugHost, "ERROR", err)
		}
	}()

	// -------------------------------------------------------------------------
	// Start API Service

	log.Info(ctx, "startup", "status", "initializing V1 API support")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	cfgMux := mux.Config{
		Shutdown:       shutdown,
		Log:            log,
		Auth:           authClient,
		Converter:      converter,
		Bank:           bankClient,
		DB:             db,
		AnteUSD:        cfg.Game.AnteUSD,
		ActiveKID:      cfg.Auth.ActiveKID,
		BankTimeout:    cfg.Bank.Timeout,
		ConnectTimeout: cfg.Game.ConnectTimeout,
	}

	api := http.Server{
		Addr:         cfg.Web.APIHost,
		Handler:      mux.WebAPI(cfgMux, all.Routes(), mux.WithCORS(cfg.Web.CORSAllowedOrigins)),
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
		IdleTimeout:  cfg.Web.IdleTimeout,
		ErrorLog:     logger.NewStdLogger(log, logger.LevelError),
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Info(ctx, "startup", "status", "api router started", "host", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	// -------------------------------------------------------------------------
	// Shutdown

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		log.Info(ctx, "shutdown", "status", "shutdown started", "signal", sig)
		defer log.Info(ctx, "shutdown", "status", "shutdown complete", "signal", sig)

		ctx, cancel := context.WithTimeout(context.Background(), cfg.Web.ShutdownTimeout)
		defer cancel()

		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}
