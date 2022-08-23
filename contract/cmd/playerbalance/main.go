package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ardanlabs/liarsdice/contract/sol/go/contract"
	"github.com/ardanlabs/liarsdice/foundation/smartcontract/smart"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	ctx := context.Background()

	player := common.HexToAddress(smart.Player2Address)

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

	tranOpts, err := client.NewCallOpts(ctx)
	if err != nil {
		return err
	}

	// =========================================================================
	amount, err := contract.PlayerBalance(tranOpts, player)
	if err != nil {
		return err
	}

	fmt.Println(amount)

	return nil
}

// newContract constructs a SimpleCoin contract.
func newContract(client *smart.Client) (*contract.Contract, error) {
	data, err := os.ReadFile("zarf/contract/id.env")
	if err != nil {
		return nil, fmt.Errorf("readfile: %w", err)
	}
	contractID := string(data)
	fmt.Println("contractID:", contractID)

	contract, err := contract.NewContract(common.HexToAddress(contractID), client.ContractBackend())
	if err != nil {
		return nil, fmt.Errorf("NewContract: %w", err)
	}

	return contract, nil
}