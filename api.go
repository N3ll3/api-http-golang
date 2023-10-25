package main

import (
	db "api-http/DB"
	"encoding/json"
	"log"
	"net/http"
)

func getArtistsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
  }
	result := db.GetArtists();
	jsonStr, err := json.Marshal(result)
	if err != nil {
		 w.WriteHeader(http.StatusInternalServerError)
     return
	}
  w.Write([]byte(jsonStr))
}

func apiKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("Api-Key")
		 w.Write([]byte(apiKey))
		next.ServeHTTP(w, r)
	})
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	 w.Write([]byte("PONG"))
}


func main() {
		http.HandleFunc("/ping", pingHandler)
		artistsHandler := http.HandlerFunc(getArtistsHandler)
		http.Handle("/artists/", apiKeyMiddleware(artistsHandler))
    log.Fatal(http.ListenAndServe(":8000", nil))
}