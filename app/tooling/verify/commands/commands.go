// Package commands provides all the different command options and logic.
package commands

import (
	"context"
	"fmt"

	"github.com/ardanlabs/liarsdice/business/core/bank"
	"github.com/ardanlabs/liarsdice/foundation/smart/contract"
	"github.com/ardanlabs/liarsdice/foundation/smart/currency"
	"github.com/ethereum/go-ethereum/common"
)

// Balance returns the current balance of the specified address.
func Balance(ctx context.Context, converter currency.Converter, v Values) error {
	b, err := bank.New(ctx, v.Network, v.KeyFile, v.PassPhrase, v.ContractID)
	if err != nil {
		return err
	}

	gwei, err := b.AccountBalance(ctx, v.Address)
	if err != nil {
		return err
	}

	fmt.Println("Address         :", v.Address)
	fmt.Println("WEI             :", currency.GWei2Wei(gwei))
	fmt.Println("GWEI            :", gwei)
	fmt.Println("USD             :", converter.GWei2USD(gwei))

	return nil
}

// Transaction returns the transaction and receipt information for the specified
// transaction. The txHex is expected to be in a 0x format.
func Transaction(ctx context.Context, converter currency.Converter, v Values) error {
	client, err := contract.NewClient(ctx, v.Network, v.KeyFile, v.PassPhrase)
	if err != nil {
		return err
	}

	fmt.Println("fromAddress     :", client.Account())
	fmt.Println("transaction     :", v.Hex)

	txHash := common.HexToHash(v.Hex)
	tx, pending, err := client.TransactionByHash(ctx, txHash)
	if err != nil {
		fmt.Println("tx status       : Not Found")
		return err
	}

	if pending {
		fmt.Println("tx status       : Pending")
		return nil
	}

	fmt.Print(converter.FmtTransaction(tx))

	receipt, err := client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return err
	}

	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	return nil
}
