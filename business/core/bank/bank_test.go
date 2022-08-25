package bank_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ardanlabs/liarsdice/business/core/bank"
	"github.com/ardanlabs/liarsdice/contract/sol/go/contract"
	"github.com/ardanlabs/liarsdice/foundation/smartcontract/smart"
)

const (
	OwnerAddress    = "0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd"
	OwnerKeyPath    = "../../../zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	OwnerPassPhrase = "123"

	Player1Address    = "0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7"
	Player1KeyPath    = "../../../zarf/ethereum/keystore/UTC--2022-05-13T16-59-42.277071000Z--0070742ff6003c3e809e78d524f0fe5dcc5ba7f7"
	Player1PassPhrase = "123"

	Player2Address    = "0x8e113078adf6888b7ba84967f299f29aece24c55"
	Player2KeyPath    = "../../../zarf/ethereum/keystore/UTC--2022-05-13T16-57-20.203544000Z--8e113078adf6888b7ba84967f299f29aece24c55"
	Player2PassPhrase = "123"
)

func Test_PlayerBalance(t *testing.T) {
	ctx := context.Background()

	contractID, err := deployContract(ctx)
	if err != nil {
		t.Fatalf("error deploying a new contract: %s", err)
	}

	ownerClient, err := bank.New(ctx, smart.NetworkHTTPLocalhost, OwnerKeyPath, OwnerPassPhrase, contractID)
	if err != nil {
		t.Fatalf("error creating new bank for owner: %s", err)
	}

	playerClient, err := bank.New(ctx, smart.NetworkHTTPLocalhost, Player1KeyPath, Player1PassPhrase, contractID)
	if err != nil {
		t.Fatalf("error creating new bank for player: %s", err)
	}

	err = playerClient.Deposit(ctx, Player1Address, 40000000)
	if err != nil {
		t.Fatalf("error making deposit: %s", err)
	}

	amount, err := ownerClient.Balance(ctx, Player1Address)
	if err != nil {
		t.Fatalf("error getting the player balance: %s", err)
	}

	expectedWeiAmount := big.NewInt(40000000000000000)
	if amount.Cmp(expectedWeiAmount) != 0 {
		t.Fatalf("expecting balance to be %d; got %d", expectedWeiAmount, amount)
	}

	err = playerClient.Deposit(ctx, Player1Address, 40000000)
	if err != nil {
		t.Fatalf("error making deposit: %s", err)
	}

	amount, err = ownerClient.Balance(ctx, Player1Address)
	if err != nil {
		t.Fatalf("error getting the player balance: %s", err)
	}

	expectedWeiAmount = big.NewInt(80000000000000000)
	if amount.Cmp(expectedWeiAmount) != 0 {
		t.Fatalf("expecting balance to be %d; got %d", expectedWeiAmount, amount)
	}
}

func Test_Withdraw(t *testing.T) {
	ctx := context.Background()

	contractID, err := deployContract(ctx)
	if err != nil {
		t.Fatalf("error deploying a new contract: %s", err)
	}

	ownerClient, err := bank.New(ctx, smart.NetworkHTTPLocalhost, OwnerKeyPath, OwnerPassPhrase, contractID)
	if err != nil {
		t.Fatalf("error creating new bank for owner: %s", err)
	}

	playerClient, err := bank.New(ctx, smart.NetworkHTTPLocalhost, Player1KeyPath, Player1PassPhrase, contractID)
	if err != nil {
		t.Fatalf("error creating new bank for owner: %s", err)
	}

	err = playerClient.Deposit(ctx, Player1Address, 40000000)
	if err != nil {
		t.Fatalf("error making deposit: %s", err)
	}

	err = playerClient.Withdraw(ctx, Player1Address)
	if err != nil {
		t.Fatalf("error calling Withdraw: %s", err)
	}

	balance, err := ownerClient.Balance(ctx, Player1Address)
	if err != nil {
		t.Fatalf("error calling Balance: %s", err)
	}

	if balance.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("expecting player balance to be 0; got %d", balance)
	}

	// TODO: You need to check the Wallet Balance before and after this as well.
}

