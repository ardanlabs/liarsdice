package smart

import (
	"fmt"
	"math/big"

	ethUnit "github.com/DeOne4eg/eth-unit-converter"
)

// These values are used to calculate Wei values in both GWei and USD.
// https://nomics.com/markets/gwei-gwei/usd-united-states-dollar
var (
	OneGWeiInUSD = big.NewFloat(0.00000255) // 8/23/22
	OneUSDInGwei = big.NewFloat(391424.02)  // $1 USD to GWei
)

// Wei2USD converts Wei to USD.
func Wei2USD(amountWei *big.Int) string {
	unit := ethUnit.NewWei(amountWei)
	gWeiAmount := unit.GWei()

	return GWei2USD(gWeiAmount)
}

// GWei2USD converts GWei to USD.
func GWei2USD(amountGWei *big.Float) string {
	cost := big.NewFloat(0).Mul(amountGWei, OneGWeiInUSD)
	costFloat, _ := cost.Float64()
	return fmt.Sprintf("%.8f", costFloat)
}

// Wei2GWei converts the wei unit into a GWei for display.
func Wei2GWei(amountWei *big.Int) *big.Float {
	unit := ethUnit.NewWei(amountWei)
	return unit.GWei()
}

// GWei2Wei converts the wei unit into a GWei for display.
func GWei2Wei(amountGWei *big.Float) *big.Int {
	unit := ethUnit.NewGWei(amountGWei)
	return unit.Wei()
}

// USD2Wei converts USD to Wei.
func USD2Wei(amountGWei *big.Float) *big.Int {
	gwei := big.NewFloat(0).Mul(amountGWei, OneUSDInGwei)
	multiplier := big.NewFloat(1e9)
	v, _ := big.NewFloat(0).Mul(gwei, multiplier).Int64()
	return big.NewInt(0).SetInt64(v)
}

// USD2GWei converts USD to GWei.
func USD2GWei(amount *big.Float) *big.Float {
	return big.NewFloat(0).Mul(amount, OneUSDInGwei)
}
