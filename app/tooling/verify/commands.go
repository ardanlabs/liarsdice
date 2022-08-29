package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ardanlabs/liarsdice/business/core/bank"
	"github.com/ardanlabs/liarsdice/foundation/smart/contract"
	"github.com/ardanlabs/liarsdice/foundation/smart/currency"
	"github.com/ethereum/go-ethereum/common"
)

func balance(ctx context.Context, network string, address string, contractID string) error {
	_, err := bank.New(ctx, network, keyPath, passPhrase, contractID)
	if err != nil {
		return err
	}

	return nil
}

func txHash(ctx context.Context, network string, hash string) error {
	converter, err := currency.NewConverter(coinMarketCapKey)
	if err != nil {
		log.Println("unable to create converter, using default:", err)
		converter = currency.NewDefaultConverter()
	}

	client, err := contract.NewClient(ctx, network, keyPath, passPhrase)
	if err != nil {
		return err
	}

	oneETHToUSD, oneUSDToETH := converter.Values()

	fmt.Println("network         :", network)
	fmt.Println("fromAddress     :", client.Account())
	fmt.Println("oneETHToUSD     :", oneETHToUSD)
	fmt.Println("oneUSDToETH     :", oneUSDToETH)

	// =========================================================================

	fmt.Println("hash to check   :", hash)

	txHash := common.HexToHash(hash)
	tx, pending, err := client.TransactionByHash(ctx, txHash)
	if err != nil {
		log.Print("status            : Not Found")
		return err
	}

	if pending {
		log.Print("status            : Pending")
		return nil
	}
	log.Print("status            : Completed")
	fmt.Print(converter.FmtTransaction(tx))

	receipt, err := client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return err
	}

	fmt.Println("Output          :")
	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	return nil
}
