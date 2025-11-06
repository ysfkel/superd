package uniswap

import (
	"context"
	"math/big"

	// "univ3-swapper/core/bindings"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ykel/quote_engine/core/bindings"
	"github.com/ykel/quote_engine/core/config"
)

type Factory struct {
	contract *bindings.IUniswapV3Factory
}

func NewFactory(client *ethclient.Client) *Factory {
	contract, _ := bindings.NewIUniswapV3Factory(config.App.FactoryAddress, client)
	return &Factory{contract}
}

func (c *Factory) GetPool(ctx context.Context, tokenIn common.Address, tokenOut common.Address, fee *big.Int) (*common.Address, error) {

	callOps := &bind.CallOpts{Context: ctx}

	result, err := c.contract.GetPool(callOps, tokenIn, tokenOut, fee)

	if err != nil {
		return nil, err
	}

	return &result, nil

}
