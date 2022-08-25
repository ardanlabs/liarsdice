package smart

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

// FmtBalanceSheet produces a easy to read format of the starting and ending
// balance for the connected account.
func (c *Client) FmtBalanceSheet(ctx context.Context, startingBalance *big.Int) string {
	endingBalance, err := c.CurrentBalance(ctx)
	if err != nil {
		return ""
	}

	diff, err := CalculateBalanceDiff(ctx, startingBalance, endingBalance)
	if err != nil {
		return ""
	}

	return formatBalanceDiff(diff)
}

// FmtTransaction produces a easy to read format of the specified transaction.
func FmtTransaction(tx *types.Transaction) string {
	tcd := CalculateTransactionDetails(tx)

	return formatTranCostDetails(tcd)
}

// FmtTransaction produces a easy to read format of the specified transaction.
func FmtTransactionReceipt(receipt *types.Receipt, gasPrice *big.Int) string {
	rcd := CalculateReceiptDetails(receipt, gasPrice)

	var b bytes.Buffer

	b.WriteString(formatReceiptCostDetails(rcd))
	b.WriteString(formatLogs(ExtractLogs(receipt)))

	return b.String()
}

// =============================================================================

// formatTranCostDetails displays the transaction cost details.
func formatTranCostDetails(tcd TransactionDetails) string {
	var b bytes.Buffer

	fmt.Fprintf(&b, "\nTransaction Details\n")
	fmt.Fprintf(&b, "----------------------------------------------------\n")
	fmt.Fprintf(&b, "sent            : %v\n", tcd.Hash)
	fmt.Fprintf(&b, "gas limit       : %v\n", tcd.GasLimit)
	fmt.Fprintf(&b, "gas offer price : %v GWei\n", tcd.GasOfferPriceGWei)
	fmt.Fprintf(&b, "value           : %v GWei\n", tcd.Value)
	fmt.Fprintf(&b, "max gas price   : %v GWei\n", tcd.MaxGasPriceGWei)
	fmt.Fprintf(&b, "max gas price   : %v USD\n", tcd.MaxGasPriceUSD)

	return b.String()
}

// formatReceiptCostDetails displays the receipt cost details.
func formatReceiptCostDetails(rcd ReceiptDetails) string {
	var b bytes.Buffer

	fmt.Fprintf(&b, "\nReceipt Details\n")
	fmt.Fprintf(&b, "----------------------------------------------------\n")
	fmt.Fprintf(&b, "status          : %v\n", rcd.Status)
	fmt.Fprintf(&b, "gas used        : %v\n", rcd.GasUsed)
	fmt.Fprintf(&b, "gas price       : %v GWei\n", rcd.GasPriceGWei)
	fmt.Fprintf(&b, "gas price       : %v USD\n", rcd.GasPriceUSD)
	fmt.Fprintf(&b, "final gas cost  : %v GWei\n", rcd.FinalCostGWei)
	fmt.Fprintf(&b, "final gas cost  : %v USD\n", rcd.FinalCostUSD)

	return b.String()
}

// formatLogs takes the slice of log information and displays it.
func formatLogs(logs []string) string {
	var b bytes.Buffer

	fmt.Fprintf(&b, "\nLogs\n")
	fmt.Fprintf(&b, "----------------------------------------------------\n")
	for _, log := range logs {
		fmt.Fprintln(&b, log)
	}

	return b.String()
}

// formatBalanceDiff outputs the start and ending balances with difference.
func formatBalanceDiff(bd BalanceDiff) string {
	var b bytes.Buffer

	fmt.Fprintf(&b, "\nBalance\n")
	fmt.Fprintf(&b, "----------------------------------------------------\n")
	fmt.Fprintf(&b, "balance before  : %v GWei\n", bd.BeforeGWei)
	fmt.Fprintf(&b, "balance after   : %v GWei\n", bd.AfterGWei)
	fmt.Fprintf(&b, "balance diff    : %v GWei\n", bd.DiffGWei)
	fmt.Fprintf(&b, "balance diff    : %v USD\n", bd.DiffUSD)

	return b.String()
}
