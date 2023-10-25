package main

import (
	"api-http/handler"
	"log"
	"net/http"
)

func apiKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("Api-Key")
		log.Println(apiKey)
		next.ServeHTTP(w, r)
	})
}


func main() {
		http.HandleFunc("/ping", handler.PingHandler)

		getArtistsHandler := http.HandlerFunc(handler.GetArtistsHandler)
		http.Handle("/artists/", apiKeyMiddleware(getArtistsHandler))

		postArtistHandler := http.HandlerFunc(handler.PostArtistHandler)
		http.Handle("/artist", apiKeyMiddleware(postArtistHandler))
    log.Fatal(http.ListenAndServe(":8000", nil))
}