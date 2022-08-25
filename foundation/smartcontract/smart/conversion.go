package smart

import (
	"fmt"
	"math/big"

	ethUnit "github.com/DeOne4eg/eth-unit-converter"
)

// These values are used to calculate Wei values in both GWei and USD.
// https://nomics.com/markets/gwei-gwei/usd-united-states-dollar
var (
	GWeiConv  = big.NewInt(1_000_000_000)
	GWeiPrice = big.NewFloat(0.00000255) // 8/23/22
	USDPrice  = big.NewFloat(391424.02)  // $1 USD to GWei
)

// Wei2USD converts Wei to USD.
func Wei2USD(amount *big.Int) string {
	unit := ethUnit.NewWei(amount)
	gWeiAmount := unit.GWei()

	// Multiple the current price of GWei to the USD.
	cost := big.NewFloat(0).Mul(gWeiAmount, GWeiPrice)
	costFloat, _ := cost.Float64()
	return fmt.Sprintf("%.8f", costFloat)
}

// Wei2GWei converts the wei unit into a GWei for display.
func Wei2GWei(amount *big.Int) string {
	unit := ethUnit.NewWei(amount)
	return unit.GWei().String()
}

// USD2Wei converts USD to Wei.
func USD2Wei(amount *big.Float) *big.Int {
	gwei := big.NewFloat(0).Mul(amount, USDPrice)
	multiplier := big.NewFloat(1e9)
	v, _ := big.NewFloat(0).Mul(gwei, multiplier).Int64()
	return big.NewInt(0).SetInt64(v)
}
