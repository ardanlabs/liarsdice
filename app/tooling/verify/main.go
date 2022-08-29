package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/ardanlabs/liarsdice/app/tooling/verify/commands"
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

	f, err := commands.Parse()
	if err != nil {
		return fmt.Errorf("parse commands: %v", err)
	}

	fmt.Println("network         :", f.Network)

	switch {
	case f.TX != "":
		return commands.TXHash(ctx, f.Network, f.TX)

	case f.Balance != "":
		return commands.Balance(ctx, f.Network, f.Balance, f.ContractID)
	}

	return nil
}
