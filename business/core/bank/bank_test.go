package bank

import (
	"context"
	"math/big"
	"testing"

	"github.com/ardanlabs/liarsdice/contract/sol/go/contract"
	"github.com/ardanlabs/liarsdice/foundation/smartcontract/smart"
	"github.com/ethereum/go-ethereum/common"
)

const (
	PrimaryKeyPath    = "../../../zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	PrimaryPassPhrase = "123"

	Player1Address    = "0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7"
	Player1KeyPath    = "../../../zarf/ethereum/keystore/UTC--2022-05-13T16-59-42.277071000Z--0070742ff6003c3e809e78d524f0fe5dcc5ba7f7"
	Player1PassPhrase = "123"
)

func TestPlayerBalance(t *testing.T) {
	ctx := context.Background()

	// Deploy a new contract.
	contractID, err := deployContract(ctx)
	if err != nil {
		t.Fatalf("error deploying a new contract: %s", err)
	}

	// Create a bank for the contract owner.
	ownerClient, err := New(ctx, smart.NetworkHTTPLocalhost, PrimaryKeyPath, PrimaryPassPhrase, contractID)
	if err != nil {
		t.Fatalf("error creating new bank for owner: %s", err)
	}

	// Create a bank for the player.
	playerClient, err := New(ctx, smart.NetworkHTTPLocalhost, Player1KeyPath, Player1PassPhrase, contractID)
	if err != nil {
		t.Fatalf("error creating new bank for player: %s", err)
	}

	// =========================================================================
	// Make a deposit as player.

	tranOpts, err := playerClient.client.NewTransactOpts(ctx, 300000, 40000000)
	if err != nil {
		t.Fatalf("error creating player transaction: %s", err)
	}

	tx, err := playerClient.contract.Deposit(tranOpts)
	if err != nil {
		t.Fatalf("error making deposit: %s", err)
	}

	_, err = playerClient.client.WaitMined(ctx, tx)
	if err != nil {
		t.Fatalf("error mining the block: %s", err)
	}

	// =========================================================================
	// Check the player balance as an owner.

	callOpts, err := ownerClient.client.NewCallOpts(ctx)
	if err != nil {
		t.Fatalf("error creating owner transaction: %s", err)
	}

	amount, err := ownerClient.contract.PlayerBalance(callOpts, common.HexToAddress(Player1Address))
	if err != nil {
		t.Fatalf("error getting the player balance: %s", err)
	}

	expectedWeiAmount := big.NewInt(40000000000000000)
	if amount.Cmp(expectedWeiAmount) != 0 {
		t.Fatalf("expecting balance to be %d; got %d", expectedWeiAmount, amount)
	}

	// ==========================================================================
	// Make a new player deposit.

	tranOpts, err = playerClient.client.NewTransactOpts(ctx, 300000, 40000000)
	if err != nil {
		t.Fatalf("error creating player transaction: %s", err)
	}

	tx, err = playerClient.contract.Deposit(tranOpts)
	if err != nil {
		t.Fatalf("error making deposit: %s", err)
	}

	_, err = playerClient.client.WaitMined(ctx, tx)
	if err != nil {
		t.Fatalf("error mining the block: %s", err)
	}

	// =========================================================================
	// Check the player balance as an owner.

	amount, err = ownerClient.contract.PlayerBalance(callOpts, common.HexToAddress(Player1Address))
	if err != nil {
		t.Fatalf("error getting the player balance: %s", err)
	}

	expectedWeiAmount = big.NewInt(80000000000000000)
	if amount.Cmp(expectedWeiAmount) != 0 {
		t.Fatalf("expecting balance to be %d; got %d", expectedWeiAmount, amount)
	}
}

// ============================================================================

func deployContract(ctx context.Context) (string, error) {
	// Create a client as an owner.
	client, err := smart.Connect(ctx, smart.NetworkHTTPLocalhost, PrimaryKeyPath, PrimaryPassPhrase)
	if err != nil {
		return "", err
	}

	tranOpts, err := client.NewTransactOpts(ctx, 3_000_000, 0)
	if err != nil {
		return "", err
	}

	address, tx, _, err := contract.DeployContract(tranOpts, client.ContractBackend())
	if err != nil {
		return "", err
	}

	_, err = client.WaitMined(ctx, tx)
	if err != nil {
		return "", err
	}

	return address.String(), nil
}
