package tokens

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ykel/quote_engine/core/bindings"
)

type Weth struct {
	contract *bindings.IWETH
}

func NewWeth(client *ethclient.Client, address common.Address) (*Weth, error) {
	contract, err := bindings.NewIWETH(address, client)

	if err != nil {
		return nil, err
	}
	return &Weth{contract}, nil
}

func (c *Weth) BalanceOf(ctx context.Context, account common.Address) (*big.Int, error) {

	callOps := &bind.CallOpts{Context: ctx}

	result, err := c.contract.BalanceOf(callOps, account)

	if err != nil {
		return nil, err
	}

	return result, nil

}
func (c *Weth) Deposit(amount *big.Int, opts *bind.TransactOpts) (*types.Transaction, error) {

	opts.Value = amount

	result, err := c.contract.Deposit(opts)

	if err != nil {
		return nil, err
	}

	return result, nil

}
