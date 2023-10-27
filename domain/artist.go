package domain

import (
	"log"
	"regexp"
)

type Artist struct {
	Id     string // Identifiant de l'artiste
	Name   string // Nom de l'artiste
	Tracks []Track // Liste de pistes associées à l'artiste
}

// IsValid vérifie si l'identifiant de l'artiste est valide en utilisant une expression régulière.
func (artist Artist) IsValid() bool {
	pattern := `^[0-9a-zA-Z]{22}$`
	regex, errRegex := regexp.Compile(pattern)
	if errRegex != nil {
		log.Fatal(errRegex)
	}
	isMatch := regex.MatchString(artist.Id)
	return isMatch
}