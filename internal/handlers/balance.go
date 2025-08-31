package handlers

import (
    "encoding/json"
    "net/http"

    "Infra-Dev-Trial-Task/internal/auth"
    "Infra-Dev-Trial-Task/internal/cache"
    "Infra-Dev-Trial-Task/internal/ratelimit"
    "Infra-Dev-Trial-Task/internal/rpc"
)

type RequestBody struct {
    Wallets []string `json:"wallets"`
}

type ResponseBody struct {
    Wallet  string  `json:"wallet"`
    Balance float64 `json:"balance"`
    Error   string  `json:"error,omitempty"`
}

func GetBalanceHandler(w http.ResponseWriter, r *http.Request) {
    apiKey := r.Header.Get("X-API-KEY")
    if !auth.ValidateAPIKey(apiKey) {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    ip := r.RemoteAddr
    if !ratelimit.Allow(ip) {
        http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
        return
    }

    var req RequestBody
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }

    results := make([]ResponseBody, len(req.Wallets))
    for i, wallet := range req.Wallets {
        bal, err := cache.GetOrFetch(wallet, func() (float64, error) {
            return rpc.GetBalance(wallet)
        })
        if err != nil {
            results[i] = ResponseBody{Wallet: wallet, Error: err.Error()}
        } else {
            results[i] = ResponseBody{Wallet: wallet, Balance: bal}
        }
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(results)
}
