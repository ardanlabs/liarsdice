package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ardanlabs/liarsdice/business/contract/go/bank"
	"github.com/ardanlabs/liarsdice/foundation/smart/contract"
	"github.com/ardanlabs/liarsdice/foundation/smart/currency"
	"github.com/ethereum/go-ethereum/log"
)

// Harded this here for now just to make life easier.
const (
	primaryKeyPath    = "zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	primaryPassPhrase = "123"
	coinMarketCapKey  = "a8cd12fb-d056-423f-877b-659046af0aa5"
	network           = contract.NetworkLocalhost
)

func main() {
	if err := run(); err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}

func run() (dErr error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	converter, err := currency.NewConverter(coinMarketCapKey)
	if err != nil {
		fmt.Println("unable to create converter, using default:", err)
		converter = currency.NewDefaultConverter()
	}

	client, err := contract.NewClient(ctx, network, primaryKeyPath, primaryPassPhrase)
	if err != nil {
		return err
	}

	oneETHToUSD, oneUSDToETH := converter.Values()

	fmt.Println("\nSettings")
	fmt.Println("----------------------------------------------------")
	fmt.Println("network         :", network)
	fmt.Println("key address     :", client.Address())
	fmt.Println("oneETHToUSD     :", oneETHToUSD)
	fmt.Println("oneUSDToETH     :", oneUSDToETH)

	// =========================================================================

	startingBalance, err := client.CurrentBalance(ctx)
	if err != nil {
		return err
	}
	defer func() {
		endingBalance, err := client.CurrentBalance(ctx)
		if err != nil {
			dErr = err
			return
		}
		fmt.Print(converter.FmtBalanceSheet(startingBalance, endingBalance))
	}()

	// =========================================================================

	const gasLimit = 3000000
	const valueGwei = 0.0
	tranOpts, err := client.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		return err
	}

	// =========================================================================

	address, tx, _, err := bank.DeployBank(tranOpts, client.ContractBackend())
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransaction(tx))

	os.MkdirAll("zarf/contract/", 0755)
	if err := os.WriteFile("zarf/contract/id.env", []byte(address.Hex()), 0666); err != nil {
		return err
	}

	fmt.Println("\nContract Details")
	fmt.Println("----------------------------------------------------")
	fmt.Println("contract id     :", address.Hex())
	fmt.Printf("export GAME_CONTRACT_ID=%s\n", address.Hex())

	// =========================================================================

	clientWait, err := contract.NewClient(ctx, network, primaryKeyPath, primaryPassPhrase)
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
