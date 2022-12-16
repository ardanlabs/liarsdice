package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	"time"

	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/common"
)

const defaultBankNetwork = "http://geth-service.liars-system.svc.cluster.local:8545"

// transactionCmd represents the transaction command
var transactionCmd = &cobra.Command{
	Use:   "transaction",
	Short: "Examine transaction",
	Long:  `Examine a transaction directly`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tranID, err := cmd.Flags().GetString(transaction)
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
		defer cancel()

		converter, ethClient, _, err := getDependencies(ctx, cmd, "")
		if err != nil {
			return err
		}

		return getTransaction(ctx, converter, ethClient, tranID)
	},
}

func init() {
	rootCmd.AddCommand(transactionCmd)

	transactionCmd.Flags().StringP(transaction, shortName[transaction], "", "Show transaction details for the specified transaction hash.")
}

func getTransaction(ctx context.Context, converter *currency.Converter, ethClient *ethereum.Client, tranID string) error {
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
}
