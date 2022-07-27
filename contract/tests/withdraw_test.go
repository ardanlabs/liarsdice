package tests

import (
	"context"
	"math/big"
	"strings"
	"testing"

	"github.com/ardanlabs/liarsdice/foundation/smartcontract/smart"
	"github.com/ethereum/go-ethereum/common"
)

func TestWithdraw(t *testing.T) {
	ctx := context.Background()

	// Deploy the contract so it can be tested.
	contract, err := deployContract(ctx, PrimaryKeyPath, PrimaryPassPhrase)
	if err != nil {
		t.Fatalf("unexpected deploy error: %s", err)
	}

	// Connect as Player1.
	client, err := smart.Connect(ctx, smart.NetworkHTTPLocalhost, Player1KeyPath, Player1PassPhrase)
	if err != nil {
		t.Fatalf("unexpected Connect error: %s", err)
	}

	const gasLimit = 300000
	const valueGwei = 40000000

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

	// Create a new transaction to avoid nonce errors.
	tranOptsWithdraw, err := client.NewTransactOpts(ctx, gasLimit, valueGwei)
	if err != nil {
		t.Fatalf("unexpected NewTransactOpts error: %s", err)
	}

	// We are taking the full amount out of the player's balance.
	txWithdraw, err := contract.Withdraw(tranOptsWithdraw)
	if err != nil {
		t.Fatalf("unexpected Withdraw error: %s", err)
	}

	_, err = client.WaitMined(ctx, txWithdraw)
	if err != nil {
		t.Fatalf("unexpected WaitMined error: %s", err)
	}

	//==========================================================================

	// Connect as the Owner and check the player's balance, it should be zero.
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
		t.Fatalf("expecting player's balance to be zero; got %d", amount)
	}
}

func TestWithdrawWithoutBalance(t *testing.T) {
	ctx := context.Background()

	// Deploy the contract so it can be tested.
	contract, err := deployContract(ctx, PrimaryKeyPath, PrimaryPassPhrase)
	if err != nil {
		t.Fatalf("unexpected deploy error: %s", err)
	}

	// Connect as Player1.
	client, err := smart.Connect(ctx, smart.NetworkHTTPLocalhost, Player1KeyPath, Player1PassPhrase)
	if err != nil {
		t.Fatalf("unexpected Connect error: %s", err)
	}

	const gasLimit = 300000
	const valueGwei = 0

	tranOpts, err := client.NewTransactOpts(ctx, gasLimit, valueGwei)
	if err != nil {
		t.Fatalf("unexpected NewTransactOpts error: %s", err)
	}

	tx, err := contract.Withdraw(tranOpts)
	if err != nil {
		t.Fatalf("unexpected Withdraw error: %s", err)
	}

	_, err = client.WaitMined(ctx, tx)
	if err == nil {
		t.Fatal("expecting revert error; got nil")
	}

	if !strings.Contains(err.Error(), "not enough balance") {
		t.Fatal("expecting revert message 'not enough balance'")
	}
}
