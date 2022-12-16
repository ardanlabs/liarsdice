package cmd

import (
	"context"
	"encoding/pem"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"io/fs"
	"os"
	"path"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/liarsdice/foundation/vault"
)

// vaultCmd represents the vault command
var vaultAddKeysCmd = &cobra.Command{
	Use:   "add-keys",
	Short: "Add pem keys to vault.",
	Long:  `Used to load pem keys into vault for liars dice.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		vaultConfig, err := getVaultConfig(cmd)
		if err != nil {
			return err
		}

		vaultSrv, err := vault.New(vaultConfig)
		if err != nil {
			return fmt.Errorf("constructing vault: %w", err)
		}

		kFolder, err := cmd.Flags().GetString(keysFolder)
		if err != nil {
			return err
		}

		err = loadKeys(ctx, os.DirFS(kFolder), vaultSrv)
		if err != nil {
			return err
		}

		ksPath, err := cmd.Flags().GetString(keyStorePath)
		if err != nil {
			return err
		}

		passPhrase, err := cmd.Flags().GetString(passPhrase)
		if err != nil {
			return err
		}

		err = loadBankKeys(ctx, ksPath, passPhrase, vaultSrv)
		if err != nil {
			return err
		}

		return nil
	},
}

const defaultKeysFolder = "zarf/keys/"

func init() {
	vaultCmd.AddCommand(vaultAddKeysCmd)
	vaultAddKeysCmd.Flags().StringP(keysFolder, shortName[keysFolder], defaultKeysFolder, "The folder of keys to be loaded into vault")
}

func loadBankKeys(ctx context.Context, ksPath, passPhrase string, vaultSrv *vault.Vault) error {
	fSys := os.DirFS(ksPath)

	fn := func(fileName string, dirEntry fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("walkdir failure: %w", err)
		}

		if dirEntry.IsDir() {
			return nil
		}

		privateKey, err := ethereum.PrivateKeyByKeyFile(fmt.Sprintf("%s%s", ksPath, fileName), passPhrase)
		if err != nil {
			return fmt.Errorf("capture private key: %s", err)
		}

		kid := strings.Split(dirEntry.Name(), "Z--")
		if len(kid) != 2 {
			return fmt.Errorf("misformed file name: %s", dirEntry.Name())
		}
		fmt.Println("Loading kid:", kid[1])

		privatePEM := pem.EncodeToMemory(&pem.Block{
			Type:  "EC PRIVATE KEY",
			Bytes: crypto.FromECDSA(privateKey),
		})

		if err := vaultSrv.AddPrivateKey(ctx, kid[1], privatePEM); err != nil {
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

func loadKeys(ctx context.Context, fSys fs.FS, vaultSrv *vault.Vault) error {

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

		if err := vaultSrv.AddPrivateKey(ctx, kid, privatePEM); err != nil {
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
