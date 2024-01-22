// Package cmd implements all the functionality needed to manage liarsdice.
package cmd

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	scBank "github.com/ardanlabs/liarsdice/business/contract/go/bank"
	"github.com/ardanlabs/liarsdice/business/core/bank"
	"github.com/ardanlabs/liarsdice/foundation/logger"
	"github.com/ardanlabs/liarsdice/foundation/web"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "admin",
	Short: "A small tool to manage liars dice",
	Long: `Provides the ability to deploy the contract, move funds, check balances as well as
initialize vault and load keys into it.`,
	Version: "1.0.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

const (
	defaultNetwork          = "http://geth-service.liars-system.svc.cluster.local:8545"
	defaultCoinMarketCapKey = "a8cd12fb-d056-423f-877b-659046af0aa5"
	defaultKeyPath          = "zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	defaultPassPhrase       = "123"
	defaultFileKey          = "6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	defaultKeyStorePath     = "zarf/ethereum/keystore/"
)

func init() {
	rootCmd.PersistentFlags().StringP(network, shortName[network], defaultNetwork, "Sets the network to use.")
	rootCmd.PersistentFlags().StringP(keyCoin, shortName[keyCoin], defaultCoinMarketCapKey, "Key that references market cap.")
	rootCmd.PersistentFlags().StringP(keyPath, shortName[keyPath], defaultKeyPath, "The key path to use.")
	rootCmd.PersistentFlags().StringP(keyStorePath, shortName[keyStorePath], defaultKeyStorePath, "The key path to use.")
	rootCmd.PersistentFlags().StringP(passPhrase, shortName[passPhrase], defaultPassPhrase, "The pass phrase to use.")
	rootCmd.PersistentFlags().StringP(
		contractID,
		shortName[contractID],
		getEnv("GAME_CONTRACT_ID", ""),
		"Sets the Contract ID to use.",
	)
}

func getDependencies(ctx context.Context, cmd *cobra.Command, fileKeyKey string) (*currency.Converter, *ethereum.Client, *bank.Bank, error) {
	coinMarketCapKey, err := cmd.Flags().GetString(keyCoin)
	if err != nil {
		return nil, nil, nil, err
	}

	bankNetwork, err := cmd.Flags().GetString(network)
	if err != nil {
		return nil, nil, nil, err
	}

	keyStorePath, err := cmd.Flags().GetString(keyStorePath)
	if err != nil {
		return nil, nil, nil, err
	}

	var fileKey string
	switch len(fileKeyKey) {
	case 0:
		fileKey = defaultFileKey
	default:
		fileKey = fileKeyKey
	}

	keyFile, err := findKeyFile(keyStorePath, fileKey)
	if err != nil {
		return nil, nil, nil, err
	}

	passPhrase, err := cmd.Flags().GetString(passPhrase)
	if err != nil {
		return nil, nil, nil, err
	}

	converter, err := currency.NewConverter(scBank.BankMetaData.ABI, coinMarketCapKey)
	if err != nil {
		converter = currency.NewDefaultConverter(scBank.BankMetaData.ABI)
	}

	backend, err := ethereum.CreateDialedBackend(ctx, bankNetwork)
	if err != nil {
		return nil, nil, nil, errors.New("ethereum backend")
	}
	defer backend.Close()

	privateKey, err := ethereum.PrivateKeyByKeyFile(keyFile, passPhrase)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("capture private key: %w", err)
	}

	ethClient, err := ethereum.NewClient(backend, privateKey)
	if err != nil {
		return nil, nil, nil, err
	}

	contractID, err := cmd.Flags().GetString(contractID)
	if err != nil {
		return nil, nil, nil, err
	}

	var buf bytes.Buffer
	log := logger.New(&buf, logger.LevelInfo, "TEST", func(context.Context) string { return web.GetTraceID(ctx) })

	bankClient, err := bank.New(ctx, log, backend, privateKey, common.HexToAddress(contractID))
	if err != nil {
		return nil, nil, nil, fmt.Errorf("connecting to bankClient: %w", err)
	}

	// -------------------------------------------------------------------------
	// Display the settings and execute the specified command.

	oneETHToUSD, oneUSDToETH := converter.Values()
	fmt.Println("\nSettings")
	fmt.Println("----------------------------------------------------")
	fmt.Println("network         :", bankNetwork)
	fmt.Println("private key     :", keyFile)
	fmt.Println("passphrase      :", passPhrase)
	fmt.Println("oneETHToUSD     :", oneETHToUSD)
	fmt.Println("oneUSDToETH     :", oneUSDToETH)
	fmt.Println("key address     :", bankClient.Client().Address())
	fmt.Println("contract id     :", contractID)

	return converter, ethClient, bankClient, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// findKeyFile searches the keystore for the specified address key file.
func findKeyFile(keyStorePath string, address string) (string, error) {
	keyStorePath = strings.TrimSuffix(keyStorePath, "/")
	errFound := errors.New("found")

	var filePath string
	fn := func(fileName string, dirEntry fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("walkdir failure: %w", err)
		}

		if dirEntry.IsDir() {
			return nil
		}

		if strings.Contains(strings.ToLower(fileName), strings.ToLower(address[2:])) {
			filePath = fmt.Sprintf("%s/%s", keyStorePath, fileName)
			return errFound
		}

		return nil
	}

	if err := fs.WalkDir(os.DirFS(keyStorePath), ".", fn); err != nil {
		if errors.Is(err, errFound) {
			return filePath, nil
		}
		return "", fmt.Errorf("walking directory: %w", err)
	}

	return "", errors.New("not found")
}
