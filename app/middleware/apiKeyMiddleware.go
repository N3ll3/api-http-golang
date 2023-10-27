package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// ApiKeyMiddleware est un middleware qui vérifie la présence et la validité de la clé API dans l'en-tête de la requête.
func ApiKeyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Charge les variables d'environnement à partir du fichier .env
        godotenv.Load(".env")
        // Récupère le hachage attendu de la clé API depuis les variables d'environnement
        expectedHash := os.Getenv("API_KEY_SHA256")
        // Récupère la clé API de l'en-tête de la requête
        apiKey := r.Header.Get("Api-Key")
        if apiKey == "" {
            http.Error(w, "Missing API key", http.StatusUnauthorized)
            return
        }
        
        //  Hache la clé API fournie
        hash := sha256.Sum256([]byte(apiKey))
        apiKeyHash := hex.EncodeToString(hash[:])

        // Vérifie si le hachage de la clé API correspond au hachage attendu
        if apiKeyHash != expectedHash {
            http.Error(w, "Invalid API key", http.StatusUnauthorized)
            return
        }
        // Passe la requête au gestionnaire suivant si la clé API est valide      
        next.ServeHTTP(w, r)
	})
}
