package handler

import (
	Errors "api-http/api/error"
	database "api-http/db"
	"api-http/domain"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetArtistsHandler(w http.ResponseWriter, r *http.Request) {
	result, err := database.GetArtists()
	if err != nil {
			if apiErr, ok := err.(*Errors.ApiError); ok {
					w.WriteHeader(apiErr.ResponseCode())
			} else {
					// Handle the case where the error is not an ApiError
					w.WriteHeader(http.StatusInternalServerError) // Set a default status code
			}
			return
    }
	jsonStr, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(jsonStr))
}

func PostArtistHandler(w http.ResponseWriter, r *http.Request) {
	var payload domain.Artist

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	if !payload.IsValid() {
		http.Error(w, "Id non valid", http.StatusBadRequest)
		return
	}
	err := database.AddArtist(payload)
		if err != nil {
			if apiErr, ok := err.(*Errors.ApiError); ok {
					w.WriteHeader(apiErr.ResponseCode())
			} else {
					// Handle the case where the error is not an ApiError
					w.WriteHeader(http.StatusInternalServerError) // Set a default status code
			}
			return
    }

}

func PostArtistTrackHandler(w http.ResponseWriter, r *http.Request) {
	var payload domain.Track

	urlVars := mux.Vars(r)
	artistId := urlVars["id"]

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	if !payload.IsValid() {
		http.Error(w, "Id non valid", http.StatusBadRequest)
		return
	}
	err := database.AddArtistTrack(payload, artistId)
	if err != nil {
			if apiErr, ok := err.(*Errors.ApiError); ok {
					w.WriteHeader(apiErr.ResponseCode())
			} else {
					// Handle the case where the error is not an ApiError
					w.WriteHeader(http.StatusInternalServerError) // Set a default status code
			}
			return
	}

}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	database.Connection()
	w.Write([]byte("PONG"))
}
