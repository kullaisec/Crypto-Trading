package handlers

import (
    "net/http"
    "log"
    "github.com/gorilla/websocket"
    "bitrail/trading"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Printf("WebSocket upgrade failed: %v", err)
        return
    }
    defer conn.Close()
    
    userID := r.Header.Get("X-User-ID")
    log.Printf("User %s connected to trading WebSocket", userID)
    
    trading.HandleTrading(conn, userID)
}
