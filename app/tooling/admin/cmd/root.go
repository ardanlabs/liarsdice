/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const defaultNetwork = "http://geth-service.liars-system.svc.cluster.local:8545"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "admin",
	Short: "A small tool to manage liars dice",
	Long: `Provides the ability to deploy the contract, move funds, check balances as well as
initialize vault and load keys into it.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.admin.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().StringP("network", "n", defaultNetwork, "Sets the network to use.")
}
