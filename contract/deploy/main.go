package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/ardanlabs/liarsdice/contract/sol/go/contract"
	"github.com/ardanlabs/liarsdice/foundation/smartcontract/smart"
	"github.com/ethereum/go-ethereum/log"
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

	client, err := smart.Connect(ctx, network, smart.PrimaryKeyPath, smart.PrimaryPassPhrase)
	if err != nil {
		return err
	}

	fmt.Println("network    :", network)
	fmt.Println("fromAddress:", client.Account)

	// =========================================================================

	startingBalance, err := client.CurrentBalance(ctx)
	if err != nil {
		return err
	}
	fmt.Println("starting Ba:", smart.Wei2GWei(startingBalance))
	defer client.DisplayBalanceSheet(ctx, startingBalance)

	// =========================================================================

	const gasLimit = 3000000
	const valueGwei = 0
	tranOpts, err := client.NewTransactOpts(ctx, gasLimit, valueGwei)
	if err != nil {
		return err
	}

	// =========================================================================

	address, tx, _, err := contract.DeployContract(tranOpts, client.ContractBackend())
	if err != nil {
		return err
	}
	client.DisplayTransaction(tx)

	os.MkdirAll("zarf/contract/", 0755)
	if err := os.WriteFile("zarf/contract/id.env", []byte(address.Hex()), 0666); err != nil {
		return err
	}

	fmt.Println("\nContract Details")
	fmt.Println("----------------------------------------------------")
	fmt.Println("Contract ID:", address.Hex())
	fmt.Println("Please export GAME_CONTRACT_ID=", address.Hex())

	// =========================================================================

	log.Root().SetHandler(log.StdoutHandler)

	for {
		fmt.Println("*** Establish new connection ***")

		client, err := smart.Connect(ctx, network, smart.PrimaryKeyPath, smart.PrimaryPassPhrase)
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		receipt, err := client.WaitMined(ctx, tx)
		cancel()

		if err != nil {
			continue
		}

		client.DisplayTransactionReceipt(receipt, tx)
		break
	}

	return nil
}
