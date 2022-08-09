package bank

import (
	"context"
	"math/big"

	"github.com/ardanlabs/liarsdice/foundation/smartcontract/smart"
)

// Bank represents a bank that allows for the reconciling of a game and
// information about player balances.
type Bank struct {
	client *smart.Client
}

// NewBank returns a new bank with the ability to manage the game money.
func NewBank(ctx context.Context, network string, keyPath string, passPhrase string) (*Bank, error) {
	client, err := smart.Connect(ctx, network, keyPath, passPhrase)
	if err != nil {
		return nil, err
	}

	bank := Bank{
		client: client,
	}

	return &bank, nil
}

// PlayerBalance will return the specified player's balance.
func (b *Bank) PlayerBalance(address string) (*big.Int, error) {
	return nil, nil
}

// Reconcile will apply with ante to the winner and losers and provide the
// house the game fee.
func (b *Bank) Reconcile(winner string, losers []string, ante uint, gameFee uint) error {
	return nil
}
