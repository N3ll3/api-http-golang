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

// GetArtistsHandler récupère une liste d'artistes et les renvoie au format JSON.
func GetArtistsHandler(w http.ResponseWriter, r *http.Request) {
	// Récupère les artistes depuis la base de données
	result, err := database.GetArtists()

	// Gère les erreurs et défini le code d'état HTTP approprié
	if err != nil {
			if apiErr, ok := err.(*Errors.ApiError); ok {
					w.WriteHeader(apiErr.ResponseCode())
			} else {
					w.WriteHeader(http.StatusInternalServerError)
			}
			return
    }
	// Convertit le résultat en JSON et renvoie la réponse
	jsonStr, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(jsonStr))
}

// PostArtistHandler traite la demande d'ajout d'un nouvel artiste à la base de données.
func PostArtistHandler(w http.ResponseWriter, r *http.Request) {
	// Récupération du payload de l'artiste depuis le body de la requête
	var payload domain.Artist
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	//Controle de la validation du format de l'id de l'artiste
	if !payload.IsValid() {
		http.Error(w, "Id non valid", http.StatusBadRequest)
		return
	}
	// Ajoute l'artiste à la base de données et gère les erreurs
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

// PostArtistByURLHandler traite de la demande d'ajout d'un artiste en utilisant une URL Spotify.
func PostArtistByURLHandler(w http.ResponseWriter, r *http.Request) {
	// Récupération du payload depuis le body de la requêt
	type Payload struct {
		Spotify_url string
	}
	var payload Payload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	//1 .Parse de l'url_spotify pour récupérer l'id de l'artiste
	parsedURL, err := url.Parse(payload.Spotify_url)
	pathParts := path.Clean(parsedURL.Path)
	segments := strings.Split(pathParts, "/")
	artistId := segments[len(segments)-1]

	//2. Récupération du nom de l'artist via l'API spotify
	artistName, err := api.GetNameArtistFromSpotify(artistId) 
	if err != nil {
		if apiErr, ok := err.(*Errors.ApiError); ok {
				w.WriteHeader(apiErr.ResponseCode())
		} else {
				w.WriteHeader(http.StatusInternalServerError) 
		}
		return
	}
	//3 . Ajouter l'artiste à la base de données et gère les erreurs
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

// PostArtistTrackHandler traite la demande d'ajout d'un track au catalogue d'un artiste.
func PostArtistTrackHandler(w http.ResponseWriter, r *http.Request) {
	// Récupération du payload : trackID/name et de l'ID de l'artiste depuis les variables URL
	var payload domain.Track

	urlVars := mux.Vars(r)
	artistId := urlVars["id"]

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	//Controle de la validation du format de l'id du track
	if !payload.IsValid() {
		http.Error(w, "Id non valid", http.StatusBadRequest)
		return
	}
	// Ajoute la piste au catalogue de l'artiste et gère les erreurs
	err := database.AddArtistTrack(payload, artistId)
	if err != nil {
			if apiErr, ok := err.(*Errors.ApiError); ok {
					w.WriteHeader(apiErr.ResponseCode())
			} else {
					w.WriteHeader(http.StatusInternalServerError)
			}
			return
	}

}

// PingHandler vérifie la connexion à la base de données et répond "PONG".
func PingHandler(w http.ResponseWriter, r *http.Request) {
	database.Connection()
	w.Write([]byte("PONG"))
}
