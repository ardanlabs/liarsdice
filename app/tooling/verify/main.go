package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ardanlabs/liarsdice/app/tooling/verify/commands"
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

	f, v, err := commands.Parse()
	if err != nil {
		return err
	}

	converter, err := currency.NewConverter(v.CoinMarketCapKey)
	if err != nil {
		converter = currency.NewDefaultConverter()
	}
	oneETHToUSD, oneUSDToETH := converter.Values()

	bank, err := bank.New(ctx, v.Network, v.KeyFile, v.PassPhrase, v.ContractID)
	if err != nil {
		return err
	}

	fmt.Println("\nSettings")
	fmt.Println("----------------------------------------------------")
	fmt.Println("network         :", v.Network)
	fmt.Println("privatekey      :", v.KeyFile)
	fmt.Println("passphrase      :", v.PassPhrase)
	fmt.Println("oneETHToUSD     :", oneETHToUSD)
	fmt.Println("oneUSDToETH     :", oneUSDToETH)
	fmt.Println("key address     :", bank.Client().Address())

	if _, exists := f["t"]; exists {
		return commands.Transaction(ctx, converter, bank, v)
	}
	if _, exists := f["b"]; exists {
		return commands.Balance(ctx, converter, bank, v)
	}
	if _, exists := f["w"]; exists {
		return commands.Wallet(ctx, converter, bank, v)
	}

	return errors.New("no functional command provided")
}
