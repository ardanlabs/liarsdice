// Package commands provides all the different command options and logic.
package commands

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	scbank "github.com/ardanlabs/liarsdice/business/contract/go/bank"
	"github.com/ardanlabs/liarsdice/business/core/bank"
	"github.com/ardanlabs/liarsdice/foundation/smart/contract"
	"github.com/ardanlabs/liarsdice/foundation/smart/currency"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

// Withdraw will remove any balance from the contract back into the calling account.
func Withdraw(ctx context.Context, converter currency.Converter, bank *bank.Bank, v Values) error {
	tx, receipt, err := bank.Withdraw(ctx)
	if err != nil {
		return err
	}

	fmt.Print(converter.FmtTransaction(tx))
	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	return nil
}

// Deploy will deploy the smart contract to the configured network.
func Deploy(ctx context.Context, converter currency.Converter, bank *bank.Bank, v Values) (err error) {
	client := bank.Client()

	startingBalance, err := client.CurrentBalance(ctx)
	if err != nil {
		return err
	}
	defer func() {
		endingBalance, dErr := client.CurrentBalance(ctx)
		if dErr != nil {
			err = dErr
			return
		}
		fmt.Print(converter.FmtBalanceSheet(startingBalance, endingBalance))
	}()

	// =========================================================================

	const gasLimit = 1600000
	const valueGwei = 0.0
	tranOpts, err := client.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		return err
	}

	// =========================================================================

	address, tx, _, err := scbank.DeployBank(tranOpts, client.ContractBackend())
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransaction(tx))

	fmt.Println("\nContract Details")
	fmt.Println("----------------------------------------------------")
	fmt.Println("contract id     :", address.Hex())
	fmt.Printf("export GAME_CONTRACT_ID=%s\n", address.Hex())

	// =========================================================================

	clientWait, err := contract.NewClient(ctx, v.Network, v.KeyFile, v.PassPhrase)
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

// Wallet returns the current wallet balance
func Wallet(ctx context.Context, converter currency.Converter, bank *bank.Bank, v Values) error {
	wei, err := bank.WalletBalance(ctx)
	if err != nil {
		return err
	}

	fmt.Println("\nWallet Balance")
	fmt.Println("----------------------------------------------------")
	fmt.Println("wei             :", wei)
	fmt.Println("gwei            :", currency.Wei2GWei(wei))
	fmt.Println("usd             :", converter.Wei2USD(wei))

	return nil
}

// Balance returns the current balance of the specified address.
func Balance(ctx context.Context, converter currency.Converter, bank *bank.Bank, v Values) error {
	gwei, err := bank.AccountBalance(ctx, v.Address)
	if err != nil {
		return err
	}

	fmt.Println("\nGame Balance")
	fmt.Println("----------------------------------------------------")
	fmt.Println("account         :", v.Address)
	fmt.Println("wei             :", currency.GWei2Wei(gwei))
	fmt.Println("gwei            :", gwei)
	fmt.Println("usd             :", converter.GWei2USD(gwei))

	return nil
}

// Transaction returns the transaction and receipt information for the specified
// transaction. The txHex is expected to be in a 0x format.
func Transaction(ctx context.Context, converter currency.Converter, bank *bank.Bank, v Values) error {
	txHash := common.HexToHash(v.Hex)
	tx, pending, err := bank.Client().TransactionByHash(ctx, txHash)
	if err != nil {
		return err
	}

	if pending {
		return errors.New("transaction pending")
	}

	fmt.Print(converter.FmtTransaction(tx))

	receipt, err := bank.Client().TransactionReceipt(ctx, txHash)
	if err != nil {
		return err
	}

	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	return nil
}
