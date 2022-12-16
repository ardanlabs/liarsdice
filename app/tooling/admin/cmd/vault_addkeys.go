package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"io/fs"
	"os"
	"path"
	"strings"

	"github.com/ardanlabs/liarsdice/foundation/vault"
)

// vaultCmd represents the vault command
var vaultAddKeysCmd = &cobra.Command{
	Use:   "add-keys",
	Short: "Add pem keys to vault.",
	Long:  `Used to load pem keys into vault for liars dice.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		vaultConfig, err := getVaultConfig(cmd)
		if err != nil {
			return err
		}

		kFolder, err := cmd.Flags().GetString(keysFolder)
		if err != nil {
			return err
		}

		return loadKeys(os.DirFS(kFolder), vaultConfig)
	},
}

const defaultKeysFolder = "zarf/keys/"

func init() {
	vaultCmd.AddCommand(vaultAddKeysCmd)
	vaultAddKeysCmd.Flags().StringP(keysFolder, shortName[keysFolder], defaultKeysFolder, "The folder of keys to be loaded into vault")
}

func loadKeys(fSys fs.FS, vaultConfig vault.Config) error {
	vaultSrv, err := vault.New(vaultConfig)
	if err != nil {
		return fmt.Errorf("constructing vault: %w", err)
	}

	fn := func(fileName string, dirEntry fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("walkdir failure: %w", err)
		}

		if dirEntry.IsDir() {
			return nil
		}

		if path.Ext(fileName) != ".pem" {
			return nil
		}

		file, err := fSys.Open(fileName)
		if err != nil {
			return fmt.Errorf("opening key file: %w", err)
		}
		defer func() {
			err = file.Close()
			if err != nil {
				fmt.Printf("Error closing file: %s", err)
			}
		}()

		// limit PEM file size to 1 megabyte. This should be reasonable for
		// almost any PEM file and prevents shenanigans like linking the file
		// to /dev/random or something like that.
		privatePEM, err := io.ReadAll(io.LimitReader(file, 1024*1024))
		if err != nil {
			return fmt.Errorf("reading auth private key: %w", err)
		}

		kid := strings.TrimSuffix(dirEntry.Name(), ".pem")
		fmt.Println("Loading kid:", kid)

		if err := vaultSrv.AddPrivateKey(context.Background(), kid, privatePEM); err != nil {
			return fmt.Errorf("put: %w", err)
		}

		return nil
	}

	fmt.Print("\n")
	if err := fs.WalkDir(fSys, ".", fn); err != nil {
		return fmt.Errorf("walking directory: %w", err)
	}

	return nil
}
