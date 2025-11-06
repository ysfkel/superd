package main

import (
	"fmt"

	"github.com/ykel/quote_engine/cmd/cli"
	"github.com/ykel/quote_engine/core/config"
)

func main() {
	config.MustLoad()

	fmt.Println("=== Configuration Loaded ===")
	fmt.Printf("RPC URL: %s\n", config.App.RPCURL)
	fmt.Printf("Chain ID: %d\n", config.App.ChainID)
	fmt.Printf("Router: %s\n", config.App.RouterAddress.Hex())
	fmt.Printf("Factory: %s\n", config.App.FactoryAddress.Hex())
	fmt.Printf("Quoter: %s\n", config.App.QuoterAddress.Hex())
	fmt.Printf("WETH: %s\n", config.App.WETHAddress.Hex())
	fmt.Printf("USDC: %s\n", config.App.USDCAddress)
	fmt.Printf("USDT: %s\n", config.App.USDTAddress)
	fmt.Printf("Default Fee: %d (%.2f%%)\n\n", config.App.DefaultFee, float64(config.App.DefaultFee)/10000)

	cli.RunCli()

}
