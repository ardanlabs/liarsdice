package bank

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ardanlabs/liarsdice/contract/sol/go/contract"
	"github.com/ardanlabs/liarsdice/foundation/smartcontract/smart"
	"github.com/ethereum/go-ethereum/common"
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

// Balance will return the specified address balance.
func (b *Bank) Balance(ctx context.Context, address string) (*big.Int, error) {
	tranOpts, err := b.client.NewCallOpts(ctx)
	if err != nil {
		return nil, err
	}

	player := common.HexToAddress(address)

	return b.contract.PlayerBalance(tranOpts, player)
}

// Reconcile will apply with ante to the winner and losers and provide the
// house the game fee.
func (b *Bank) Reconcile(ctx context.Context, winner string, losers []string, ante uint, gameFee uint) error {
	return nil
}
