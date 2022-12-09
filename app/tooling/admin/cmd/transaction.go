/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// transactionCmd represents the transaction command
var transactionCmd = &cobra.Command{
	Use:   "transaction",
	Short: "Examine transaction",
	Long:  `Examine a transaction directly`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		fmt.Println("transaction called")
	},
}

func init() {
	rootCmd.AddCommand(transactionCmd)

	transactionCmd.Flags().StringP("transaction", "t", "", "Show transaction details for the specified transaction hash.")
}
