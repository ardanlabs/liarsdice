// Package commands provides all the different command options and logic.
package commands

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	scbank "github.com/ardanlabs/liarsdice/business/contract/go/bank"
	"github.com/ardanlabs/liarsdice/business/core/bank"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

// Deposit will move money from the wallet into the game contract.
func Deposit(ctx context.Context, converter *currency.Converter, bank *bank.Bank, amountUSD float64) error {
	fmt.Println("\nDeposit Details")
	fmt.Println("----------------------------------------------------")
	fmt.Println("address         :", bank.Client().Address())
	fmt.Println("amount          :", amountUSD)

	amountGWei := converter.USD2GWei(big.NewFloat(amountUSD))
	tx, receipt, err := bank.Deposit(ctx, amountGWei)
	if err != nil {
		return err
	}

	fmt.Print(converter.FmtTransaction(tx))
	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	return nil
}

// Withdraw will remove money from the game contract back into the wallet.
func Withdraw(ctx context.Context, converter *currency.Converter, bank *bank.Bank) error {
	fmt.Println("\nWithdraw Details")
	fmt.Println("----------------------------------------------------")
	fmt.Println("address         :", bank.Client().Address())

	tx, receipt, err := bank.Withdraw(ctx)
	if err != nil {
		return err
	}

	fmt.Print(converter.FmtTransaction(tx))
	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	return nil
}

// Balance returns the current balance of the specified address.
func Balance(ctx context.Context, converter *currency.Converter, bank *bank.Bank, address string) error {
	fmt.Println("\nGame Balance")
	fmt.Println("----------------------------------------------------")
	fmt.Println("account         :", address)

	gwei, err := bank.AccountBalance(ctx, address)
	if err != nil {
		return err
	}

	fmt.Println("wei             :", currency.GWei2Wei(gwei))
	fmt.Println("gwei            :", gwei)
	fmt.Println("usd             :", converter.GWei2USD(gwei))

	return nil
}

// =============================================================================

// Deploy will deploy the smart contract to the configured network.
func Deploy(ctx context.Context, converter *currency.Converter, ethereum *ethereum.Ethereum) (err error) {
	startingBalance, err := ethereum.Balance(ctx)
	if err != nil {
		return err
	}
	defer func() {
		endingBalance, dErr := ethereum.Balance(ctx)
		if dErr != nil {
			err = dErr
			return
		}
		fmt.Print(converter.FmtBalanceSheet(startingBalance, endingBalance))
	}()

	// =========================================================================

	const gasLimit = 1700000
	const valueGwei = 0.0
	tranOpts, err := ethereum.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		return err
	}

	// =========================================================================

	address, tx, _, err := scbank.DeployBank(tranOpts, ethereum.RawClient())
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransaction(tx))

	fmt.Println("\nContract Details")
	fmt.Println("----------------------------------------------------")
	fmt.Println("contract id     :", address.Hex())
	fmt.Printf("export GAME_CONTRACT_ID=%s\n", address.Hex())

	// =========================================================================

	clientWait, err := ethereum.Copy(ctx)
	if err != nil {
		return err
	}

	fmt.Println("\nWaiting Logs")
	fmt.Println("----------------------------------------------------")
	log.Root().SetHandler(log.StdoutHandler)

	receipt, err := clientWait.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	return nil
}

// Wallet returns the current wallet balance for the specified address.
func Wallet(ctx context.Context, converter *currency.Converter, ethereum *ethereum.Ethereum, address string) error {
	fmt.Println("\nWallet Balance")
	fmt.Println("----------------------------------------------------")
	fmt.Println("account         :", address)

	wei, err := ethereum.BalanceAt(ctx, address)
	if err != nil {
		return err
	}

	fmt.Println("wei             :", wei)
	fmt.Println("gwei            :", currency.Wei2GWei(wei))
	fmt.Println("usd             :", converter.Wei2USD(wei))

	return nil
}

// Transaction returns the transaction and receipt information for the specified
// transaction. The txHex is expected to be in a 0x format.
func Transaction(ctx context.Context, converter *currency.Converter, ethereum *ethereum.Ethereum, tranID string) error {
	fmt.Println("\nTransaction ID")
	fmt.Println("----------------------------------------------------")
	fmt.Println("tran id         :", tranID)

	txHash := common.HexToHash(tranID)
	tx, pending, err := ethereum.TransactionByHash(ctx, txHash)
	if err != nil {
		return err
	}

	if pending {
		return errors.New("transaction pending")
	}

	fmt.Print(converter.FmtTransaction(tx))

	receipt, err := ethereum.TransactionReceipt(ctx, txHash)
	if err != nil {
		return err
	}

	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	return nil
}
