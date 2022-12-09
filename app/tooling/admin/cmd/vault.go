/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// vaultCmd represents the vault command
var vaultCmd = &cobra.Command{
	Use:   "vault",
	Short: "Utility tasks to manage Vault",
	Long: `Can be used to initialize or unseal a Vault instance as well as load keys into vault
for liars dice.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		fmt.Println("vault called")
	},
}

func init() {
	rootCmd.AddCommand(vaultCmd)
	vaultCmd.Flags().BoolP("init", "i", false, "Initialize or unseal a Vault instance")
	vaultCmd.Flags().BoolP("add-keys", "k", false, "Add keys to Vault")
	vaultCmd.MarkFlagsMutuallyExclusive("init", "add-keys")
}
