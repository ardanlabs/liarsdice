package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	scBank "github.com/ardanlabs/liarsdice/business/contract/go/bank"
	"github.com/ardanlabs/liarsdice/business/core/bank"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "admin",
	Short: "A small tool to manage liars dice",
	Long: `Provides the ability to deploy the contract, move funds, check balances as well as
initialize vault and load keys into it.`,
	Version: "set me",
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
	defaultFileKey          = "6327a38415c53ffb36c11db55ea74cc9cb4976fd"
)

func init() {
	rootCmd.PersistentFlags().StringP("network", "n", defaultNetwork, "Sets the network to use.")
	rootCmd.PersistentFlags().StringP("coin-market-cap-key", "c", defaultCoinMarketCapKey, "Key that references market cap.")
	rootCmd.PersistentFlags().StringP("key-path", "k", "", "The key path to use.")
	rootCmd.PersistentFlags().StringP("file-key", "f", defaultFileKey, "The file key to use.")
	rootCmd.PersistentFlags().StringP("passphrase", "p", "", "The pass phrase to use.")
}

func getDependencies(ctx context.Context, cmd *cobra.Command) (*currency.Converter, *ethereum.Client, *bank.Bank, error) {
	coinMarketCapKey, err := cmd.Flags().GetString("coin-market-cap-key")
	if err != nil {
		return nil, nil, nil, err
	}

	bankNetwork, err := cmd.Flags().GetString("network")
	if err != nil {
		return nil, nil, nil, err
	}

	fileKey, err := cmd.Flags().GetString("file-key")
	if err != nil {
		return nil, nil, nil, err
	}

	passPhrase, err := cmd.Flags().GetString("passphrase")
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

	privateKey, err := ethereum.PrivateKeyByKeyFile(fileKey, passPhrase)
	if err != nil {
		return nil, nil, nil, errors.New("capture private key")
	}

	ethClient, err := ethereum.NewClient(backend, privateKey)
	if err != nil {
		return nil, nil, nil, err
	}

	contractID, err := cmd.Flags().GetString("contract-id")
	if err != nil {
		return nil, nil, nil, err
	}

	bankClient, err := bank.New(ctx, nil, backend, privateKey, common.HexToAddress(contractID))
	if err != nil {
		return nil, nil, nil, fmt.Errorf("connecting to bankClient: %w", err)
	}

	return converter, ethClient, bankClient, nil
}
