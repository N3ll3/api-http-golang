package main

import (
	"api-http/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func apiKeyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("Api-Key")
		log.Println(apiKey)
		next(w, r)
	})
}


func main() {
	router := mux.NewRouter()

	router.HandleFunc("/ping", handler.PingHandler).Methods("GET")

	router.HandleFunc("/artists/", apiKeyMiddleware(handler.GetArtistsHandler)).Methods("GET")
	router.HandleFunc("/artist", apiKeyMiddleware(handler.PostArtistHandler)).Methods("POST")
	router.HandleFunc("/artist/{id}/track", apiKeyMiddleware(handler.PostArtistTrackHandler)).Methods("POST")
	
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8000", nil))
}