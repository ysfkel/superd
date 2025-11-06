package util

import (
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func ToBaseUnits(amount float64, decimals uint8) *big.Int {
	base := new(big.Float).SetFloat64(amount)
	multiplier := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil))
	base.Mul(base, multiplier)

	// Convert to big.Int
	result := new(big.Int)
	base.Int(result)
	return result
}

func ReadString(prompt string) string {
	fmt.Println(prompt)
	var input string
	fmt.Scanln(&input)
	return strings.TrimSpace(input)
}

func ReadFloat(prompt string) (float64, error) {
	fmt.Println(prompt)
	var input float64
	_, err := fmt.Scanln(&input)
	return input, err
}

func ReadAddress(prompt string) (common.Address, error) {
	fmt.Println(prompt)
	var input string
	fmt.Scanln(&input)
	input = strings.TrimSpace(input)

	if !common.IsHexAddress(input) {
		return common.Address{}, fmt.Errorf("invalid address format")
	}

	return common.HexToAddress(input), nil
}

func CalculateMinimumAmountInt(amount *big.Int, slippagePercent float64) *big.Int {
	bps := int64(slippagePercent * 100)
	multiplier := big.NewInt(10000 - bps)
	result := new(big.Int).Mul(amount, multiplier)
	return new(big.Int).Div(result, big.NewInt(10000))
}

func CalculateMaximumAmountInt(amount *big.Int, slippagePercent float64) *big.Int {
	bps := int64(slippagePercent * 100)
	multiplier := big.NewInt(10000 + bps)
	result := new(big.Int).Mul(amount, multiplier)
	return new(big.Int).Div(result, big.NewInt(10000))
}

func FormatTokenAmount(amount *big.Int, decimals uint8) string {
	if amount == nil {
		return "0"
	}

	fAmount := new(big.Float).SetInt(amount)
	divisor := new(big.Float).SetFloat64(math.Pow10(int(decimals)))
	human := new(big.Float).Quo(fAmount, divisor)

	return human.Text('f', 6) // show up to 6 decimal places
}
