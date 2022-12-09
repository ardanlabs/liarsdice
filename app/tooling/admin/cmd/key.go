/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

admin key -f :: keyfile
admin key -p :: passphrase
admin key -k :: keycoin
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	defaultFileKey    = "0x6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	defaultPassphrase = "123"
	defaultKeyCoin    = "a8cd12fb-d056-423f-877b-659046af0aa5"
)

// keyCmd represents the key command
var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		fmt.Println("key called")
	},
}

func init() {
	rootCmd.AddCommand(keyCmd)

	keyCmd.Flags().StringP("file-key", "f", defaultFileKey, "Sets the private key file to use for blockchain calls.")
	keyCmd.Flags().StringP("passphrase", "p", defaultPassphrase, "Sets the pass phrase for the key file.")
	keyCmd.Flags().StringP("key-coin", "k", defaultKeyCoin, "Sets the key for the coin market cap API.")
}
