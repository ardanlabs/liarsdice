package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"

	ldc "github.com/ardanlabs/liarsdice/contract/sol/go"
	"github.com/ardanlabs/liarsdice/foundation/smartcontract/smart"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	ctx := context.Background()

	const rawurl = smart.NetworkLocalhost
	client, err := smart.Connect(ctx, rawurl, smart.PrimaryKeyPath, smart.PrimaryPassPhrase)
	if err != nil {
		return err
	}

	fmt.Println("fromAddress:", client.Account)

	// =========================================================================

	contract, err := newContract(client)
	if err != nil {
		return err
	}

	// =========================================================================

	startingBalance, err := client.CurrentBalance(ctx)
	if err != nil {
		return err
	}
	defer client.DisplayBalanceSheet(ctx, startingBalance)

	// =========================================================================

	const gasLimit = 300000
	const valueGwei = 0
	tranOpts, err := client.NewTransactOpts(ctx, gasLimit, valueGwei)
	if err != nil {
		return err
	}

	// =========================================================================

	to := common.HexToAddress("0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7")

	tx, err := contract.PlaceAnte(tranOpts, to, big.NewInt(10000000000))
	if err != nil {
		return err
	}

	client.DisplayTransaction(tx)

	receipt, err := client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	client.DisplayTransactionReceipt(receipt, tx)

	return nil
}

// newContract constructs a SimpleCoin contract.
func newContract(client *smart.Client) (*ldc.Contract, error) {
	data, err := os.ReadFile("zarf/contract/id.env")
	if err != nil {
		return nil, fmt.Errorf("readfile: %w", err)
	}
	contractID := string(data)
	fmt.Println("contractID:", contractID)

	contract, err := ldc.NewContract(common.HexToAddress(contractID), client.ContractBackend())
	if err != nil {
		return nil, fmt.Errorf("NewContract: %w", err)
	}

	return contract, nil
}
