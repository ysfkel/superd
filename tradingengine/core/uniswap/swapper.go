package uniswap

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ykel/quote_engine/core/bindings"
)

type ExactInputSingleParams = bindings.ISwapRouterExactInputSingleParams
type ExactOutputSingleParams = bindings.ISwapRouterExactOutputSingleParams
type Swapper struct {
	client  *ethclient.Client
	router  *bindings.ISwapRouter
	address common.Address
}

func NewSwapper(client *ethclient.Client, routerAddress common.Address) (*Swapper, error) {
	router, err := bindings.NewISwapRouter(routerAddress, client)
	if err != nil {
		return nil, err
	}

	return &Swapper{
		client:  client,
		router:  router,
		address: routerAddress,
	}, nil
}

func (s *Swapper) SwapExactInputSingle(
	opts *bind.TransactOpts,
	req SwapInRequest,
) (*types.Transaction, error) {

	deadline := big.NewInt(time.Now().Add(time.Duration(req.DeadlineMinutes) * time.Minute).Unix())

	params := ExactInputSingleParams{
		TokenIn:           req.TokenIn,
		TokenOut:          req.TokenOut,
		Fee:               req.Fee,
		Recipient:         opts.From,
		Deadline:          deadline,
		AmountIn:          req.AmountIn,
		AmountOutMinimum:  req.AmountOutMinimum,
		SqrtPriceLimitX96: big.NewInt(0),
	}

	tx, err := s.router.ExactInputSingle(opts, params)
	if err != nil {
		return nil, fmt.Errorf("swap execution failed: %w", err)
	}

	return tx, nil
}

func (s *Swapper) SwapExactOutputSingle(
	opts *bind.TransactOpts,
	req SwapOutRequest,
) (*types.Transaction, error) {

	deadline := big.NewInt(time.Now().Add(time.Duration(req.DeadlineMinutes) * time.Minute).Unix())

	params := ExactOutputSingleParams{
		TokenIn:           req.TokenIn,
		TokenOut:          req.TokenOut,
		Fee:               req.Fee,
		Recipient:         opts.From,
		Deadline:          deadline,
		AmountOut:         req.AmountOut,
		AmountInMaximum:   req.AmountInMaximum,
		SqrtPriceLimitX96: big.NewInt(0),
	}

	tx, err := s.router.ExactOutputSingle(opts, params)
	if err != nil {
		return nil, fmt.Errorf("swap execution failed: %w", err)
	}

	return tx, nil
}
