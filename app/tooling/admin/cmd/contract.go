package cmd

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ardanlabs/ethereum/currency"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/ardanlabs/liarsdice/business/core/bank"
)

// contractCmd represents the contract command
var contractCmd = &cobra.Command{
	Use:   "contract",
	Short: "Manage contract related items",
	Long:  `Manage contract: deploy contract, show balance, add and remove money, etc.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Flags().NFlag() == 0 {
			return cmd.Help()
		}

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
		defer cancel()

		converter, _, bankClient, err := getDependencies(ctx, cmd)
		if err != nil {
			return err
		}

		getBalance, err := cmd.Flags().GetString("balance")
		if err != nil {
			return nil
		}

		addMoney, err := cmd.Flags().GetString("add-money")
		if err != nil {
			return nil
		}

		removeMoney, err := cmd.Flags().GetString("remove-money")
		if err != nil {
			return nil
		}

		if len(getBalance) != 0 {
			//contract, err := cmd.Flags().GetString("contract")
			//if err != nil {
			//	return err
			//}

			return balance(ctx, converter, bankClient, getBalance)
		}

		if len(addMoney) != 0 {
			amountUSD, err := cmd.Flags().GetFloat64("money")
			if err != nil {
				return err
			}

			if amountUSD == 0 {
				return errors.New("must set money value to greater than 0")
			}

			return deposit(ctx, converter, bankClient, amountUSD)
		}

		if len(removeMoney) != 0 {
			return withdraw(ctx, converter, bankClient)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(contractCmd)
	contractCmd.Flags().StringP("balance", "b", "", "Show the smart contract balance.")
	contractCmd.Flags().StringP("add-money", "a", "", "Deposit USD into the game contract.")
	contractCmd.Flags().StringP("remove-money", "r", "", "Withdraw money from the game contract.")
	contractCmd.MarkFlagsMutuallyExclusive("balance", "add-money", "remove-money")

	contractCmd.Flags().StringP("contract", "c", "", "Provides the contract id for required calls.")
	contractCmd.Flags().Float64P("money", "m", 0, "Sets the amount of USD to use.")
}

func deposit(ctx context.Context, converter *currency.Converter, bankClient *bank.Bank, amountUSD float64) error {
	fmt.Println("\nDeposit Details")
	fmt.Println("----------------------------------------------------")
	fmt.Println("address         :", bankClient.Client().Address())
	fmt.Println("amount          :", amountUSD)

	amountGWei := converter.USD2GWei(big.NewFloat(amountUSD))
	tx, receipt, err := bankClient.Deposit(ctx, amountGWei)
	if err != nil {
		return err
	}

	fmt.Print(converter.FmtTransaction(tx))
	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	return nil
}

func withdraw(ctx context.Context, converter *currency.Converter, bankClient *bank.Bank) error {
	fmt.Println("\nWithdraw Details")
	fmt.Println("----------------------------------------------------")
	fmt.Println("address         :", bankClient.Client().Address())

	tx, receipt, err := bankClient.Withdraw(ctx)
	if err != nil {
		return err
	}

	fmt.Print(converter.FmtTransaction(tx))
	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	return nil
}

func balance(ctx context.Context, converter *currency.Converter, bankClient *bank.Bank, address string) error {
	fmt.Println("\nGame Balance")
	fmt.Println("----------------------------------------------------")
	fmt.Println("account         :", address)

	gwei, err := bankClient.AccountBalance(ctx, common.HexToAddress(address))
	if err != nil {
		return err
	}

	fmt.Println("wei             :", currency.GWei2Wei(gwei))
	fmt.Println("gwei            :", gwei)
	fmt.Println("usd             :", converter.GWei2USD(gwei))

	return nil
}
