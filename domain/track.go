package domain

import (
	"log"
	"regexp"
)

type Track struct {
	Id   string // Identifiant du track
	Name string // Nom du track
}

// IsValid vérifie si l'identifiant du track est valide en utilisant une expression régulière.
func (track Track) IsValid() bool {
	pattern := `^[0-9a-zA-Z]{22}$`
	regex, errRegex := regexp.Compile(pattern)
	if errRegex != nil {
		log.Fatal(errRegex)
	}
	isMatch := regex.MatchString(track.Id)
	return isMatch
}