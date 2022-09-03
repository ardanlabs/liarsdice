package main

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"time"

	"github.com/ardanlabs/liarsdice/app/tooling/admin/commands"
	"github.com/ardanlabs/liarsdice/business/core/bank"
	"github.com/ardanlabs/liarsdice/foundation/smart/currency"
)

const (
	keyStorePath = "zarf/ethereum/keystore/"
	passPhrase   = "123"
)

func main() {
	if len(os.Args) == 1 {
		commands.PrintUsage()
		return
	}

	if err := run(); err != nil {
		fmt.Println("\nError")
		fmt.Println("----------------------------------------------------")
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	// =========================================================================
	// Parse flags for settings.

	flags, args, err := commands.Parse()
	if err != nil {
		return err
	}

	if _, exists := flags["h"]; exists {
		commands.PrintUsage()
		return nil
	}

	// =========================================================================
	// Find the key file for the specified address.

	keyFile, err := findKeyFile(keyStorePath, args.Address)
	if err != nil {
		return err
	}

	// =========================================================================
	// Construct the converted for ETH to USD conversions.

	converter, err := currency.NewConverter(args.CoinMarketCapKey)
	if err != nil {
		converter = currency.NewDefaultConverter()
	}
	oneETHToUSD, oneUSDToETH := converter.Values()

	// =========================================================================
	// Construct the bank API.

	bank, err := bank.New(ctx, args.Network, keyFile, args.PassPhrase, args.ContractID)
	if err != nil {
		return err
	}

	// =========================================================================
	// Display the settings and execute the specified command.

	fmt.Println("\nSettings")
	fmt.Println("----------------------------------------------------")
	fmt.Println("network         :", args.Network)
	fmt.Println("privatekey      :", keyFile)
	fmt.Println("passphrase      :", args.PassPhrase)
	fmt.Println("oneETHToUSD     :", oneETHToUSD)
	fmt.Println("oneUSDToETH     :", oneUSDToETH)
	fmt.Println("key address     :", bank.Client().Address())
	fmt.Println("contract id     :", args.ContractID)

	if _, exists := flags["t"]; exists {
		fmt.Println("tran id         :", args.TranID)
		return commands.Transaction(ctx, converter, bank, args.TranID)
	}
	if _, exists := flags["b"]; exists {
		fmt.Println("address         :", args.Address)
		return commands.Balance(ctx, converter, bank, args.Address)
	}
	if _, exists := flags["w"]; exists {
		return commands.Wallet(ctx, converter, bank)
	}
	if _, exists := flags["d"]; exists {
		return commands.Deploy(ctx, converter, bank, args, keyFile)
	}
	if _, exists := flags["a"]; exists {
		fmt.Println("amount          :", args.Amount)
		return commands.Deposit(ctx, converter, bank, args.Amount)
	}
	if _, exists := flags["r"]; exists {
		return commands.Withdraw(ctx, converter, bank)
	}

	return errors.New("no functional command provided")
}

// =============================================================================

// findKeyFile searches the keystore for the specified address key file.
func findKeyFile(keyStorePath string, address string) (string, error) {
	keyStorePath = strings.TrimSuffix(keyStorePath, "/")
	errFound := errors.New("found")

	var filePath string
	fn := func(fileName string, dirEntry fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("walkdir failure: %w", err)
		}

		if dirEntry.IsDir() {
			return nil
		}

		if strings.Contains(fileName, address[2:]) {
			filePath = fmt.Sprintf("%s/%s", keyStorePath, fileName)
			return errFound
		}

		return nil
	}

	if err := fs.WalkDir(os.DirFS(keyStorePath), ".", fn); err != nil {
		if errors.Is(err, errFound) {
			return filePath, nil
		}
		return "", fmt.Errorf("walking directory: %w", err)
	}

	return "", errors.New("not found")
}
