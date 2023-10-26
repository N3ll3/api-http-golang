package domain

import (
	"log"
	"regexp"
)

type Track struct {
	Id   string
	Name string
}

func (track Track) IsValid() bool {
	pattern := `^[0-9a-zA-Z]{22}$`
	regex, errRegex := regexp.Compile(pattern)
	if errRegex != nil {
		log.Fatal(errRegex)
	}
	isMatch := regex.MatchString(track.Id)
	return isMatch
}