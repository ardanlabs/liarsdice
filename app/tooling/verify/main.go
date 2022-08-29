package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ardanlabs/liarsdice/foundation/smart/contract"
	"github.com/ardanlabs/liarsdice/foundation/smart/currency"
	"github.com/ethereum/go-ethereum/common"
)

// flags represent the values from the command line.
type flags struct {
	txHash string
}

var f flags

func init() {
	flag.StringVar(&f.txHash, "t", "", "the transaction hash")
	flag.StringVar(&f.txHash, "txhash", "", "the transaction hash")
}

// Harded this here for now just to make life easier.
const (
	primaryKeyPath    = "zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	primaryPassPhrase = "123"
	coinMarketCapKey  = "a8cd12fb-d056-423f-877b-659046af0aa5"
	network           = contract.NetworkLocalhost
)

func main() {
	if err := run(); err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}

func run() (dErr error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	converter, err := currency.NewConverter(coinMarketCapKey)
	if err != nil {
		fmt.Println("unable to create converter, using default:", err)
		converter = currency.NewDefaultConverter()
	}

	client, err := contract.NewClient(ctx, network, primaryKeyPath, primaryPassPhrase)
	if err != nil {
		return err
	}

	oneETHToUSD, oneUSDToETH := converter.Values()

	fmt.Println("network         :", network)
	fmt.Println("fromAddress     :", client.Account())
	fmt.Println("oneETHToUSD     :", oneETHToUSD)
	fmt.Println("oneUSDToETH     :", oneUSDToETH)

	// =========================================================================

	txHash := common.HexToHash(f.txHash)
	tx, pending, err := client.TransactionByHash(ctx, txHash)
	if err != nil {
		return err
	}

	fmt.Print("Transaction Pending")
	if pending {
		fmt.Println("Transaction Pending")
		return nil
	}
	fmt.Print(converter.FmtTransaction(tx))

	receipt, err := client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return err
	}

	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	return nil
}
