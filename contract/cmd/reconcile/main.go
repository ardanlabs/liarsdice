package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	ldc "github.com/ardanlabs/liarsdice/contract/sol/go"
	"github.com/ardanlabs/liarsdice/foundation/smartcontract/smart"
	"github.com/ethereum/go-ethereum/common"
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

	winner := common.HexToAddress("0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7")
	loser := common.HexToAddress("0x8e113078adf6888b7ba84967f299f29aece24c55")
	ante := smart.USD2Wei(big.NewInt(10))
	fee := smart.USD2Wei(big.NewInt(1))

	tx, err := contract.Reconcile(tranOpts, winner, []common.Address{loser}, ante, fee)
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
