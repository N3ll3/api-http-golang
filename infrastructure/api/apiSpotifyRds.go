package api

import (
	Errors "api-http/app/error"
	"context"
	"errors"
	"log"
	"os"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

// GetNameArtistFromSpotify récupère le nom d'un artiste à partir de son ID Spotify.
func GetNameArtistFromSpotify(artistId string) (string, error) {
	// Récupère les informations d'identification Spotify depuis les variables d'environnement.
	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_SECRET")
	// Configure l'authentification client pour obtenir un jeton d'accès Spotify.
	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     spotify.TokenURL,
	}

	// Obtient un jeton d'accès en utilisant les informations d'identification.
	accessToken, err := config.Token(context.Background())
	if err != nil {
		log.Printf("%v", err)
		return "", Errors.NewApiError(errors.New("Failed to connect to spotify url"), 500)
	}
	// Crée un client Spotify avec le jeton d'accès.
	client := spotify.Authenticator{}.NewClient(accessToken)
	// Convertit l'ID de l'artiste en un format Spotify.
	spotifyArtistID := spotify.ID(artistId)
// Récupère les informations sur l'artiste depuis l'API Spotify. 
	spotifyArtist, err :=  client.GetArtist(spotifyArtistID)

	if err != nil {
		log.Printf("%v", err)
		return "", Errors.NewApiError(errors.New("Failed to connect to spotify url"), 500)
	}
 // Renvoie le nom de l'artiste récupéré depuis Spotify.
	return spotifyArtist.Name, nil
}
