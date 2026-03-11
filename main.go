package main

import (
    "log"
    "net/http"
    "bitrail/handlers"
    "bitrail/middleware"
)

func main() {
    mux := http.NewServeMux()
    
    mux.HandleFunc("/api/trade/ws", middleware.SessionAuth(handlers.HandleWebSocket))
    mux.HandleFunc("/api/orders", middleware.SessionAuth(handlers.GetOrders))
    mux.HandleFunc("/api/balance", middleware.SessionAuth(handlers.GetBalance))
    
    log.Println("BitRail trading server starting on :8443")
    log.Fatal(http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", mux))
}
