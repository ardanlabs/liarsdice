package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ardanlabs/liarsdice/contract/sol/go/contract"
	"github.com/ardanlabs/liarsdice/foundation/smartcontract/smart"
	"github.com/ethereum/go-ethereum/log"
)

// Harded this here for now just to make life easier.
const (
	primaryKeyPath    = "zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	primaryPassPhrase = "123"
)

func main() {
	if err := run(); err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()
	network := smart.NetworkLocalhost

	client, err := smart.Connect(ctx, network, primaryKeyPath, primaryPassPhrase)
	if err != nil {
		return err
	}

	fmt.Println("network    :", network)
	fmt.Println("fromAddress:", client.Account())

	// =========================================================================

	startingBalance, err := client.CurrentBalance(ctx)
	if err != nil {
		return err
	}
	fmt.Println("starting Ba:", smart.Wei2GWei(startingBalance))
	defer func() {
		fmt.Print(client.FmtBalanceSheet(ctx, startingBalance))
	}()

	// =========================================================================

	const gasLimit = 3000000
	const valueGwei = 0.0
	tranOpts, err := client.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		return err
	}

	// =========================================================================

	address, tx, _, err := contract.DeployContract(tranOpts, client.ContractBackend())
	if err != nil {
		return err
	}
	fmt.Print(smart.FmtTransaction(tx))

	os.MkdirAll("zarf/contract/", 0755)
	if err := os.WriteFile("zarf/contract/id.env", []byte(address.Hex()), 0666); err != nil {
		return err
	}

	fmt.Println("\nContract Details")
	fmt.Println("----------------------------------------------------")
	fmt.Println("Contract ID:", address.Hex())
	fmt.Printf("Please export GAME_CONTRACT_ID=%s\n", address.Hex())

	// =========================================================================

	clientWait, err := smart.Connect(ctx, network, primaryKeyPath, primaryPassPhrase)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	log.Root().SetHandler(log.StdoutHandler)
	receipt, err := clientWait.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	fmt.Print(smart.FmtTransactionReceipt(receipt, tx.GasPrice()))

	return nil
}