func Test_WithdrawWithoutBalance(t *testing.T) {
	ctx := context.Background()

	contractID, err := deployContract(ctx)
	if err != nil {
		t.Fatalf("error deploying a new contract: %s", err)
	}

	playerClient, err := bank.New(ctx, smart.NetworkHTTPLocalhost, Player1KeyPath, Player1PassPhrase, contractID)
	if err != nil {
		t.Fatalf("error creating new bank for owner: %s", err)
	}

	err = playerClient.Withdraw(ctx, Player1Address)
	if err == nil {
		t.Fatal("expecting error when trying to withdraw from an empty balance")
	}
}

func Test_Reconcile(t *testing.T) {
	ctx := context.Background()

	contractID, err := deployContract(ctx)
	if err != nil {
		t.Fatalf("error deploying a new contract: %s", err)
	}

	ownerClient, err := bank.New(ctx, smart.NetworkHTTPLocalhost, OwnerKeyPath, OwnerPassPhrase, contractID)
	if err != nil {
		t.Fatalf("error creating new bank for owner: %s", err)
	}

	player1Client, err := bank.New(ctx, smart.NetworkHTTPLocalhost, Player1KeyPath, Player1PassPhrase, contractID)
	if err != nil {
		t.Fatalf("error creating new bank for player 1: %s", err)
	}

	player2Client, err := bank.New(ctx, smart.NetworkHTTPLocalhost, Player2KeyPath, Player2PassPhrase, contractID)
	if err != nil {
		t.Fatalf("error creating new bank for player 2: %s", err)
	}

	err = player1Client.Deposit(ctx, Player1Address, 40000000)
	if err != nil {
		t.Fatalf("error making deposit player 1: %s", err)
	}

	err = player2Client.Deposit(ctx, Player2Address, 20000000)
	if err != nil {
		t.Fatalf("error making deposit for player 2: %s", err)
	}

	ante := big.NewInt(20000000000000000)
	fee := big.NewInt(10000000000000000)

	losingAccounts := []string{Player2Address}

	tx, receipt, err := ownerClient.Reconcile(ctx, Player1Address, losingAccounts, ante, fee)
	if err != nil {
		t.Fatalf("error calling Reconcile: %s", err)
	}

	t.Log(smart.FmtTransaction(tx))
	t.Log(smart.FmtTransactionReceipt(receipt, tx.GasPrice()))

	player1Balance, err := ownerClient.Balance(ctx, Player1Address)
	if err != nil {
		t.Fatalf("error calling balance for player 1: %s", err)
	}

	winnerBalance := big.NewInt(50000000000000000)

	if player1Balance.Cmp(winnerBalance) != 0 {
		t.Fatalf("expecting winner player balance to be %d; got %d", winnerBalance, player1Balance)
	}

	player2Balance, err := ownerClient.Balance(ctx, Player2Address)
	if err != nil {
		t.Fatalf("error calling balance for player 2: %s", err)
	}

	if player2Balance.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("expecting loser player balance to be %d; got %d", 0, player2Balance)
	}

	contractBalance, err := ownerClient.Balance(ctx, OwnerAddress)
	if err != nil {
		t.Fatalf("error calling balance for owner: %s", err)
	}

	if contractBalance.Cmp(fee) != 0 {
		t.Fatalf("expecting owner balance to be %d; got %d", fee, contractBalance)
	}
}

func deployContract(ctx context.Context) (string, error) {
	client, err := smart.Connect(ctx, smart.NetworkHTTPLocalhost, OwnerKeyPath, OwnerPassPhrase)
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

	if _, err := client.WaitMined(ctx, tx); err != nil {
		return "", err
	}

	return address.String(), nil
}
