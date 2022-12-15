package cmd

import (
	"fmt"
	"github.com/ardanlabs/liarsdice/foundation/vault"
	"github.com/spf13/cobra"
)

// vaultCmd represents the vault command
var vaultCmd = &cobra.Command{
	Use:   "vault",
	Short: "Utility tasks to manage Vault",
	Long: `Can be used to initialize or unseal a Vault instance as well as load keys into vault
for liars dice.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmd.Help()
		}

		fmt.Println("vault called")
		return nil
	},
}

const (
	defaultVaultAddress = "http://vault-service.liars-system.svc.cluster.local:8200"
	defaultMountPath    = "secret"
	defaultToken        = "mytoken"
)

func init() {
	rootCmd.AddCommand(vaultCmd)
	vaultCmd.PersistentFlags().StringP(vaultAddress, shortName[vaultAddress], defaultVaultAddress, "The network address of our vault server")
	vaultCmd.PersistentFlags().StringP(mountPath, shortName[mountPath], defaultMountPath, "The mount path we want to use in vault")
	vaultCmd.PersistentFlags().StringP(token, shortName[token], defaultToken, "The (non-root) token used to access vault")
}

func getVaultConfig(cmd *cobra.Command) (vault.Config, error) {
	var vaultConfig vault.Config
	var err error

	vaultConfig.Address, err = cmd.Flags().GetString("vault-address")
	if err != nil {
		return vault.Config{}, err
	}

	vaultConfig.MountPath, err = cmd.Flags().GetString("mount-path")
	if err != nil {
		return vault.Config{}, err
	}

	vaultConfig.Token, err = cmd.Flags().GetString("token")
	if err != nil {
		return vault.Config{}, err
	}

	return vaultConfig, nil
}
