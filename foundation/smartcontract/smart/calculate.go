package smart

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// TransactionDetails provides details about a transaction and it's cost.
type TransactionDetails struct {
	Hash              string
	GasLimit          uint64
	GasOfferPriceGWei string
	Value             string
	MaxGasPriceGWei   string
	MaxGasPriceUSD    string
}

// CalculateTranactionDetails performs calculations on the transaction.
func CalculateTransactionDetails(tx *types.Transaction) TransactionDetails {
	return TransactionDetails{
		Hash:              tx.Hash().Hex(),
		GasLimit:          tx.Gas(),
		GasOfferPriceGWei: Wei2GWei(tx.GasPrice()).String(),
		Value:             Wei2GWei(tx.Value()).String(),
		MaxGasPriceGWei:   Wei2GWei(tx.Cost()).String(),
		MaxGasPriceUSD:    Wei2USD(tx.Cost()),
	}
}

// =============================================================================

// ReceiptDetails provides details about a receipt and it's cost.
type ReceiptDetails struct {
	Status        uint64
	GasUsed       uint64
	GasPriceGWei  string
	GasPriceUSD   string
	FinalCostGWei string
	FinalCostUSD  string
}

// CalculateReceiptDetails performs calculations on the receipt.
func CalculateReceiptDetails(receipt *types.Receipt, gasPrice *big.Int) ReceiptDetails {
	cost := big.NewInt(0).Mul(big.NewInt(int64(receipt.GasUsed)), gasPrice)

	return ReceiptDetails{
		Status:        receipt.Status,
		GasUsed:       receipt.GasUsed,
		GasPriceGWei:  Wei2GWei(gasPrice).String(),
		GasPriceUSD:   Wei2USD(gasPrice),
		FinalCostGWei: Wei2GWei(cost).String(),
		FinalCostUSD:  Wei2USD(cost),
	}
}

// ExtractLogs extracts the logging events from the receipt.
func ExtractLogs(receipt *types.Receipt) []string {
	var logs []string

	if len(receipt.Logs) > 0 {

		// We have a particular log event that if we find, we can separate
		// from the rest of the events.
		topicLog := crypto.Keccak256Hash([]byte("EventLog(string)"))

		// Iterate over the logs and separate.
		for _, v := range receipt.Logs {
			if v.Topics[0] == topicLog {
				l := v.Data[63]
				logs = append(logs, string(v.Data[64:64+l]))
			}
		}
	}

	return logs
}

// =============================================================================

// BalanceDiff performs calculations on the starting and ending balance.
type BalanceDiff struct {
	BeforeGWei string
	AfterGWei  string
	DiffGWei   string
	DiffUSD    string
}

// CalculateBalanceDiff performs calculations on the starting and ending balance.
func CalculateBalanceDiff(ctx context.Context, startingBalance *big.Int, endingBalance *big.Int) (BalanceDiff, error) {
	cost := big.NewInt(0).Sub(startingBalance, endingBalance)

	return BalanceDiff{
		BeforeGWei: Wei2GWei(startingBalance).String(),
		AfterGWei:  Wei2GWei(endingBalance).String(),
		DiffGWei:   Wei2GWei(cost).String(),
		DiffUSD:    Wei2USD(cost),
	}, nil
}
