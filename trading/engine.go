package trading

import (
    "encoding/json"
    "log"
    "github.com/gorilla/websocket"
    "bitrail/database"
)

type TradeOrder struct {
    Action string  `json:"action"`
    Symbol string  `json:"symbol"`
    Amount float64 `json:"amount"`
    Price  float64 `json:"price"`
}

func HandleTrading(conn *websocket.Conn, userID string) {
    for {
        var order TradeOrder
        err := conn.ReadJSON(&order)
        if err != nil {
            log.Printf("Read error: %v", err)
            break
        }
        
        result := database.ExecuteTrade(userID, order)
        conn.WriteJSON(result)
    }
}
