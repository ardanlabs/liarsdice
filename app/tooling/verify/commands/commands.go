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

// Harded this here for now just to make life easier.
const (
	keyPath          = "zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	passPhrase       = "123"
	coinMarketCapKey = "a8cd12fb-d056-423f-877b-659046af0aa5"
)

func Balance(ctx context.Context, network string, address string, contractID string) error {
	b, err := bank.New(ctx, network, keyPath, passPhrase, contractID)
	if err != nil {
		return err
	}

	gwei, err := b.AccountBalance(ctx, address)
	if err != nil {
		return err
	}

	converter, err := currency.NewConverter(coinMarketCapKey)
	if err != nil {
		converter = currency.NewDefaultConverter()
	}

	fmt.Println("Balances        :", address)
	fmt.Println("WEI             :", currency.GWei2Wei(gwei))
	fmt.Println("GWEI            :", gwei)
	fmt.Println("USD             :", converter.GWei2USD(gwei))

	return nil
}

func TXHash(ctx context.Context, network string, hash string) error {
	converter, err := currency.NewConverter(coinMarketCapKey)
	if err != nil {
		converter = currency.NewDefaultConverter()
	}

	client, err := contract.NewClient(ctx, network, keyPath, passPhrase)
	if err != nil {
		return err
	}

	fmt.Println("fromAddress     :", client.Account())
	fmt.Println("hash to check   :", hash)

	txHash := common.HexToHash(hash)
	tx, pending, err := client.TransactionByHash(ctx, txHash)
	if err != nil {
		fmt.Print("status            : Not Found")
		return err
	}

	if pending {
		fmt.Print("status            : Pending")
		return nil
	}
	fmt.Print("status            : Completed")
	fmt.Print(converter.FmtTransaction(tx))

	receipt, err := client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return err
	}

	fmt.Println("Output          :")
	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	return nil
}
