package api

import (
	Errors "api-http/app/error"
	"errors"
	"io"
	"log"
	"net/http"
)

func GetNameArtistFromSpotify(url string) (string, error) {
	response, err := http.Get(url)
		if err != nil {
			log.Printf("%v", err)
			return "", Errors.NewApiError(errors.New("Failed to connect to spotify url"), 500)
		}
		defer response.Body.Close()

		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			log.Printf("%v", err)
			return "", Errors.NewApiError(errors.New("Failed to fetch name of artist from spotify url"), 500)
		}
		return string(responseBody), nil
}