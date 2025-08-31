
package rpc

import (
    "context"
    "fmt"
    "os"

    "github.com/gagliardetto/solana-go"
    "github.com/gagliardetto/solana-go/rpc"
)

var client *rpc.Client

func init() {
    rpcURL := os.Getenv("RPC_URL")
    if rpcURL == "" {
        rpcURL = rpc.MainNetBeta_RPC
    }
    client = rpc.New(rpcURL)
}

// GetBalance fetches the SOL balance for a given wallet address
func GetBalance(wallet string) (float64, error) {
    pubKey, err := solana.PublicKeyFromBase58(wallet)
    if err != nil {
        return 0, fmt.Errorf("invalid wallet address: %v", err)
    }

    // Add Commitment parameter
    resp, err := client.GetBalance(
        context.Background(),
        pubKey,
        rpc.CommitmentFinalized, // Use Finalized commitment
    )
    if err != nil {
        return 0, fmt.Errorf("failed to fetch balance: %v", err)
    }

    balance := float64(resp.Value) / 1e9 // lamports → SOL
    return balance, nil
}
