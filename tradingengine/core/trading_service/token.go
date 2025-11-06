package trading_service

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ykel/quote_engine/core/config"
	"github.com/ykel/quote_engine/core/tokens"
)

func GetDecimals(client *ethclient.Client, token common.Address) (uint8, error) {

	erc20, err := tokens.NewErc20(client, token)

	if err != nil {
		return 0, err
	}

	decimals, err := erc20.Decimals(context.Background())

	if err != nil {
		return 0, err
	}

	return decimals, nil

}

func GetBalanceOf(token common.Address, client *ethclient.Client) (*big.Int, error) {

	erc20, err := tokens.NewErc20(client, token)

	if err != nil {
		return nil, err
	}

	balance, err := erc20.BalanceOf(context.Background(), config.App.WalletAddress)

	if err != nil {
		return nil, err
	}

	return balance, nil
}
