package main

import (
    "log"
    "net/http"
    "os"

    "Infra-Dev-Trial-Task/internal/handlers"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/api/get-balance", handlers.GetBalanceHandler)

    log.Printf("Server listening on :%s", port)
    if err := http.ListenAndServe(":"+port, mux); err != nil {
        log.Fatal(err)
    }
}
