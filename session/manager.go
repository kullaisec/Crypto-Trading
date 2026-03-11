package session

import (
    "sync"
    "time"
)

var sessions = &sync.Map{}

func ValidateSession(sessionID string) (string, bool) {
    val, ok := sessions.Load(sessionID)
    if !ok {
        return "", false
    }
    
    sess := val.(Session)
    if time.Now().After(sess.Expiry) {
        sessions.Delete(sessionID)
        return "", false
    }
    
    return sess.UserID, true
}

type Session struct {
    UserID string
    Expiry time.Time
}
Total
