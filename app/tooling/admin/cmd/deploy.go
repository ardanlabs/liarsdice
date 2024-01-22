package cmd

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	scBank "github.com/ardanlabs/liarsdice/business/contract/go/bank"
	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/cobra"
)

// keyCmd represents the key command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy the bank contract",
	Long:  `This deploys the bank smart contract used by liars dice.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
		defer cancel()

		converter, ethClient, _, err := getDependencies(ctx, cmd, "")
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

	// -------------------------------------------------------------------------

	const gasLimit = 1700000
	const valueGwei = 0.0
	tranOpts, err := ethClient.NewTransactOpts(ctx, gasLimit, big.NewInt(0), big.NewFloat(valueGwei))
	if err != nil {
		return err
	}

	// -------------------------------------------------------------------------

	address, tx, _, err := scBank.DeployBank(tranOpts, ethClient.Backend)
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransaction(tx))

	fmt.Println("\nContract Details")
	fmt.Println("----------------------------------------------------")
	fmt.Println("contract id     :", address.Hex())
	fmt.Printf("export GAME_CONTRACT_ID=%s\n", address.Hex())

	// -------------------------------------------------------------------------

	fmt.Println("\nWaiting Logs")
	fmt.Println("----------------------------------------------------")

	log.SetDefault(log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stdout, log.LevelInfo, true)))

	receipt, err := ethClient.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	return nil
}
