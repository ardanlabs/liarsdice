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

	// Get player's wallet balance before the transaction.
	startingBalance, err := client.CurrentBalance(ctx)
	if err != nil {
		t.Fatalf("unexpected CurrentBalance error: %s", err)
	}

	const gasLimit = 300000
	const transactionGwei = 40000000

	tranOpts, err := client.NewTransactOpts(ctx, gasLimit, transactionGwei)
	if err != nil {
		t.Fatalf("unexpected NewTransactOpts error: %s", err)
	}

	tx, err := contract.Deposit(tranOpts)
	if err != nil {
		t.Fatalf("unexpected Deposit error: %s", err)
	}

	// The receipt gets the total gas used by the transaction.
	receipt, err := client.WaitMined(ctx, tx)
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

	// We need the transaction value from Gwei to Wei.
	valueWei := transactionGwei * math.Pow(10, 9)

	// We need it to be big.Int for calculations.
	transactionWei := big.NewInt(int64(valueWei))

	// Check the player's balance in the contract.
	if amount.Cmp(transactionWei) != 0 {
		t.Fatalf("expecting player's balance to be %d; got %d", transactionWei, amount)
	}

	// Get the player's wallet balance after the transaction.
	finalBalance, err := client.CurrentBalance(ctx)
	if err != nil {
		t.Fatalf("unexpected CurrentBalance error: %s", err)
	}

	// Calculate the total Gas cost (gas price * used gas)
	totalGasCost := big.NewInt(0)
	totalGasCost.Mul(tx.GasPrice(), big.NewInt(int64(receipt.GasUsed)))

	// Subtract the transaction amount and total gas cost from starting balance.
	expectedBalance := big.NewInt(0)
	expectedBalance.Sub(startingBalance, transactionWei)
	expectedBalance.Sub(expectedBalance, totalGasCost)

	if expectedBalance.Cmp(finalBalance) != 0 {
		t.Fatalf("expecting final balance to be %d; got %d", expectedBalance, finalBalance)
	}
}
