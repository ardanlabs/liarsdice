package cmd

import (
	"context"
	"fmt"
	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	scBank "github.com/ardanlabs/liarsdice/business/contract/go/bank"
	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/cobra"
	"math/big"
	"time"
)

// keyCmd represents the key command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
		defer cancel()

		converter, ethClient, _, err := getDependencies(ctx, cmd)
		if err != nil {
			return err
		}

		return deploy(ctx, converter, ethClient)
	},
}

func init() {
	contractCmd.AddCommand(deployCmd)
}

func deploy(ctx context.Context, converter *currency.Converter, ethClient *ethereum.Client) (err error) {
	startingBalance, err := ethClient.Balance(ctx)
	if err != nil {
		return err
	}
	defer func() {
		endingBalance, dErr := ethClient.Balance(ctx)
		if dErr != nil {
			err = dErr
			return
		}
		fmt.Print(converter.FmtBalanceSheet(startingBalance, endingBalance))
	}()

	// =========================================================================

	const gasLimit = 1700000
	const valueGwei = 0.0
	tranOpts, err := ethClient.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		return err
	}

	// =========================================================================

	address, tx, _, err := scBank.DeployBank(tranOpts, ethClient.Backend)
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransaction(tx))

	fmt.Println("\nContract Details")
	fmt.Println("----------------------------------------------------")
	fmt.Println("contract id     :", address.Hex())
	fmt.Printf("export GAME_CONTRACT_ID=%s\n", address.Hex())

	// =========================================================================

	fmt.Println("\nWaiting Logs")
	fmt.Println("----------------------------------------------------")
	log.Root().SetHandler(log.StdoutHandler)

	receipt, err := ethClient.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	return nil
}
