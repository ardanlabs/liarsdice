package cmd

import (
	"context"
	"fmt"
	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"time"
)

// walletCmd represents the wallet command
var walletCmd = &cobra.Command{
	Use:   "wallet",
	Short: "Show the wallet balance",
	Long:  `Show the wallet balance for the specified smart contract`,
	RunE: func(cmd *cobra.Command, args []string) error {
		address, err := cmd.Flags().GetString(wallet)
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
		defer cancel()

		converter, ethClient, _, err := getDependencies(ctx, cmd, "")
		if err != nil {
			return err
		}

		return getWallet(ctx, converter, ethClient, address)
	},
}

const defaultWalletAddress = "0x8e113078adf6888b7ba84967f299f29aece24c55"

func init() {
	rootCmd.AddCommand(walletCmd)

	walletCmd.Flags().StringP(wallet, shortName[wallet], defaultWalletAddress, "Wallet address")
}

func getWallet(ctx context.Context, converter *currency.Converter, ethClient *ethereum.Client, address string) error {
	fmt.Println("\nWallet Balance")
	fmt.Println("----------------------------------------------------")
	fmt.Println("account         :", address)

	wei, err := ethClient.BalanceAt(ctx, common.HexToAddress(address), nil)
	if err != nil {
		return err
	}

	fmt.Println("wei             :", wei)
	fmt.Println("gwei            :", currency.Wei2GWei(wei))
	fmt.Println("usd             :", converter.Wei2USD(wei))

	return nil
}
