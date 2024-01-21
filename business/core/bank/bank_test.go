package bank_test

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	scbank "github.com/ardanlabs/liarsdice/business/contract/go/bank"
	"github.com/ardanlabs/liarsdice/business/core/bank"
	"github.com/ethereum/go-ethereum/common"
)

var (
	backend    *ethereum.SimulatedBackend
	ownerClt   *ethereum.Client
	player1Clt *ethereum.Client
	player2Clt *ethereum.Client
)

func TestMain(m *testing.M) {
	var err error
	backend, err = ethereum.CreateSimulatedBackend(3, true, big.NewInt(100))
	if err != nil {
		fmt.Println("create backend", err)
		os.Exit(1)
	}
	defer backend.Close()

	ownerClt, err = ethereum.NewClient(backend, backend.PrivateKeys[0])
	if err != nil {
		fmt.Println("create ownerClt client", err)
		os.Exit(1)
	}

	player1Clt, err = ethereum.NewClient(backend, backend.PrivateKeys[1])
	if err != nil {
		fmt.Println("create player1Clt client", err)
		os.Exit(1)
	}

	player2Clt, err = ethereum.NewClient(backend, backend.PrivateKeys[2])
	if err != nil {
		fmt.Println("create player2Clt client", err)
		os.Exit(1)
	}

	m.Run()
}

func Test_PlayerBalance(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	contractID, err := deployContract()
	if err != nil {
		t.Fatal(err)
	}

	converter := currency.NewDefaultConverter(scbank.BankMetaData.ABI)

	// Connect player 1 to the smart contract.
	playerBank, err := bank.New(ctx, nil, backend, player1Clt.PrivateKey(), contractID)
	if err != nil {
		t.Fatalf("error creating new bank for player: %s", err)
	}

	// Deposit ~10 USD into the players account.
	depositGWei := converter.USD2GWei(big.NewFloat(10))
	if _, _, err := playerBank.Deposit(ctx, depositGWei); err != nil {
		t.Fatalf("error making deposit: %s", err)
	}

	// Check the players balance in the smart contract.
	playerBalance, err := playerBank.Balance(ctx)
	if err != nil {
		t.Fatalf("error getting the player balance: %s", err)
	}

	// The players balance should match the deposit.
	got, _ := playerBalance.Float32()
	exp, _ := depositGWei.Float32()
	if got != exp {
		t.Fatalf("expecting balance to be %f; got %f", exp, got)
	}

	// Perform a second deposit of the same amount.
	if _, _, err := playerBank.Deposit(ctx, depositGWei); err != nil {
		t.Fatalf("error making deposit: %s", err)
	}

	// Check the players balance in the smart contract.
	playerBalance, err = playerBank.Balance(ctx)
	if err != nil {
		t.Fatalf("error getting the player balance: %s", err)
	}

	// The players balance should match the two deposits.
	got, _ = playerBalance.Float32()
	exp, _ = depositGWei.Mul(depositGWei, big.NewFloat(2)).Float32()
	if got != exp {
		t.Fatalf("expecting balance to be %f; got %f", exp, got)
	}
}

func Test_Withdraw(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	contractID, err := deployContract()
	if err != nil {
		t.Fatal(err)
	}

	converter := currency.NewDefaultConverter(scbank.BankMetaData.ABI)

	// Connect player 1 to the smart contract.
	playerBank, err := bank.New(ctx, nil, backend, player1Clt.PrivateKey(), contractID)
	if err != nil {
		t.Fatalf("error creating new bank for owner: %s", err)
	}

	// =========================================================================
	// Deposit process

	// Get the starting balance.
	startingBalance, err := playerBank.EthereumBalance(ctx)
	if err != nil {
		t.Fatalf("error getting player's ethereum balance: %s", err)
	}

	// Perform a deposit from the player's wallet.
	depositGWeiAmount := converter.USD2GWei(big.NewFloat(10))
	depositTx, depositReceipt, err := playerBank.Deposit(ctx, depositGWeiAmount)
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
	currentBalance, err := playerBank.EthereumBalance(ctx)
	if err != nil {
		t.Fatalf("error getting player's wallet balance: %s", err)
	}

	// The player's wallet balance should match the deposit minus the fees.
	if currentBalance.Cmp(expectedBalance) != 0 {
		t.Fatalf("expecting final balance to be %d; got %d", expectedBalance, currentBalance)
	}

	// =========================================================================
	// Withdraw process

	// Perform a withdraw to the player's wallet.
	withdrawTx, withdrawReceipt, err := playerBank.Withdraw(ctx)
	if err != nil {
		t.Fatalf("error calling withdraw: %s", err)
	}

	// Calculate the expected balance by adding the amount withdrawn and the
	// gas fees for the transaction.
	gasCost = big.NewInt(0).Mul(withdrawTx.GasPrice(), big.NewInt(0).SetUint64(withdrawReceipt.GasUsed))
	expectedBalance = big.NewInt(0).Add(currentBalance, depositWeiAmount)
	expectedBalance.Sub(expectedBalance, gasCost)

	// Get the updated wallet balance.
	currentBalance, err = playerBank.EthereumBalance(ctx)
	if err != nil {
		t.Fatalf("error getting player's wallet balance: %s", err)
	}

	// The player's wallet balance should match the withdrawal minus the fees.
	if currentBalance.Cmp(expectedBalance) != 0 {
		t.Fatalf("expecting final balance to be %d; got %d", expectedBalance, currentBalance)
	}
}

