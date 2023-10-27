package app

import (
	"api-http/app/handler"
	"api-http/app/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() {
	router := mux.NewRouter()

	router.HandleFunc("/ping", handler.PingHandler).Methods("GET")

	router.HandleFunc("/artists/", middleware.ApiKeyMiddleware(handler.GetArtistsHandler)).Methods("GET")
	router.HandleFunc("/artist", middleware.ApiKeyMiddleware(handler.PostArtistHandler)).Methods("POST")
	router.HandleFunc("/artist/url", middleware.ApiKeyMiddleware(handler.PostArtistByURLHandler)).Methods("POST")
	router.HandleFunc("/artist/{id}/track", middleware.ApiKeyMiddleware(handler.PostArtistTrackHandler)).Methods("POST")
	
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8000", nil))
}