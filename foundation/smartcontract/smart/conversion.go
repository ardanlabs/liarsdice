package smart

import (
	"fmt"
	"math/big"

	ethUnit "github.com/DeOne4eg/eth-unit-converter"
)

// https://www.cryps.info/en/ETH_to_USD/
// https://www.cryps.info/en/USD_to_ETH/
var (
	oneEthToUSD  = big.NewFloat(1508.6386)
	oneUSDToEth  = big.NewFloat(0.00066285)
	oneGweiToUSD = big.NewFloat(0).Mul(oneEthToUSD, big.NewFloat(0.000000001))
	oneUSDToGWei = big.NewFloat(0).Mul(oneUSDToEth, big.NewFloat(1000000000))
)

// Wei2USD converts Wei to USD.
func Wei2USD(amountWei *big.Int) string {
	unit := ethUnit.NewWei(amountWei)
	gWeiAmount := unit.GWei()

	return GWei2USD(gWeiAmount)
}

// GWei2USD converts GWei to USD.
func GWei2USD(amountGWei *big.Float) string {
	cost := big.NewFloat(0).Mul(amountGWei, oneGweiToUSD)
	costFloat, _ := cost.Float64()
	return fmt.Sprintf("%.2f", costFloat)
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
func USD2Wei(amountUSD *big.Float) *big.Int {
	gwei := big.NewFloat(0).Mul(amountUSD, oneUSDToGWei)
	v, _ := big.NewFloat(0).Mul(gwei, big.NewFloat(1e9)).Int64()
	return big.NewInt(0).SetInt64(v)
}

// USD2GWei converts USD to GWei.
func USD2GWei(amountUSD *big.Float) *big.Float {
	return big.NewFloat(0).Mul(amountUSD, oneUSDToGWei)
}
