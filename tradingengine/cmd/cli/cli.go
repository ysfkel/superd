package cli

import (
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ykel/quote_engine/core/client"
	"github.com/ykel/quote_engine/core/config"
	"github.com/ykel/quote_engine/core/trading_service"
	"github.com/ykel/quote_engine/core/uniswap"
	"github.com/ykel/quote_engine/core/util"
)

func RunCli() {

	amount, tokenAddress, side := readQuoteInputs()
	client := client.MustNewRPCClient(config.App.RPCURL)

	decimals, err := trading_service.GetDecimals(client, tokenAddress)

	if err != nil {
		log.Fatal(err)
	}

	baseAmount := util.ToBaseUnits(amount, decimals)
	fmt.Println("Finding quote, pleae wait..")
	result, err := trading_service.GetQuote(tokenAddress, baseAmount, side, client)

	if err != nil {
		log.Fatalln(err)
	}

	quoteAmount := result.Amount // WETH amount from quoter

	slippageAmount := readSwapInputs(quoteAmount, side)

	if slippageAmount == nil {
		return
	}

	_, err = trading_service.ExecuteSwap(tokenAddress, quoteAmount, baseAmount, slippageAmount, side, client)

	if err != nil {
		log.Fatalln(err)
	}

	balance, err := trading_service.GetBalanceOf(tokenAddress, client)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Your Balance: ", util.FormatTokenAmount(balance, 6))

}

func readSwapInputs(amount *big.Int, side uniswap.Side) (slippageAmount *big.Int) {
	accept := util.ReadString(fmt.Sprintf("Quote: %s Accept? (y/n)", util.FormatTokenAmount(amount, 18)))

	if strings.ToLower(accept) != "yes" && strings.ToLower(accept) != "y" {
		return nil
	}

	// Show slippage options
	fmt.Println("\nSlippage Tolerance Options:")
	fmt.Println("   1. 0.1%")
	fmt.Println("   2. 0.5% (default)")
	fmt.Println("   3. 1.0%")
	fmt.Println("   4. 5.0%")

	choice := util.ReadString("\nSelect option (1-5): ")

	var slippage float64
	switch strings.TrimSpace(choice) {
	case "1":
		slippage = 0.1
	case "2":
		slippage = 0.5
	case "3":
		slippage = 1.0
	case "4":
		slippage = 5.0
	default:
		slippage = 0.5 // Default to 0.5%
		fmt.Println("Invalid choice, using default 0.5%")
	}

	if side == uniswap.Buy {
		slippageAmount = util.CalculateMaximumAmountInt(amount, slippage)
	} else {
		slippageAmount = util.CalculateMinimumAmountInt(amount, slippage)
	}

	return slippageAmount
}

func readQuoteInputs() (amount float64, tokenInput common.Address, side uniswap.Side) {

	tokenInput, err := util.ReadAddress("Enter token symbol (e.g. USDC) or Provide valid token address:")

	if err != nil {
		log.Fatal("Invalid Ethereum address format")
	}

	amount, err = util.ReadFloat("Enter trade Amount (e.g 100)")

	if err != nil {
		log.Fatal("Invalid Trade Quantity (Valid numerical value)")
	}

	sideInput := util.ReadString("Enter trade Side ? (Buy/Sell):")

	sideInput = strings.ToUpper(sideInput)

	switch sideInput {
	case "BUY":
		side = uniswap.Buy
	case "SELL":
		side = uniswap.Sell
	default:
		log.Fatal("Invalid Side (must be 'buy' or 'sell')")
	}

	return amount, tokenInput, side
}
