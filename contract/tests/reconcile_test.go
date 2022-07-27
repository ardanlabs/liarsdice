package tests

import (
	"context"
	"math/big"
	"testing"

	"github.com/ardanlabs/liarsdice/foundation/smartcontract/smart"
	"github.com/ethereum/go-ethereum/common"
)

func TestReconcile(t *testing.T) {
	ctx := context.Background()

	// Deploy the contract so it can be tested.
	contract, err := deployContract(ctx, PrimaryKeyPath, PrimaryPassPhrase)
	if err != nil {
		t.Fatalf("unexpected deploy error: %s", err)
	}

	var gasLimit uint64 = 300000
	var valueGwei uint64 = 40000000

	//==========================================================================

	// Connect as Player1 and make deposit.
	clientPlayer1, err := smart.Connect(ctx, smart.NetworkHTTPLocalhost, Player1KeyPath, Player1PassPhrase)
	if err != nil {
		t.Fatalf("unexpected Connect error: %s", err)
	}

	tranOpts, err := clientPlayer1.NewTransactOpts(ctx, gasLimit, valueGwei)
	if err != nil {
		t.Fatalf("unexpected NewTransactOpts error: %s", err)
	}

	tx, err := contract.Deposit(tranOpts)
	if err != nil {
		t.Fatalf("unexpected Deposit error: %s", err)
	}

	_, err = clientPlayer1.WaitMined(ctx, tx)
	if err != nil {
		t.Fatalf("unexpected WaitMined error: %s", err)
	}

	//==========================================================================

	// Connect as Player2 and make deposit.
	clientPlayer2, err := smart.Connect(ctx, smart.NetworkHTTPLocalhost, Player2KeyPath, Player2PassPhrase)
	if err != nil {
		t.Fatalf("unexpected Connect error: %s", err)
	}

	tranOpts, err = clientPlayer2.NewTransactOpts(ctx, gasLimit, valueGwei)
	if err != nil {
		t.Fatalf("unexpected NewTransactOpts error: %s", err)
	}

	tx, err = contract.Deposit(tranOpts)
	if err != nil {
		t.Fatalf("unexpected Deposit error: %s", err)
	}

	_, err = clientPlayer2.WaitMined(ctx, tx)
	if err != nil {
		t.Fatalf("unexpected WaitMined error: %s", err)
	}

	//==========================================================================

	// Connect as the Owner to call Reconcile.
	ownerClient, err := smart.Connect(ctx, smart.NetworkHTTPLocalhost, PrimaryKeyPath, PrimaryPassPhrase)
	if err != nil {
		t.Fatalf("unexpected Connect error: %s", err)
	}

	tranOpts, err = ownerClient.NewTransactOpts(ctx, gasLimit, 0)
	if err != nil {
		t.Fatalf("unexpected NewTransactOpts error: %s", err)
	}

	winner := common.HexToAddress(Player1Address)
	losers := []common.Address{
		common.HexToAddress(Player2Address),
	}

	// Values represented in Wei.
	ante := big.NewInt(20000000000000000)
	gameFee := big.NewInt(10000000000000000)

	_, err = contract.Reconcile(tranOpts, winner, losers, ante, gameFee)
	if err != nil {
		t.Fatalf("unexpected Reconcile error: %s", err)
	}

	_, err = ownerClient.WaitMined(ctx, tx)
	if err != nil {
		t.Fatalf("unexpected WaitMined error: %s", err)
	}

	//==========================================================================

	// Still as an Owner, check the loser's balance.
	callOpts, err := ownerClient.NewCallOpts(ctx)
	if err != nil {
		t.Fatalf("unexpected NewCallOpts error: %s", err)
	}

	amountPlayer2, err := contract.PlayerBalance(callOpts, common.HexToAddress(Player2Address))
	if err != nil {
		t.Fatalf("unexpected PlayerBalance error: %s", err)
	}

	expectedLoserBalance := big.NewInt(20000000000000000)
	if amountPlayer2.Cmp(expectedLoserBalance) != 0 {
		t.Fatalf("expecting losers's balance %d; got %d", expectedLoserBalance, amountPlayer2)
	}

	//==========================================================================

	// Still as an Owner, check the winner's balance.
	callOpts, err = ownerClient.NewCallOpts(ctx)
	if err != nil {
		t.Fatalf("unexpected NewCallOpts error: %s", err)
	}

	amountPlayer1, err := contract.PlayerBalance(callOpts, common.HexToAddress(Player1Address))
	if err != nil {
		t.Fatalf("unexpected PlayerBalance error: %s", err)
	}

	expectedWinnerBalance := big.NewInt(50000000000000000)
	if amountPlayer1.Cmp(expectedWinnerBalance) != 0 {
		t.Fatalf("expecting winner's balance %d; got %d", expectedWinnerBalance, amountPlayer1)
	}

	//==========================================================================

	// Still as an Owner, check the winner's balance.
	callOpts, err = ownerClient.NewCallOpts(ctx)
	if err != nil {
		t.Fatalf("unexpected NewCallOpts error: %s", err)
	}

	owner, err := contract.PlayerBalance(callOpts, ownerClient.Account)
	if err != nil {
		t.Fatalf("unexpected PlayerBalance error: %s", err)
	}

	expectedOwnerBalance := big.NewInt(10000000000000000)
	if owner.Cmp(expectedOwnerBalance) != 0 {
		t.Fatalf("expecting owner's balance %d; got %d", expectedOwnerBalance, owner)
	}
}
