package client

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func MustNewRPCClient(rpcURL string) *ethclient.Client {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect RPC: %v", err)
	}
	return client
}
