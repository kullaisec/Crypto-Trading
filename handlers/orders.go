package handlers

import (
    "encoding/json"
    "net/http"
    "bitrail/database"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
    userID := r.Header.Get("X-User-ID")
    
    orders, err := database.FetchUserOrders(userID)
    if err != nil {
        http.Error(w, "Failed to fetch orders", http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(orders)
}
