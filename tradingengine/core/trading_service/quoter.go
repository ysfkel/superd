package trading_service

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ykel/quote_engine/core/config"
	"github.com/ykel/quote_engine/core/uniswap"
)

func GetQuote(inputToken common.Address, amount *big.Int, side uniswap.Side, client *ethclient.Client) (*uniswap.QuoteResult, error) {

	quoter, err := uniswap.NewQuoter(client)

	if err != nil {
		return nil, err
	}

	var result *uniswap.QuoteResult

	if side == uniswap.Buy {
		// BUY TOKEN: Spend WETH, Get TOKEN
		result, err = quoter.QuoteExactOutputSingle(amount, config.App.WETHAddress, inputToken)

		if err != nil {
			return nil, err
		}
	} else {

		result, err = quoter.QuoteExactInputSingle(amount, inputToken, config.App.WETHAddress)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
