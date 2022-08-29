// Package commands provides all the different command options and logic.
package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/ardanlabs/liarsdice/business/core/bank"
	"github.com/ardanlabs/liarsdice/foundation/smart/currency"
	"github.com/ethereum/go-ethereum/common"
)

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
