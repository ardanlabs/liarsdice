package bank_test

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	scbank "github.com/ardanlabs/liarsdice/business/contract/go/bank"
	"github.com/ardanlabs/liarsdice/business/core/bank"
	"github.com/ardanlabs/liarsdice/foundation/smart/contract"
	"github.com/ardanlabs/liarsdice/foundation/smart/currency"
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
	contractID, err := deployContract()
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Need a converter for handling ETH to USD to ETH conversions.
	converter := currency.NewDefaultConverter()

	// Connect player 1 to the smart contract.
	playerClient, err := bank.New(ctx, nil, contract.NetworkHTTPLocalhost, Player1KeyPath, Player1PassPhrase, contractID)
	if err != nil {
		t.Fatalf("error creating new bank for player: %s", err)
	}

	// Deposit ~10 USD into the players account.
	depositGWei := converter.USD2GWei(big.NewFloat(10))
	if _, _, err := playerClient.Deposit(ctx, depositGWei); err != nil {
		t.Fatalf("error making deposit: %s", err)
	}

	// Check the players balance in the smart contract.
	playerBalance, err := playerClient.Balance(ctx)
	if err != nil {
		t.Fatalf("error getting the player balance: %s", err)
	}

	// The players balance should match the deposit.
	depositGWei32, _ := depositGWei.Float32()
	playerBalance32, _ := playerBalance.Float32()
	if playerBalance32 != depositGWei32 {
		t.Fatalf("expecting balance to be %f; got %f", depositGWei32, playerBalance32)
	}

	// Perform a second deposit of the same amount.
	if _, _, err := playerClient.Deposit(ctx, depositGWei); err != nil {
		t.Fatalf("error making deposit: %s", err)
	}

	// Check the players balance in the smart contract.
	playerBalance, err = playerClient.Balance(ctx)
	if err != nil {
		t.Fatalf("error getting the player balance: %s", err)
	}

	// The players balance should match the two deposits.
	expectedGWei32, _ := depositGWei.Mul(depositGWei, big.NewFloat(2)).Float32()
	playerBalance32, _ = playerBalance.Float32()
	if playerBalance32 != expectedGWei32 {
		t.Fatalf("expecting balance to be %f; got %f", expectedGWei32, playerBalance32)
	}
}

func Test_Withdraw(t *testing.T) {
	contractID, err := deployContract()
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Need a converter for handling ETH to USD to ETH conversions.
	converter := currency.NewDefaultConverter()

	// Connect player 1 to the smart contract.
	playerClient, err := bank.New(ctx, nil, contract.NetworkHTTPLocalhost, Player1KeyPath, Player1PassPhrase, contractID)
	if err != nil {
		t.Fatalf("error creating new bank for owner: %s", err)
	}

	// =========================================================================
	// Deposit process

	// Get the starting balance.
	startingBalance, err := playerClient.WalletBalance(ctx)
	if err != nil {
		t.Fatalf("error getting player's wallet balance: %s", err)
	}

	// Perform a deposit from the player's wallet.
	depositGWeiAmount := converter.USD2GWei(big.NewFloat(10))
	depositTx, depositReceipt, err := playerClient.Deposit(ctx, depositGWeiAmount)
	if err != nil {
		t.Fatalf("error making deposit: %s", err)
	}

	// Calculate the expected balance by subtracting the amount deposited and the
	// gas fees for the transaction.
	gasCost := big.NewInt(0).Mul(depositTx.GasPrice(), big.NewInt(0).SetUint64(depositReceipt.GasUsed))
	depositWeiAmount := currency.GWei2Wei(depositGWeiAmount)
	expectedBalance := big.NewInt(0).Sub(startingBalance, depositWeiAmount)
	expectedBalance.Sub(expectedBalance, gasCost)

	// Get the updated wallet balance.
	currentBalance, err := playerClient.WalletBalance(ctx)
	if err != nil {
		t.Fatalf("error getting player's wallet balance: %s", err)
	}

	// The player's wallet balance should match the deposit minus the fees.
	if expectedBalance.Cmp(currentBalance) != 0 {
		t.Fatalf("expecting final balance to be %d; got %d", expectedBalance, currentBalance)
	}

	// =========================================================================
	// Withdraw process

	// Perform a withdraw to the player's wallet.
	withdrawTx, withdrawReceipt, err := playerClient.Withdraw(ctx)
	if err != nil {
		t.Fatalf("error calling withdraw: %s", err)
	}

	// Calculate the expected balance by adding the amount withdrawn and the
	// gas fees for the transaction.
	gasCost = big.NewInt(0).Mul(withdrawTx.GasPrice(), big.NewInt(0).SetUint64(withdrawReceipt.GasUsed))
	expectedBalance = big.NewInt(0).Add(currentBalance, depositWeiAmount)
	expectedBalance.Sub(expectedBalance, gasCost)

	// Get the updated wallet balance.
	currentBalance, err = playerClient.WalletBalance(ctx)
	if err != nil {
		t.Fatalf("error getting player's wallet balance: %s", err)
	}

	// The player's wallet balance should match the withdrawal minus the fees.
	if expectedBalance.Cmp(currentBalance) != 0 {
		t.Fatalf("expecting final balance to be %d; got %d", expectedBalance, currentBalance)
	}
}

