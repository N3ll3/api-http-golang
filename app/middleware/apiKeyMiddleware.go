package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func ApiKeyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		godotenv.Load(".env")
	  expectedHash := os.Getenv("API_KEY_SHA256")
		apiKey := r.Header.Get("Api-Key")
        if apiKey == "" {
            http.Error(w, "Missing API key", http.StatusUnauthorized)
            return
        }
        
        // Hash the provided API key
        hash := sha256.Sum256([]byte(apiKey))
        apiKeyHash := hex.EncodeToString(hash[:])
        
        if apiKeyHash != expectedHash {
            http.Error(w, "Invalid API key", http.StatusUnauthorized)
            return
        }      
        next.ServeHTTP(w, r)
	})
}
