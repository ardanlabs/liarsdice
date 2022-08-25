// Package smart provides smart contract support.
package smart

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"

	"github.com/ardanlabs/liarsdice/foundation/smartcontract/etherscan"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Set of networks supported by the smart package.
const (
	NetworkHTTPLocalhost = "http://localhost:8545"
	NetworkLocalhost     = "zarf/ethereum/geth.ipc"
	NetworkGoerli        = "https://rpc.ankr.com/eth_goerli"
)

// Harded this here for now just to make life easier.
const (
	PrimaryKeyPath    = "zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	PrimaryPassPhrase = "123"

	Player1Address    = "0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7"
	Player1KeyPath    = "zarf/ethereum/keystore/UTC--2022-05-13T16-59-42.277071000Z--0070742ff6003c3e809e78d524f0fe5dcc5ba7f7"
	Player1PassPhrase = "123"

	Player2Address    = "0x8e113078adf6888b7ba84967f299f29aece24c55"
	Player2KeyPath    = "zarf/ethereum/keystore/UTC--2022-05-13T16-57-20.203544000Z--8e113078adf6888b7ba84967f299f29aece24c55"
	Player2PassPhrase = "123"
)

// =============================================================================

// Client provides an API for working with smart contracts.
type Client struct {
	Network string
	Account common.Address

	privateKey *ecdsa.PrivateKey
	chainID    *big.Int
	ethClient  *ethclient.Client
	etherscan  *etherscan.Etherscan
}

// Connect provides boilerplate for connecting to the geth service using
// an IPC socket created by the geth service on startup.
func Connect(ctx context.Context, network string, keyPath string, passPhrase string) (*Client, error) {
	ethClient, err := ethclient.Dial(network)
	if err != nil {
		return nil, fmt.Errorf("dial network: %w", err)
	}

	privateKey, err := privateKeyByKeyFile(keyPath, passPhrase)
	if err != nil {
		return nil, fmt.Errorf("extract private key: %w", err)
	}

	chainID, err := ethClient.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("capture chain id: %w", err)
	}

	var eth *etherscan.Etherscan
	etherscanApiKey := os.Getenv("ETHERSCAN")
	if etherscanApiKey != "" {
		eth = etherscan.New(etherscanApiKey)
	}

	c := Client{
		Network: network,
		Account: crypto.PubkeyToAddress(privateKey.PublicKey),

		privateKey: privateKey,
		chainID:    chainID,
		ethClient:  ethClient,
		etherscan:  eth,
	}

	return &c, nil
}

// NewCallOpts constructs a new CallOpts which is used to call contract methods
// that does not require a transaction.
func (c *Client) NewCallOpts(ctx context.Context) (*bind.CallOpts, error) {
	call := bind.CallOpts{
		Pending: true,
		From:    c.Account,
		Context: ctx,
	}

	return &call, nil
}

// NewTransaction constructs a new TransactOpts which is the collection of
// authorization data required to create a valid Ethereum transaction.
func (c *Client) NewTransactOpts(ctx context.Context, gasLimit uint64, valueGwei uint64) (*bind.TransactOpts, error) {
	nonce, err := c.ethClient.PendingNonceAt(ctx, c.Account)
	if err != nil {
		return nil, fmt.Errorf("retrieving next nonce: %w", err)
	}

	gasPrice, err := c.ethClient.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("retrieving suggested gas price: %w", err)
	}

	tranOpts, err := bind.NewKeyedTransactorWithChainID(c.privateKey, c.chainID)
	if err != nil {
		return nil, fmt.Errorf("keying transaction: %w", err)
	}

	// The value must be converted from Gwei to Wei.
	valueWei := big.NewInt(0).Mul(big.NewInt(int64(valueGwei)), GWeiConv)

	tranOpts.Nonce = big.NewInt(int64(nonce))
	tranOpts.Value = valueWei
	tranOpts.GasLimit = gasLimit // The maximum amount of Gas you are willing to pay for.
	tranOpts.GasPrice = gasPrice // What you will agree to pay per unit of gas.

	return tranOpts, nil
}

// WaitMined will wait for the transaction to be minded and return a receipt.
func (c *Client) WaitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	receipt, err := bind.WaitMined(ctx, c.ethClient, tx)
	if err != nil {
		return nil, fmt.Errorf("waiting for tx to be mined: %w", err)
	}

	if receipt.Status == 0 {
		err := c.extractError(ctx, tx)
		return nil, fmt.Errorf("extracting tx error: %w", err)
	}

	return receipt, nil
}

// Transaction returns a transaction value for the specified transaction hash.
func (c *Client) Transaction(ctx context.Context, txHash common.Hash) (*types.Transaction, bool, error) {
	return c.ethClient.TransactionByHash(ctx, txHash)
}

// TransactionReceipt returns a receipt value for the specified transaction hash.
func (c *Client) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return c.ethClient.TransactionReceipt(ctx, txHash)
}

// BaseFee calculates the base fee from the block for this receipt.
func (c *Client) BaseFee(receipt *types.Receipt) *big.Int {
	block, err := c.ethClient.BlockByNumber(context.Background(), receipt.BlockNumber)
	if err != nil {
		return big.NewInt(0)
	}
	return block.BaseFee()
}

// CurrentBalance retrieves the current balance for the account.
func (c *Client) CurrentBalance(ctx context.Context) (*big.Int, error) {
	balance, err := c.ethClient.BalanceAt(ctx, c.Account, nil)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

// ContractBackend returns the ethereum client. This is needed for smart
// contract creation and other calls.
func (c *Client) ContractBackend() *ethclient.Client {
	return c.ethClient
}

// =============================================================================

// extractError checks the failed transaction for the error message.
func (c *Client) extractError(ctx context.Context, tx *types.Transaction) error {
	msg := ethereum.CallMsg{
		From:     c.Account,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}

	_, err := c.ethClient.CallContract(ctx, msg, nil)
	return err
}

// privateKeyByKeyFile opens a key file for the private key.
func privateKeyByKeyFile(keyPath string, passPhrase string) (*ecdsa.PrivateKey, error) {
	data, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}

	key, err := keystore.DecryptKey(data, passPhrase)
	if err != nil {
		return nil, err
	}

	return key.PrivateKey, nil
}
