package handler

import (
	Errors "api-http/app/error"
	"api-http/domain"
	"api-http/infrastructure/api"
	database "api-http/infrastructure/db"
	"encoding/json"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/gorilla/mux"
)

func GetArtistsHandler(w http.ResponseWriter, r *http.Request) {
	result, err := database.GetArtists()
	if err != nil {
			if apiErr, ok := err.(*Errors.ApiError); ok {
					w.WriteHeader(apiErr.ResponseCode())
			} else {
					w.WriteHeader(http.StatusInternalServerError)
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
					w.WriteHeader(http.StatusInternalServerError) 
			}
			return
    }

}


func PostArtistByURLHandler(w http.ResponseWriter, r *http.Request) {
	type Payload struct {
		Spotify_url string
	}
	var payload Payload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	//1 .parse de l'url_spotify pour recuperer l'id de l'artiste
	parsedURL, err := url.Parse(payload.Spotify_url)
	pathParts := path.Clean(parsedURL.Path)
	segments := strings.Split(pathParts, "/")
	artistId := segments[len(segments)-1]

	//2. faire un fetch du nom de l'artist sur l'API spotify
	artistName, err := api.GetNameArtistFromSpotify(artistId) 
	if err != nil {
		if apiErr, ok := err.(*Errors.ApiError); ok {
				w.WriteHeader(apiErr.ResponseCode())
		} else {
				w.WriteHeader(http.StatusInternalServerError) 
		}
		return
	}
	//3 . Ajouter l'artist
	artist := domain.Artist{
		Id: artistId,
		Name: artistName,
	}

	errAdd := database.AddArtist(artist)
	if errAdd != nil {
		if apiErr, ok := errAdd.(*Errors.ApiError); ok {
				w.WriteHeader(apiErr.ResponseCode())
		} else {
				w.WriteHeader(http.StatusInternalServerError) 
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
