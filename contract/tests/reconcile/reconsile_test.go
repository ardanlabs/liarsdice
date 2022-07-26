package tests

import (
	"context"
	"math/big"
	"testing"

	ldc "github.com/ardanlabs/liarsdice/contract/sol/go"
	"github.com/ardanlabs/liarsdice/foundation/smartcontract/smart"
	"github.com/ethereum/go-ethereum/common"
)

const (
	PrimaryKeyPath    = "../UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	PrimaryPassPhrase = "123"

	Player1Address    = "0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7"
	Player1KeyPath    = "../UTC--2022-05-13T16-59-42.277071000Z--0070742ff6003c3e809e78d524f0fe5dcc5ba7f7"
	Player1PassPhrase = "123"

	Player2Address    = "0x8e113078adf6888b7ba84967f299f29aece24c55"
	Player2KeyPath    = "../UTC--2022-05-13T16-57-20.203544000Z--8e113078adf6888b7ba84967f299f29aece24c55"
	Player2PassPhrase = "123"
)

var contract *ldc.Contract
var ctx context.Context

func setup() error {
	ctx = context.Background()
	var err error

	client, err := smart.Connect(ctx, smart.NetworkHTTPLocalhost, PrimaryKeyPath, PrimaryPassPhrase)
	if err != nil {
		return err
	}

	tranOpts, err := client.NewTransactOpts(ctx, 3000000, 0)
	if err != nil {
		return err
	}

	address, _, _, err := ldc.DeployContract(tranOpts, client.ContractBackend())
	if err != nil {
		return err
	}

	contract, err = ldc.NewContract(address, client.ContractBackend())
	if err != nil {
		return err
	}

	return nil
}

func TestReconcile(t *testing.T) {
	err := setup()
	if err != nil {
		t.Fatalf("unexpected setup error: %s", err)
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
