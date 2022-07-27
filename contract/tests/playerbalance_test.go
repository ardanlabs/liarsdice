package tests

import (
	"context"
	"math"
	"math/big"
	"testing"

	"github.com/ardanlabs/liarsdice/foundation/smartcontract/smart"
	"github.com/ethereum/go-ethereum/common"
)

func TestPlayerBalance(t *testing.T) {
	ctx := context.Background()

	// Deploy the contract so it can be tested.
	contract, err := deployContract(ctx, PrimaryKeyPath, PrimaryPassPhrase)
	if err != nil {
		t.Fatalf("unexpected deploy error: %s", err)
	}

	// Connect as a Player1 and make deposit.
	client, err := smart.Connect(ctx, smart.NetworkHTTPLocalhost, Player1KeyPath, Player1PassPhrase)
	if err != nil {
		t.Fatalf("unexpected Connect error: %s", err)
	}

	const gasLimit = 300000
	const valueGwei = 40000000

	// 1 Wei == 1 Gwei * (10^9).
	valueWei := valueGwei * math.Pow(10, 9)
	expectedWei := big.NewInt(int64(valueWei))

	tranOpts, err := client.NewTransactOpts(ctx, gasLimit, valueGwei)
	if err != nil {
		t.Fatalf("unexpected NewTransactOpts error: %s", err)
	}

	tx, err := contract.Deposit(tranOpts)
	if err != nil {
		t.Fatalf("unexpected Deposit error: %s", err)
	}

	_, err = client.WaitMined(ctx, tx)
	if err != nil {
		t.Fatalf("unexpected WaitMined error: %s", err)
	}

	//==========================================================================

	// Connect as the Owner to get the player's balance.
	ownerClient, err := smart.Connect(ctx, smart.NetworkHTTPLocalhost, PrimaryKeyPath, PrimaryPassPhrase)
	if err != nil {
		t.Fatalf("unexpected Connect error: %s", err)
	}

	callOpts, err := ownerClient.NewCallOpts(ctx)
	if err != nil {
		t.Fatalf("unexpected NewCallOpts error: %s", err)
	}

	amount, err := contract.PlayerBalance(callOpts, common.HexToAddress(Player1Address))
	if err != nil {
		t.Fatalf("unexpected PlayerBalance error: %s", err)
	}

	if amount.Cmp(expectedWei) != 0 {
		t.Fatalf("expecting player's balance to be %d; got %d", expectedWei, amount)
	}
}

func TestEmptyPlayerBalance(t *testing.T) {
	ctx := context.Background()

	// Deploy the contract so it can be tested.
	contract, err := deployContract(ctx, PrimaryKeyPath, PrimaryPassPhrase)
	if err != nil {
		t.Fatalf("unexpected deploy error: %s", err)
	}

	// Connect as the Owner to get the player's balance.
	ownerClient, err := smart.Connect(ctx, smart.NetworkHTTPLocalhost, PrimaryKeyPath, PrimaryPassPhrase)
	if err != nil {
		t.Fatalf("unexpected Connect error: %s", err)
	}

	callOpts, err := ownerClient.NewCallOpts(ctx)
	if err != nil {
		t.Fatalf("unexpected NewCallOpts error: %s", err)
	}

	amount, err := contract.PlayerBalance(callOpts, common.HexToAddress(Player1Address))
	if err != nil {
		t.Fatalf("unexpected PlayerBalance error: %s", err)
	}

	if amount.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("expecting deposit %d; got %d", 0, amount)
	}
}
