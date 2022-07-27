package tests

import (
	"context"
	"math"
	"math/big"
	"testing"

	"github.com/ardanlabs/liarsdice/foundation/smartcontract/smart"
	"github.com/ethereum/go-ethereum/common"
)

func TestDeposit(t *testing.T) {
	ctx := context.Background()

	// Deploy the contract so it can be tested.
	contract, err := deployContract(ctx, PrimaryKeyPath, PrimaryPassPhrase)
	if err != nil {
		t.Fatalf("unexpected deploy error: %s", err)
	}

	// Create a client connection on behalf of player1.
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

	// WE NEED MORE. WE NEED TO CHECK THE PLAYERS WALLET BALANCE AS WELL TO BE ACCURATE.
}
