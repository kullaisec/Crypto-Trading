package middleware

import (
    "net/http"
    "bitrail/session"
)

func SessionAuth(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        cookie, err := r.Cookie("bitrail_session")
        if err != nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        
        userID, valid := session.ValidateSession(cookie.Value)
        if !valid {
            http.Error(w, "Invalid session", http.StatusUnauthorized)
            return
        }
        
        r.Header.Set("X-User-ID", userID)
        next(w, r)
    }
}
