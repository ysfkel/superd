package wallet

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CreateSigner(ethClient *ethclient.Client) (*bind.TransactOpts, error) {

	pk, addr, err := LoadWalletKeys()

	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	chainID, err := ethClient.ChainID(ctx)
	if err != nil {
		//log.Fatal("Failed to get chain ID:", err)
		return nil, err
	}

	nonce, err := ethClient.PendingNonceAt(ctx, addr)
	if err != nil {
		// log.Fatal("Failed to get nonce:", err)
		return nil, err
	}

	gasPrice, err := ethClient.SuggestGasPrice(ctx)
	if err != nil {
		// log.Fatal("Failed to get gas price:", err)
		return nil, err
	}

	// Create transaction signer
	auth, err := bind.NewKeyedTransactorWithChainID(pk, chainID)
	if err != nil {
		//log.Fatal("Failed to create transactor:", err)
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = 500000
	auth.GasPrice = gasPrice
	auth.Context = ctx

	return auth, nil
}

func LoadWalletKeys() (*ecdsa.PrivateKey, common.Address, error) {
	privHex := os.Getenv("PRIVATE_KEY")
	if privHex == "" {
		return nil, common.Address{}, fmt.Errorf("PRIVATE_KEY not set in environment")
	}

	// Normalize and strip 0x prefix
	privHex = strings.TrimPrefix(strings.ToLower(strings.TrimSpace(privHex)), "0x")

	privKey, err := crypto.HexToECDSA(privHex)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("invalid PRIVATE_KEY: %w", err)
	}

	address := crypto.PubkeyToAddress(privKey.PublicKey)

	return privKey, address, nil
}
