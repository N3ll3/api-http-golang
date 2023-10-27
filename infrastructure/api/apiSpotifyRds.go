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

type GETSpotifyResponse struct {
  Id string
  Name string
}

func GetNameArtistFromSpotify(artistId string) (string, error) {
	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_SECRET")
	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     spotify.TokenURL,
	}

	accessToken, err := config.Token(context.Background())
	if err != nil {
		log.Printf("%v", err)
		return "", Errors.NewApiError(errors.New("Failed to connect to spotify url"), 500)
	}

	client := spotify.Authenticator{}.NewClient(accessToken)
	spotifyArtistID := spotify.ID(artistId)

	spotifyArtist, err :=  client.GetArtist(spotifyArtistID)

	if err != nil {
		log.Printf("%v", err)
		return "", Errors.NewApiError(errors.New("Failed to connect to spotify url"), 500)
	}

	return spotifyArtist.Name, nil
}
