package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ardanlabs/liarsdice/app/tooling/admin/commands"
	"github.com/ardanlabs/liarsdice/business/core/bank"
	"github.com/ardanlabs/liarsdice/foundation/smart/currency"
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

	flags, args, err := commands.Parse()
	if err != nil {
		return err
	}

	if _, exists := flags["h"]; exists {
		commands.PrintUsage()
		return nil
	}

	converter, err := currency.NewConverter(args.CoinMarketCapKey)
	if err != nil {
		converter = currency.NewDefaultConverter()
	}
	oneETHToUSD, oneUSDToETH := converter.Values()

	bank, err := bank.New(ctx, args.Network, args.KeyFile, args.PassPhrase, args.ContractID)
	if err != nil {
		return err
	}

	fmt.Println("\nSettings")
	fmt.Println("----------------------------------------------------")
	fmt.Println("network         :", args.Network)
	fmt.Println("privatekey      :", args.KeyFile)
	fmt.Println("passphrase      :", args.PassPhrase)
	fmt.Println("oneETHToUSD     :", oneETHToUSD)
	fmt.Println("oneUSDToETH     :", oneUSDToETH)
	fmt.Println("key address     :", bank.Client().Address())
	fmt.Println("contract id     :", args.ContractID)

	if _, exists := flags["t"]; exists {
		return commands.Transaction(ctx, converter, bank, args.Hex)
	}
	if _, exists := flags["b"]; exists {
		return commands.Balance(ctx, converter, bank, args.Address)
	}
	if _, exists := flags["w"]; exists {
		return commands.Wallet(ctx, converter, bank)
	}
	if _, exists := flags["d"]; exists {
		return commands.Deploy(ctx, converter, bank, args)
	}
	if _, exists := flags["a"]; exists {
		return commands.Deposit(ctx, converter, bank, args.Amount)
	}
	if _, exists := flags["r"]; exists {
		return commands.Withdraw(ctx, converter, bank)
	}

	return errors.New("no functional command provided")
}
