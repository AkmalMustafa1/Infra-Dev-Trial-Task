
package rpc

import (
    "context"
    "fmt"
    "os"

    "github.com/gagliardetto/solana-go"
    "github.com/gagliardetto/solana-go/rpc"
)

// client is initialized with the RPC URL from environment variable
var client *rpc.Client

func init() {
    rpcURL := os.Getenv("RPC_URL")
    if rpcURL == "" {
        rpcURL = rpc.MainNetBeta_RPC // fallback
    }
    client = rpc.New(rpcURL)
}

// GetBalance fetches the SOL balance for a given wallet address
func GetBalance(wallet string) (float64, error) {
    pubKey, err := solana.PublicKeyFromBase58(wallet)
    if err != nil {
        return 0, fmt.Errorf("invalid wallet address: %v", err)
    }

    resp, err := client.GetBalance(context.Background(), pubKey)
    if err != nil {
        return 0, fmt.Errorf("failed to fetch balance: %v", err)
    }

    // Convert lamports to SOL
    balance := float64(resp.Value) / 1e9
    return balance, nil
}