func Test_WithdrawWithoutBalance(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	contractID, err := deployContract()
	if err != nil {
		t.Fatal(err)
	}

	// Connect player 1 to the smart contract.
	playerBank, err := bank.New(ctx, nil, backend, player1Clt.PrivateKey(), contractID)
	if err != nil {
		t.Fatalf("error creating new bank for owner: %s", err)
	}

	// Perform a withdraw to the player's wallet.
	if _, _, err := playerBank.Withdraw(ctx); err == nil {
		t.Fatal("expecting error when trying to withdraw from an empty balance")
	}
}

func Test_Reconcile(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	contractID, err := deployContract()
	if err != nil {
		t.Fatal(err)
	}

	// Need a converter for handling ETH to USD to ETH conversions.
	converter := currency.NewDefaultConverter(scbank.BankMetaData.ABI)

	// Connect owner to the smart contract.
	ownerBank, err := bank.New(ctx, nil, backend, ownerClt.PrivateKey(), contractID)
	if err != nil {
		t.Fatalf("error creating new bank for owner: %s", err)
	}

	// Connect player 1 to the smart contract.
	player1Bank, err := bank.New(ctx, nil, backend, player1Clt.PrivateKey(), contractID)
	if err != nil {
		t.Fatalf("error creating new bank for player 1: %s", err)
	}

	// Connect player 2 to the smart contract.
	player2Bank, err := bank.New(ctx, nil, backend, player2Clt.PrivateKey(), contractID)
	if err != nil {
		t.Fatalf("error creating new bank for player 2: %s", err)
	}

	// Deposit ~$10 USD into the players account.
	player1DepositGWei := converter.USD2GWei(big.NewFloat(10))
	if _, _, err := player1Bank.Deposit(ctx, player1DepositGWei); err != nil {
		t.Fatalf("error making deposit player 1: %s", err)
	}

	// Deposit ~$20 USD into the players account.
	player2DepositGWei := converter.USD2GWei(big.NewFloat(20))
	if _, _, err := player2Bank.Deposit(ctx, player2DepositGWei); err != nil {
		t.Fatalf("error making deposit for player 2: %s", err)
	}

	// Set the ante and fees.
	anteGwei := converter.USD2GWei(big.NewFloat(5))
	feeGwei := converter.USD2GWei(big.NewFloat(5))

	// Reconcile with player 1 as the winner and player 2 as the loser.
	_, _, err = ownerBank.Reconcile(ctx, player1Clt.Address(), []common.Address{player2Clt.Address()}, anteGwei, feeGwei)
	if err != nil {
		t.Fatalf("error calling Reconcile: %s", err)
	}

	// Capture player 1 balance in the smart contract.
	player1Balance, err := player1Bank.Balance(ctx)
	if err != nil {
		t.Fatalf("error calling balance for player 1: %s", err)
	}

	// The winner should have $15 USD.
	got, _ := player1Balance.Float32()
	exp, _ := converter.USD2GWei(big.NewFloat(15)).Float32()
	if got != exp {
		t.Fatalf("expecting winner player balance to be %f; got %f", exp, got)
	}

	// Capture player 2 balance in the smart contract.
	player2Balance, err := player2Bank.Balance(ctx)
	if err != nil {
		t.Fatalf("error calling balance for player 2: %s", err)
	}

	// The loser should have $15 USD.
	got, _ = player2Balance.Float32()
	exp, _ = converter.USD2GWei(big.NewFloat(15)).Float32()
	if got != exp {
		t.Fatalf("expecting loser player balance to be %v; got %v", exp, got)
	}

	// Capture owber balance in the smart contract.
	ownerBalance, err := ownerBank.Balance(ctx)
	if err != nil {
		t.Fatalf("error calling balance for owner: %s", err)
	}

	// The owner should have $5 USD.
	got, _ = ownerBalance.Float32()
	exp, _ = feeGwei.Float32()
	if got != exp {
		t.Fatalf("expecting owner balance to be %f; got %f", exp, got)
	}
}

// =============================================================================

func deployContract() (common.Address, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("Deploying Contract ...")
	defer fmt.Println("Deployed")

	contractID, err := smartContract(ctx)
	if err != nil {
		fmt.Println("error deploying a new contract:", err)

		var empty common.Address
		return empty, err
	}

	return contractID, nil
}

func smartContract(ctx context.Context) (common.Address, error) {
	var empty common.Address

	tranOpts, err := ownerClt.NewTransactOpts(ctx, 10_000_000, big.NewInt(0), big.NewFloat(0))
	if err != nil {
		return empty, err
	}

	address, tx, _, err := scbank.DeployBank(tranOpts, ownerClt.Backend)
	if err != nil {
		return empty, err
	}

	if _, err := ownerClt.WaitMined(ctx, tx); err != nil {
		return empty, err
	}

	return address, nil
}
