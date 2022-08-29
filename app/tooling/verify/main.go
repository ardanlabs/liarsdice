package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ardanlabs/liarsdice/app/tooling/verify/commands"
	"github.com/ardanlabs/liarsdice/foundation/smart/currency"
)

func main() {
	if len(os.Args) == 1 {
		commands.PrintUsage()
		return
	}

	if err := run(); err != nil {
		fmt.Println("ERROR           :", err)
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

	fmt.Println("\nSettings")
	fmt.Println("----------------------------------------------------")

	fmt.Println("network         :", v.Network)
	fmt.Println("privatekey      :", v.KeyFile)
	fmt.Println("passphrase      :", v.PassPhrase)
	fmt.Println("oneETHToUSD     :", oneETHToUSD)
	fmt.Println("oneUSDToETH     :", oneUSDToETH)

	if _, exists := f["t"]; exists {
		return commands.Transaction(ctx, converter, v)
	}
	if _, exists := f["b"]; exists {
		return commands.Balance(ctx, converter, v)
	}

	return errors.New("no functional command provided")
}