func Test_WithdrawWithoutBalance(t *testing.T) {
	contractID, err := deployContract()
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect player 1 to the smart contract.
	playerClient, err := bank.New(ctx, nil, contract.NetworkHTTPLocalhost, Player1KeyPath, Player1PassPhrase, contractID)
	if err != nil {
		t.Fatalf("error creating new bank for owner: %s", err)
	}

	// Perform a withdraw to the player's wallet.
	if _, _, err := playerClient.Withdraw(ctx); err == nil {
		t.Fatal("expecting error when trying to withdraw from an empty balance")
	}
}

func Test_Reconcile(t *testing.T) {
	contractID, err := deployContract()
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Need a converter for handling ETH to USD to ETH conversions.
	converter := currency.NewDefaultConverter()

	// Connect owner to the smart contract.
	ownerClient, err := bank.New(ctx, nil, contract.NetworkHTTPLocalhost, OwnerKeyPath, OwnerPassPhrase, contractID)
	if err != nil {
		t.Fatalf("error creating new bank for owner: %s", err)
	}

	// Connect player 1 to the smart contract.
	player1Client, err := bank.New(ctx, nil, contract.NetworkHTTPLocalhost, Player1KeyPath, Player1PassPhrase, contractID)
	if err != nil {
		t.Fatalf("error creating new bank for player 1: %s", err)
	}

	// Connect player 2 to the smart contract.
	player2Client, err := bank.New(ctx, nil, contract.NetworkHTTPLocalhost, Player2KeyPath, Player2PassPhrase, contractID)
	if err != nil {
		t.Fatalf("error creating new bank for player 2: %s", err)
	}

	// Deposit ~$10 USD into the players account.
	player1DepositGWei := converter.USD2GWei(big.NewFloat(10))
	if _, _, err := player1Client.Deposit(ctx, player1DepositGWei); err != nil {
		t.Fatalf("error making deposit player 1: %s", err)
	}

	// Deposit ~$20 USD into the players account.
	player2DepositGWei := converter.USD2GWei(big.NewFloat(20))
	if _, _, err := player2Client.Deposit(ctx, player2DepositGWei); err != nil {
		t.Fatalf("error making deposit for player 2: %s", err)
	}

	// Set the ante and fees.
	anteGwei := converter.USD2GWei(big.NewFloat(5))
	feeGwei := converter.USD2GWei(big.NewFloat(5))

	// Reconcile with player 1 as the winner and player 2 as the loser.
	tx, receipt, err := ownerClient.Reconcile(ctx, Player1Address, []string{Player2Address}, anteGwei, feeGwei)
	if err != nil {
		t.Fatalf("error calling Reconcile: %s", err)
	}

	// Log the results of the reconcile transaction.
	t.Log(converter.FmtTransaction(tx))
	t.Log(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	// Capture player 1 balance in the smart contract.
	player1Balance, err := player1Client.Balance(ctx)
	if err != nil {
		t.Fatalf("error calling balance for player 1: %s", err)
	}

	// The winner should have $15 USD.
	winnerBalanceGWei32, _ := converter.USD2GWei(big.NewFloat(15)).Float32()
	player1Balance32, _ := player1Balance.Float32()
	if player1Balance32 != winnerBalanceGWei32 {
		t.Fatalf("expecting winner player balance to be %f; got %f", winnerBalanceGWei32, player1Balance32)
	}

	// Capture player 2 balance in the smart contract.
	player2Balance, err := player2Client.Balance(ctx)
	if err != nil {
		t.Fatalf("error calling balance for player 2: %s", err)
	}

	// The loser should have $15 USD.
	losingBalanceGWei32, _ := converter.USD2GWei(big.NewFloat(15)).Float32()
	player2Balance32, _ := player2Balance.Float32()
	if player2Balance32 != losingBalanceGWei32 {
		t.Fatalf("expecting loser player balance to be %f; got %f", losingBalanceGWei32, player2Balance32)
	}

	// Capture owber balance in the smart contract.
	ownerBalance, err := ownerClient.Balance(ctx)
	if err != nil {
		t.Fatalf("error calling balance for owner: %s", err)
	}

	// The owner should have $5 USD.
	feeGwei32, _ := feeGwei.Float32()
	ownerBalance32, _ := ownerBalance.Float32()
	if ownerBalance32 != feeGwei32 {
		t.Fatalf("expecting owner balance to be %f; got %f", feeGwei32, ownerBalance32)
	}
}

// =============================================================================

func deployContract() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("*** Deploying Contract ***")

	contractID, err := smartContract(ctx)
	if err != nil {
		fmt.Println("error deploying a new contract:", err)
		return "", err
	}

	return contractID, nil
}

func smartContract(ctx context.Context) (string, error) {
	client, err := contract.NewClient(ctx, contract.NetworkHTTPLocalhost, OwnerKeyPath, OwnerPassPhrase)
	if err != nil {
		return "", err
	}

	tranOpts, err := client.NewTransactOpts(ctx, 3_000_000, big.NewFloat(0))
	if err != nil {
		return "", err
	}

	address, tx, _, err := scbank.DeployBank(tranOpts, client.ContractBackend())
	if err != nil {
		return "", err
	}

	if _, err := client.WaitMined(ctx, tx); err != nil {
		return "", err
	}

	return address.String(), nil
}
