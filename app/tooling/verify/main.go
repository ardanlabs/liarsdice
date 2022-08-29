package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ardanlabs/liarsdice/business/core/bank"
	"github.com/ardanlabs/liarsdice/foundation/smart/contract"
	"github.com/ardanlabs/liarsdice/foundation/smart/currency"
	"github.com/ethereum/go-ethereum/common"
)

// Harded this here for now just to make life easier.
const (
	keyPath          = "zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	passPhrase       = "123"
	coinMarketCapKey = "a8cd12fb-d056-423f-877b-659046af0aa5"
	network          = contract.NetworkLocalhost
	contractID       = "0xE7811C584E23419e1952fa3158DEED345901bd0e"
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

func run() (dErr error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	f, err := Parse()
	if err != nil {
		return fmt.Errorf("parse commands: %v", err)
	}

	switch {
	case f.TXHash != "":
		return txHash(ctx, f.TXHash)
	case f.Balance != "":
		return balance(ctx, f.Balance)
	}

	return nil
}

func balance(ctx context.Context, address string) error {
	_, err := bank.New(ctx, network, keyPath, passPhrase, contractID)
	if err != nil {
		return err
	}

	return nil
}

func txHash(ctx context.Context, hash string) error {
	converter, err := currency.NewConverter(coinMarketCapKey)
	if err != nil {
		log.Println("unable to create converter, using default:", err)
		converter = currency.NewDefaultConverter()
	}

	client, err := contract.NewClient(ctx, network, keyPath, passPhrase)
	if err != nil {
		return err
	}

	oneETHToUSD, oneUSDToETH := converter.Values()

	fmt.Println("network         :", network)
	fmt.Println("fromAddress     :", client.Account())
	fmt.Println("oneETHToUSD     :", oneETHToUSD)
	fmt.Println("oneUSDToETH     :", oneUSDToETH)

	// =========================================================================

	fmt.Println("hash to check   :", hash)

	txHash := common.HexToHash(hash)
	tx, pending, err := client.TransactionByHash(ctx, txHash)
	if err != nil {
		log.Print("status            : Not Found")
		return err
	}

	if pending {
		log.Print("status            : Pending")
		return nil
	}
	log.Print("status            : Completed")
	fmt.Print(converter.FmtTransaction(tx))

	receipt, err := client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return err
	}

	fmt.Println("Output          :")
	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	return nil
}
