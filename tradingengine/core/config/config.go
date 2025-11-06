package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	"github.com/ykel/quote_engine/core/wallet"
)

// Config holds all application configuration
type Config struct {
	// Network
	RPCURL  string
	ChainID int64

	// Wallet
	WalletAddress common.Address

	// Uniswap V3 Addresses
	RouterAddress  common.Address
	FactoryAddress common.Address
	QuoterAddress  common.Address

	// Token Addresses
	WETHAddress common.Address
	USDCAddress common.Address
	USDTAddress common.Address

	// Defaults
	DefaultFee      int64 // basis points (3000 = 0.3%)
	DefaultDeadline int64 // minutes
	DefaultGasLimit uint64
	DefaultSlippage float64 // 0.01 = 1%
}

var (
	// Global config instance
	App *Config
)

// Load reads environment variables and initializes the global config
func Load() error {
	// Load .env file (ignore error if it doesn't exist)
	_ = godotenv.Load()

	// Get all required values (will fail if not set)
	rpcURL, err := getEnvRequired("RPC_URL")
	if err != nil {
		return err
	}

	_, walletAddress, err := wallet.LoadWalletKeys()
	if err != nil {
		return err
	}

	chainID, err := getEnvInt64Required("CHAIN_ID")
	if err != nil {
		return err
	}

	routerAddr, err := getEnvRequired("ROUTER_ADDRESS")
	if err != nil {
		return err
	}

	factoryAddr, err := getEnvRequired("FACTORY_ADDRESS")
	if err != nil {
		return err
	}

	quoterAddr, err := getEnvRequired("QUOTER_ADDRESS")
	if err != nil {
		return err
	}

	wethAddr, err := getEnvRequired("WETH_ADDRESS")
	if err != nil {
		return err
	}

	usdcAddr, err := getEnvRequired("USDC_ADDRESS")
	if err != nil {
		return err
	}

	usdtAddr, err := getEnvRequired("USDT_ADDRESS")
	if err != nil {
		return err
	}
	defaultFee, err := getEnvInt64Required("DEFAULT_FEE")
	if err != nil {
		return err
	}

	defaultDeadline, err := getEnvInt64Required("DEFAULT_DEADLINE_MINUTES")
	if err != nil {
		return err
	}

	defaultGasLimit, err := getEnvUint64Required("DEFAULT_GAS_LIMIT")
	if err != nil {
		return err
	}

	defaultSlippage, err := getEnvFloat64Required("DEFAULT_SLIPPAGE")
	if err != nil {
		return err
	}

	App = &Config{
		// Network
		RPCURL:        rpcURL,
		ChainID:       chainID,
		WalletAddress: walletAddress,

		// Wallet
		// PrivateKey: strings.TrimPrefix(privateKey, "0x"),

		// Uniswap V3 Addresses
		RouterAddress:  common.HexToAddress(routerAddr),
		FactoryAddress: common.HexToAddress(factoryAddr),
		QuoterAddress:  common.HexToAddress(quoterAddr),

		// Token Addresses
		WETHAddress: common.HexToAddress(wethAddr),
		USDCAddress: common.HexToAddress(usdcAddr),
		USDTAddress: common.HexToAddress(usdtAddr),
		// Defaults
		DefaultFee:      defaultFee,
		DefaultDeadline: defaultDeadline,
		DefaultGasLimit: defaultGasLimit,
		DefaultSlippage: defaultSlippage,
	}

	// Validate configuration
	if err := App.Validate(); err != nil {
		return err
	}

	return nil
}

// MustLoad loads config and panics on error
func MustLoad() {
	if err := Load(); err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}
}

// Validate checks that required configuration is present and valid
func (c *Config) Validate() error {

	if c.RPCURL == "" {
		return fmt.Errorf("RPC_URL is required")
	}
	if c.WalletAddress.String() == "" {
		return fmt.Errorf("AccountAddress is required")
	}

	if c.RouterAddress == (common.Address{}) {
		return fmt.Errorf("ROUTER_ADDRESS is invalid")
	}
	if c.FactoryAddress == (common.Address{}) {
		return fmt.Errorf("FACTORY_ADDRESS is invalid")
	}
	if c.WETHAddress == (common.Address{}) {
		return fmt.Errorf("WETH_ADDRESS is invalid")
	}
	return nil
}

// Helper functions - ALL REQUIRED, NO DEFAULTS

// getEnvRequired gets an environment variable or returns an error if not set
func getEnvRequired(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("%s environment variable is required but not set", key)
	}
	return value, nil
}

// getEnvInt64Required gets an int64 environment variable or returns an error
func getEnvInt64Required(key string) (int64, error) {
	value := os.Getenv(key)
	if value == "" {
		return 0, fmt.Errorf("%s environment variable is required but not set", key)
	}
	result, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%s must be a valid integer: %w", key, err)
	}
	return result, nil
}

// getEnvUint64Required gets a uint64 environment variable or returns an error
func getEnvUint64Required(key string) (uint64, error) {
	value := os.Getenv(key)
	if value == "" {
		return 0, fmt.Errorf("%s environment variable is required but not set", key)
	}
	result, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%s must be a valid unsigned integer: %w", key, err)
	}
	return result, nil
}

// getEnvFloat64Required gets a float64 environment variable or returns an error
func getEnvFloat64Required(key string) (float64, error) {
	value := os.Getenv(key)
	if value == "" {
		return 0, fmt.Errorf("%s environment variable is required but not set", key)
	}
	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("%s must be a valid float: %w", key, err)
	}
	return result, nil
}
