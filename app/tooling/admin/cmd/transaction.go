/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/ardanlabs/ethereum/currency"
	scbank "github.com/ardanlabs/liarsdice/business/contract/go/bank"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/ardanlabs/ethereum"
	"github.com/ethereum/go-ethereum/common"
)

const (
	defaultCoinMarketCapKey = "a8cd12fb-d056-423f-877b-659046af0aa5"
	defaultBankNetwork      = "http://geth-service.liars-system.svc.cluster.local:8545"
)

// transactionCmd represents the transaction command
var transactionCmd = &cobra.Command{
	Use:   "transaction",
	Short: "Examine transaction",
	Long:  `Examine a transaction directly`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		tranID, err := cmd.Flags().GetString("transaction")
		if err != nil {
			return err
		}

		coinMarketCapKey, err := cmd.Flags().GetString("coin-market-cap-key")
		if err != nil {
			return err
		}

		bankNetwork, err := cmd.Flags().GetString("network")
		if err != nil {
			return err
		}

		keyPath, err := cmd.Flags().GetString("key-path")
		if err != nil {
			return err
		}

		passPhrase, err := cmd.Flags().GetString("passphrase")
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
		defer cancel()

		converter, err := currency.NewConverter(scbank.BankMetaData.ABI, coinMarketCapKey)
		if err != nil {
			converter = currency.NewDefaultConverter(scbank.BankMetaData.ABI)
		}

		backend, err := ethereum.CreateDialedBackend(ctx, bankNetwork)
		if err != nil {
			return errors.New("ethereum backend")
		}
		defer backend.Close()

		privateKey, err := ethereum.PrivateKeyByKeyFile(keyPath, passPhrase)
		if err != nil {
			return errors.New("capture private key")
		}

		ethClient, err := ethereum.NewClient(backend, privateKey)
		if err != nil {
			return err
		}

		fmt.Println("\nTransaction ID")
		fmt.Println("----------------------------------------------------")
		fmt.Println("tran id         :", tranID)

		txHash := common.HexToHash(tranID)
		tx, pending, err := ethClient.TransactionByHash(ctx, txHash)
		if err != nil {
			return err
		}

		if pending {
			return errors.New("transaction pending")
		}

		fmt.Print(converter.FmtTransaction(tx))

		receipt, err := ethClient.TransactionReceipt(ctx, txHash)
		if err != nil {
			return err
		}

		fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(transactionCmd)

	transactionCmd.Flags().StringP("transaction", "t", "", "Show transaction details for the specified transaction hash.")
	transactionCmd.Flags().StringP("coin-market-cap-key", "c", defaultCoinMarketCapKey, "Key that references market cap.")
	transactionCmd.Flags().StringP("network", "n", defaultBankNetwork, "The bank network to use.")
	transactionCmd.Flags().StringP("key-path", "k", "", "The key path to use.")
	transactionCmd.Flags().StringP("passphrase", "p", "", "The pass phrase to use.")
}
