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

type ERC20 struct {
	contract *bindings.IERC20
}

func NewErc20(client *ethclient.Client, erc20Address common.Address) (*ERC20, error) {
	contract, err := bindings.NewIERC20(erc20Address, client)

	if err != nil {
		return nil, err
	}
	return &ERC20{contract}, nil
}

func (c *ERC20) BalanceOf(ctx context.Context, account common.Address) (*big.Int, error) {

	callOps := &bind.CallOpts{Context: ctx}

	result, err := c.contract.BalanceOf(callOps, account)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (c *ERC20) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {

	tx, err := c.contract.Approve(opts, spender, amount)

	if err != nil {
		return nil, err
	}

	return tx, nil

}

func (c *ERC20) Decimals(ctx context.Context) (uint8, error) {

	callOps := &bind.CallOpts{Context: ctx}

	decimals, err := c.contract.Decimals(callOps)

	if err != nil {
		return decimals, err
	}

	return decimals, nil
}
