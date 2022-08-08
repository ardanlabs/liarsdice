package bank

import (
	"context"
	"math/big"

	"github.com/ardanlabs/liarsdice/foundation/smartcontract/smart"
)

// Banker interface declares the bank behaviour.
type Banker interface {
	PlayerBalance(address string) (*big.Int, error)
	Reconcile(winner string, losers []string, ante uint, gameFee uint)
}

// BankConfig contains all the information required by the bank client.
type BankConfig struct {
	Network    string
	KeyPath    string
	PassPhrase string
}

type bank struct {
	client *smart.Client
}

// NewBank returns a new Banker interface type with a smart client.
func NewBank(cfg BankConfig) (Banker, error) {
	var ctx context.Context

	client, err := smart.Connect(ctx, cfg.Network, cfg.KeyPath, cfg.PassPhrase)
	if err != nil {
		return nil, err
	}

	bank := bank{
		client: client,
	}

	return &bank, nil
}

// PlayerBalance will call the contract PlayerBalance method.
func (b *bank) PlayerBalance(address string) (*big.Int, error) {
	return nil, nil
}

// Reconcilse will call the contract Reconcile method.
func (b *bank) Reconcile(winner string, losers []string, ante uint, gameFee uint) {
}
