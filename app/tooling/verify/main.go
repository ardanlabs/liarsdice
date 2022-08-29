package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

// Harded this here for now just to make life easier.
const (
	keyPath          = "zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	passPhrase       = "123"
	coinMarketCapKey = "a8cd12fb-d056-423f-877b-659046af0aa5"
)

func main() {
	log := log.New(os.Stderr, "", 0)

	if len(os.Args) == 1 {
		PrintUsage(log)
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

	f, err := Parse()
	if err != nil {
		return fmt.Errorf("parse commands: %v", err)
	}

	switch {
	case f.TX != "":
		return txHash(ctx, f.Network, f.TX)

	case f.Balance != "":
		return balance(ctx, f.Network, f.Balance, f.ContractID)
	}

	return nil
}
