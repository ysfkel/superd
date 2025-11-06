package uniswap

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Side string

const (
	Buy  Side = "BUY"
	Sell Side = "SELL"
)

// type QuoteRequest struct {
// 	TokenAddress common.Address
// 	Quantity     *big.Int
// 	Side         Side
// }

type QuoteResult struct {
	Amount                  *big.Int
	SqrtPriceX96After       *big.Int
	InitializedTicksCrossed uint32
	GasEstimate             *big.Int
}

type QuoteResponse struct {
	Side         Side
	TokenAddress common.Address
	InputAmount  *big.Int
	OutputAmount *big.Int
}

type SwapInRequest struct {
	TokenIn          common.Address
	TokenOut         common.Address
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
	Side             Side
	Fee              *big.Int // Optional, defaults to 3000 (0.3%)
	DeadlineMinutes  int64    // Optional, defaults to 20 minutes
}

type SwapOutRequest struct {
	TokenIn         common.Address
	TokenOut        common.Address
	AmountOut       *big.Int
	AmountInMaximum *big.Int
	Side            Side
	Fee             *big.Int // Optional, defaults to 3000 (0.3%)
	DeadlineMinutes int64    // Optional, defaults to 20 minutes
}
