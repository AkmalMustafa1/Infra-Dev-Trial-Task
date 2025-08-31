package rpc

import (
    "context"
    "log"
    "os"
    "time"

    "github.com/gagliardetto/solana-go"
    "github.com/gagliardetto/solana-go/rpc"
)

var c *rpc.Client

func init() {
    rpcURL := os.Getenv("RPC_URL")
    c = rpc.New(rpcURL)
}

func GetBalance(wallet string) (float64, error) {
    pubKey, err := solana.PublicKeyFromBase58(wallet)
    if err != nil {
        return 0, err
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    log.Printf("[RPC] Fetching balance for %s", wallet)
    res, err := c.GetBalance(ctx, pubKey, rpc.CommitmentProcessed)
    if err != nil {
        return 0, err
    }
    return float64(res.Value) / 1e9, nil
}
