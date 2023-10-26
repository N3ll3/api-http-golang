package main

import (
	"api-http/api/handler"
	"api-http/api/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/ping", handler.PingHandler).Methods("GET")

	router.HandleFunc("/artists/", middleware.ApiKeyMiddleware(handler.GetArtistsHandler)).Methods("GET")
	router.HandleFunc("/artist", middleware.ApiKeyMiddleware(handler.PostArtistHandler)).Methods("POST")
	router.HandleFunc("/artist/{id}/track", middleware.ApiKeyMiddleware(handler.PostArtistTrackHandler)).Methods("POST")
	
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8000", nil))
}