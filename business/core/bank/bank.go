// Package bank represents all the transactions necessary for the game.
package bank

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	"github.com/ardanlabs/liarsdice/business/contract/go/bank"
	"github.com/ardanlabs/liarsdice/foundation/logger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Bank represents a bank that allows for the reconciling of a game and
// information about account balances.
type Bank struct {
	log        *logger.Logger
	contractID common.Address
	ethereum   *ethereum.Client
	contract   *bank.Bank
}

// New returns a new bank with the ability to manage the game money.
func New(ctx context.Context, log *logger.Logger, backend ethereum.Backend, privateKey *ecdsa.PrivateKey, contractID common.Address) (*Bank, error) {
	clt, err := ethereum.NewClient(backend, privateKey)
	if err != nil {
		return nil, fmt.Errorf("client: %w", err)
	}

	contract, err := bank.NewBank(contractID, clt.Backend)
	if err != nil {
		return nil, fmt.Errorf("new contract: %w", err)
	}

	b := Bank{
		log:        log,
		contractID: contractID,
		ethereum:   clt,
		contract:   contract,
	}

	return &b, nil
}

// ContractID returns contract id in use.
func (b *Bank) ContractID() common.Address {
	return b.contractID
}

// Client returns the underlying contract client.
func (b *Bank) Client() *ethereum.Client {
	return b.ethereum
}

// EthereumBalance returns the ethereum balance for the connected account.
func (b *Bank) EthereumBalance(ctx context.Context) (wei *big.Int, err error) {
	balance, err := b.ethereum.Balance(ctx)
	if err != nil {
		return nil, fmt.Errorf("current balance: %w", err)
	}

	return balance, nil
}

// Balance will return the bank balance for the connected account.
func (b *Bank) Balance(ctx context.Context) (GWei *big.Float, err error) {
	tranOpts, err := b.ethereum.NewCallOpts(ctx)
	if err != nil {
		return nil, fmt.Errorf("new call opts: %w", err)
	}

	wei, err := b.contract.Balance(tranOpts)
	if err != nil {
		return nil, fmt.Errorf("account balance: %w", err)
	}

	b.log.Info(ctx, "balance", "accountid", b.ethereum.Address().String(), "wei", wei)

	return currency.Wei2GWei(wei), nil
}

// AccountBalance will return the bank balance for the specified account. Only
// the owner of the smart contract can make this call.
func (b *Bank) AccountBalance(ctx context.Context, accountID common.Address) (GWei *big.Float, err error) {
	tranOpts, err := b.ethereum.NewCallOpts(ctx)
	if err != nil {
		return nil, fmt.Errorf("new call opts: %w", err)
	}

	wei, err := b.contract.AccountBalance(tranOpts, accountID)
	if err != nil {
		return nil, fmt.Errorf("account balance: %w", err)
	}

	b.log.Info(ctx, "account balance", "accountid", accountID, "wei", wei)

	return currency.Wei2GWei(wei), nil
}

// Reconcile will apply with ante to the winner and loser accounts, plus provide
// the house the game fee.
func (b *Bank) Reconcile(ctx context.Context, winningAccountID common.Address, losingAccountIDs []common.Address, anteGWei *big.Float, gameFeeGWei *big.Float) (*types.Transaction, *types.Receipt, error) {
	tranOpts, err := b.ethereum.NewTransactOpts(ctx, 0, big.NewInt(0), big.NewFloat(0))
	if err != nil {
		return nil, nil, fmt.Errorf("new trans opts: %w", err)
	}

	var loserIDs []common.Address
	loserIDs = append(loserIDs, losingAccountIDs...)

	anteWei := currency.GWei2Wei(anteGWei)
	gameFeeWei := currency.GWei2Wei(gameFeeGWei)

	tx, err := b.contract.Reconcile(tranOpts, winningAccountID, loserIDs, anteWei, gameFeeWei)
	if err != nil {
		return nil, nil, fmt.Errorf("reconcile: %w", err)
	}

	b.log.Info(ctx, "reconcile started", "anteWei", anteWei, "gameFeeWei", gameFeeWei)

	receipt, err := b.ethereum.WaitMined(ctx, tx)
	if err != nil {
		return nil, nil, fmt.Errorf("wait mined: %w", err)
	}

	b.log.Info(ctx, "reconcile completed")

	return tx, receipt, nil
}

// Deposit will add the given amount to the account's contract balance.
func (b *Bank) Deposit(ctx context.Context, amountGWei *big.Float) (*types.Transaction, *types.Receipt, error) {
	tranOpts, err := b.ethereum.NewTransactOpts(ctx, 0, big.NewInt(0), amountGWei)
	if err != nil {
		return nil, nil, fmt.Errorf("new trans opts: %w", err)
	}

	tx, err := b.contract.Deposit(tranOpts)
	if err != nil {
		return nil, nil, fmt.Errorf("deposit: %w", err)
	}

	b.log.Info(ctx, "deposit started", "accountid", b.ethereum.Address().String(), "amountGWei", amountGWei)

	receipt, err := b.ethereum.WaitMined(ctx, tx)
	if err != nil {
		return nil, nil, fmt.Errorf("wait mined: %w", err)
	}

	b.log.Info(ctx, "deposit completed")

	return tx, receipt, nil
}

// Withdraw will move all the account's balance in the contract, to the account's wallet.
func (b *Bank) Withdraw(ctx context.Context) (*types.Transaction, *types.Receipt, error) {
	tranOpts, err := b.ethereum.NewTransactOpts(ctx, 0, big.NewInt(0), big.NewFloat(0))
	if err != nil {
		return nil, nil, fmt.Errorf("new trans opts: %w", err)
	}

	tx, err := b.contract.Withdraw(tranOpts)
	if err != nil {
		return nil, nil, fmt.Errorf("withdraw: %w", err)
	}

	b.log.Info(ctx, "withdraw started", "accountid", b.ethereum.Address().String())

	receipt, err := b.ethereum.WaitMined(ctx, tx)
	if err != nil {
		return nil, nil, fmt.Errorf("wait mined: %w", err)
	}

	b.log.Info(ctx, "withdraw completed")

	return tx, receipt, nil
}
