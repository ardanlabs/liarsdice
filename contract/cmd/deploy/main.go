package main

import (
	"context"
	"fmt"
	"log"
	"os"

	contract "github.com/ardanlabs/liarsdice/contract/sol/go"
	"github.com/ardanlabs/liarsdice/foundation/smartcontract/smart"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	ctx := context.Background()

	client, err := smart.Connect(ctx, smart.NetworkLocalhost, smart.PrimaryKeyPath, smart.PrimaryPassPhrase)
	if err != nil {
		return err
	}

	fmt.Println("fromAddress:", client.Account)

	// =========================================================================

	startingBalance, err := client.CurrentBalance(ctx)
	if err != nil {
		return err
	}
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

	receipt, err := client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	client.DisplayTransactionReceipt(receipt, tx)

	return nil
}
