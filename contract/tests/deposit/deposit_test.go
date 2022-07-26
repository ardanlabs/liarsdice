package tests

import (
	"context"
	"math"
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

	const gasLimit = 3000000
	const valueGwei = 0
	tranOpts, err := client.NewTransactOpts(ctx, gasLimit, valueGwei)
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

func TestDeposit(t *testing.T) {
	err := setup()
	if err != nil {
		t.Fatalf("unexpected setup error: %s", err)
	}

	// Connect as Player1.
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
}
