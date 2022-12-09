/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

-c to specify contract
-m for amount of USD

admin contract -d
admin contract -b <smart contract balance>
admin contract -a <game contract>
admin contract -r <game contract>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// contractCmd represents the contract command
var contractCmd = &cobra.Command{
	Use:   "contract",
	Short: "Manage contract related items",
	Long:  `Manage contract: deploy contract, show balance, add and remove money, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		fmt.Println("contract called")
	},
}

func init() {
	rootCmd.AddCommand(contractCmd)
	contractCmd.Flags().BoolP("deploy", "d", false, "Deploy the smart contract.")
	contractCmd.Flags().StringP("balance", "b", "", "Show the smart contract balance.")
	contractCmd.Flags().StringP("add-money", "a", "", "Deposit USD into the game contract.")
	contractCmd.Flags().StringP("remove-money", "r", "", "Withdraw money from the game contract.")
	contractCmd.MarkFlagsMutuallyExclusive("deploy", "balance", "add-money", "remove-money")

	contractCmd.Flags().StringP("contract", "c", "", "Provides the contract id for required calls.")
	contractCmd.Flags().StringP("money", "m", "", "Sets the amount of USD to use.")
}
