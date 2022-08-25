// Package bank represents all the transactions necessary for the game.
package bank

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ardanlabs/liarsdice/contract/sol/go/contract"
	"github.com/ardanlabs/liarsdice/foundation/smartcontract/smart"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Bank represents a bank that allows for the reconciling of a game and
// information about player balances.
type Bank struct {
	client   *smart.Client
	contract *contract.Contract
}

// New returns a new bank with the ability to manage the game money.
func New(ctx context.Context, network string, keyPath string, passPhrase string, contractID string) (*Bank, error) {
	client, err := smart.Connect(ctx, network, keyPath, passPhrase)
	if err != nil {
		return nil, err
	}

	contract, err := contract.NewContract(common.HexToAddress(contractID), client.ContractBackend())
	if err != nil {
		return nil, fmt.Errorf("NewContract: %w", err)
	}

	bank := Bank{
		client:   client,
		contract: contract,
	}

	return &bank, nil
}

// Deposit will add the given amount to the player's contract balance.
func (b *Bank) Deposit(ctx context.Context, account string, amount int64) error {
	tranOpts, err := b.client.NewTransactOpts(ctx, 0, uint64(amount))
	if err != nil {
		return err
	}

	tx, err := b.contract.Deposit(tranOpts)
	if err != nil {
		return err
	}

	_, err = b.client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	return nil
}

// Balance will return the balance for the specified account.
func (b *Bank) Balance(ctx context.Context, account string) (*big.Int, error) {
	tranOpts, err := b.client.NewCallOpts(ctx)
	if err != nil {
		return nil, err
	}

	return b.contract.PlayerBalance(tranOpts, common.HexToAddress(account))
}

// Reconcile will apply with ante to the winner and loser accounts, plus provide
// the house the game fee.
func (b *Bank) Reconcile(ctx context.Context, winningAccount string, losingAccounts []string, anteWei *big.Int, gameFeeWei *big.Int) (*types.Transaction, *types.Receipt, error) {
	tranOpts, err := b.client.NewTransactOpts(ctx, 0, 0)
	if err != nil {
		return nil, nil, err
	}

	winner := common.HexToAddress(winningAccount)

	var losers []common.Address

	for _, loser := range losingAccounts {
		losers = append(losers, common.HexToAddress(loser))
	}

	tx, err := b.contract.Reconcile(tranOpts, winner, losers, anteWei, gameFeeWei)
	if err != nil {
		return nil, nil, err
	}

	receipt, err := b.client.WaitMined(ctx, tx)
	if err != nil {
		return nil, nil, err
	}

	return tx, receipt, nil
}

// Withdraw will move all the player's balance in the contract, to the player's wallet.
func (b *Bank) Withdraw(ctx context.Context, account string) error {
	tranOpts, err := b.client.NewTransactOpts(ctx, 0, 0)
	if err != nil {
		return err
	}

	tx, err := b.contract.Withdraw(tranOpts)
	if err != nil {
		return err
	}

	if _, err = b.client.WaitMined(ctx, tx); err != nil {
		return err
	}

	return nil
}

// WalletBalance returns the current balance for the account used to
// create this bank.
func (b *Bank) WalletBalance(ctx context.Context) (*big.Int, error) {
	balance, err := b.client.CurrentBalance(ctx)
	if err != nil {
		return nil, err
	}

	return balance, nil
}
