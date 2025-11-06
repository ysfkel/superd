package trading_service

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ykel/quote_engine/core/config"
	"github.com/ykel/quote_engine/core/tokens"
	"github.com/ykel/quote_engine/core/uniswap"
	"github.com/ykel/quote_engine/core/util"
	"github.com/ykel/quote_engine/core/wallet"
)

func ExecuteSwap(
	inputToken common.Address,
	quoteAmount *big.Int,
	baseAmount *big.Int,
	slippageAmount *big.Int,
	side uniswap.Side,
	client *ethclient.Client,
) (*types.Transaction, error) {
	swapper, err := uniswap.NewSwapper(client, config.App.RouterAddress)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\nSwap Details:\n")
	fmt.Printf("   Quote: %s WETH\n", util.FormatTokenAmount(quoteAmount, 18))
	fmt.Printf("   Amount: %s USDC\n", util.FormatTokenAmount(baseAmount, 6))
	fmt.Printf("   Slippage: %s WETH\n", util.FormatTokenAmount(slippageAmount, 18))

	if side == uniswap.Buy {
		return executeBuy(swapper, inputToken, baseAmount, slippageAmount, client)
	}

	return executeSell(swapper, inputToken, baseAmount, slippageAmount, client)
}

func executeBuy(
	swapper *uniswap.Swapper,
	inputToken common.Address,
	baseAmount *big.Int,
	slippageAmount *big.Int,
	client *ethclient.Client,
) (*types.Transaction, error) {

	opts, err := wallet.CreateSigner(client)
	if err != nil {
		return nil, err
	}

	opts.Value = slippageAmount
	tx, err := swapper.SwapExactOutputSingle(
		opts,
		uniswap.SwapOutRequest{
			TokenIn:         config.App.WETHAddress,
			TokenOut:        inputToken,
			AmountOut:       baseAmount,
			AmountInMaximum: slippageAmount,
			Side:            uniswap.Buy,
			Fee:             big.NewInt(config.App.DefaultFee),
			DeadlineMinutes: 20,
		},
	)

	if err != nil {
		return nil, err
	}

	return tx, nil
}

func executeSell(
	swapper *uniswap.Swapper,
	inputToken common.Address,
	baseAmount *big.Int,
	slippageAmount *big.Int,
	client *ethclient.Client,
) (*types.Transaction, error) {
	// Approve router first
	fmt.Println("\nApproving router...")
	if err := Approve(baseAmount, inputToken, client); err != nil {
		return nil, err
	}

	// Get fresh signer after approval
	opts, err := wallet.CreateSigner(client)
	if err != nil {
		return nil, err
	}

	tx, err := swapper.SwapExactInputSingle(
		opts,
		uniswap.SwapInRequest{
			TokenIn:          inputToken,
			TokenOut:         config.App.WETHAddress,
			AmountIn:         baseAmount,
			AmountOutMinimum: slippageAmount,
			Side:             uniswap.Sell,
			Fee:              big.NewInt(config.App.DefaultFee),
			DeadlineMinutes:  20,
		},
	)

	if err != nil {
		return nil, err
	}

	return tx, nil
}

func Approve(amount *big.Int, token common.Address, client *ethclient.Client) error {

	opts, err := wallet.CreateSigner(client)

	if err != nil {
		return err
	}

	erc20, err := tokens.NewErc20(client, token)

	if err != nil {
		return err
	}

	_, err = erc20.Approve(opts, config.App.RouterAddress, amount)

	if err != nil {
		return err
	}

	return nil
}
