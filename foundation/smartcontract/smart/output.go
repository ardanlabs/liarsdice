package smart

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

// WriteBalanceSheet produces a easy to read format of the starting and ending
// balance for the connected account.
func (c *Client) WriteBalanceSheet(ctx context.Context, w io.Writer, startingBalance *big.Int) {
	endingBalance, err := c.CurrentBalance(ctx)
	if err != nil {
		return
	}

	diff, err := c.CalculateBalanceDiff(ctx, startingBalance, endingBalance)
	if err != nil {
		return
	}

	w.Write(formatBalanceDiff(diff))
}

// WriteTransaction produces a easy to read format of the specified transaction.
func (c *Client) WriteTransaction(w io.Writer, tx *types.Transaction) {
	tcd := c.CalculateTransactionDetails(tx)

	w.Write(formatTranCostDetails(tcd))
}

// WriteTransaction produces a easy to read format of the specified transaction.
func (c *Client) WriteTransactionReceipt(w io.Writer, receipt *types.Receipt, gasPrice *big.Int) {
	rcd := c.CalculateReceiptDetails(receipt, gasPrice)

	w.Write(formatReceiptCostDetails(rcd))
	w.Write(formatLogs(c.ExtractLogs(receipt)))
}

// =============================================================================

// formatTranCostDetails displays the transaction cost details.
func formatTranCostDetails(tcd TransactionDetails) []byte {
	var b bytes.Buffer

	fmt.Fprintf(&b, "\nTransaction Details\n")
	fmt.Fprintf(&b, "----------------------------------------------------\n")
	fmt.Fprintf(&b, "sent            : %v\n", tcd.Hash)
	fmt.Fprintf(&b, "gas limit       : %v\n", tcd.GasLimit)
	fmt.Fprintf(&b, "gas offer price : %v GWei\n", tcd.GasOfferPriceGWei)
	fmt.Fprintf(&b, "value           : %v GWei\n", tcd.Value)
	fmt.Fprintf(&b, "max gas price   : %v GWei\n", tcd.MaxGasPriceGWei)
	fmt.Fprintf(&b, "max gas price   : %v USD\n", tcd.MaxGasPriceUSD)

	return b.Bytes()
}

// formatReceiptCostDetails displays the receipt cost details.
func formatReceiptCostDetails(rcd ReceiptDetails) []byte {
	var b bytes.Buffer

	fmt.Fprintf(&b, "\nReceipt Details\n")
	fmt.Fprintf(&b, "----------------------------------------------------\n")
	fmt.Fprintf(&b, "status          : %v\n", rcd.Status)
	fmt.Fprintf(&b, "gas used        : %v\n", rcd.GasUsed)
	fmt.Fprintf(&b, "gas price       : %v GWei\n", rcd.GasPriceGWei)
	fmt.Fprintf(&b, "gas price       : %v USD\n", rcd.GasPriceUSD)
	fmt.Fprintf(&b, "final gas cost  : %v GWei\n", rcd.FinalCostGWei)
	fmt.Fprintf(&b, "final gas cost  : %v USD\n", rcd.FinalCostUSD)

	return b.Bytes()
}

// formatLogs takes the slice of log information and displays it.
func formatLogs(logs []string) []byte {
	var b bytes.Buffer

	fmt.Fprintf(&b, "\nLogs\n")
	fmt.Fprintf(&b, "----------------------------------------------------\n")
	for _, log := range logs {
		fmt.Fprintln(&b, log)
	}

	return b.Bytes()
}

// formatBalanceDiff outputs the start and ending balances with difference.
func formatBalanceDiff(bd BalanceDiff) []byte {
	var b bytes.Buffer

	fmt.Fprintf(&b, "\nBalance\n")
	fmt.Fprintf(&b, "----------------------------------------------------\n")
	fmt.Fprintf(&b, "balance before  : %v GWei\n", bd.BeforeGWei)
	fmt.Fprintf(&b, "balance after   : %v GWei\n", bd.AfterGWei)
	fmt.Fprintf(&b, "balance diff    : %v GWei\n", bd.DiffGWei)
	fmt.Fprintf(&b, "balance diff    : %v USD\n", bd.DiffUSD)

	return b.Bytes()
}
