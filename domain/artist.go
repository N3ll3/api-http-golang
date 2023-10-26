package domain

import (
	"log"
	"regexp"
)

type Artist struct {
	Id     string
	Name   string
	Tracks []Track
}

func (artist Artist) IsValid() bool {
	pattern := `^[0-9a-zA-Z]{22}$`
	regex, errRegex := regexp.Compile(pattern)
	if errRegex != nil {
		log.Fatal(errRegex)
	}
	isMatch := regex.MatchString(artist.Id)
	return isMatch
}