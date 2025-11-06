package uniswap

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ykel/quote_engine/core/bindings"
	"github.com/ykel/quote_engine/core/config"
)

type Quoter struct {
	contract *bindings.IQuoterV2
}

func NewQuoter(client *ethclient.Client) (*Quoter, error) {

	contract, err := bindings.NewIQuoterV2(config.App.QuoterAddress, client)

	if err != nil {
		return nil, err
	}
	return &Quoter{contract}, nil
}

func (q *Quoter) QuoteExactInputSingle(amount *big.Int, tokenIn common.Address, tokenOut common.Address) (*QuoteResult, error) {

	callOps := &bind.CallOpts{Context: context.Background()}
	params := bindings.IQuoterV2QuoteExactInputSingleParams{
		TokenIn:           tokenIn,
		TokenOut:          tokenOut,
		AmountIn:          amount,
		Fee:               big.NewInt(config.App.DefaultFee),
		SqrtPriceLimitX96: big.NewInt(0),
	}

	result, err := q.contract.QuoteExactInputSingle(callOps, params)

	if err != nil {
		return nil, err
	}

	return &QuoteResult{
		Amount:                  result.AmountOut,
		SqrtPriceX96After:       result.SqrtPriceX96After,
		InitializedTicksCrossed: result.InitializedTicksCrossed,
		GasEstimate:             result.GasEstimate,
	}, nil

}

func (q *Quoter) QuoteExactOutputSingle(amount *big.Int, tokenIn common.Address, tokenOut common.Address) (*QuoteResult, error) {

	callOps := &bind.CallOpts{Context: context.Background()}

	params := bindings.IQuoterV2QuoteExactOutputSingleParams{
		TokenIn:           tokenIn,
		TokenOut:          tokenOut,
		Amount:            amount,
		Fee:               big.NewInt(config.App.DefaultFee),
		SqrtPriceLimitX96: big.NewInt(0),
	}

	result, err := q.contract.QuoteExactOutputSingle(callOps, params)

	if err != nil {
		return nil, err
	}

	return &QuoteResult{
		Amount:                  result.AmountIn,
		SqrtPriceX96After:       result.SqrtPriceX96After,
		InitializedTicksCrossed: result.InitializedTicksCrossed,
		GasEstimate:             result.GasEstimate,
	}, nil
}
