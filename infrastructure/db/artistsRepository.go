package database

import (
	Errors "api-http/app/error"
	"api-http/domain"
	"errors"
	"log"
)

// GetArtists récupère la liste des artistes depuis la base de données.
func GetArtists() ([]domain.Artist, error) {
	log.Println("GetArtists")
	db := Connection()
	 
 // Requête SQL pour obtenir les artistes et leurs pistes associées.
	rows, err := db.Query(
		`WITH ArtistTracks AS (
    SELECT sa.name AS artist,
           sa.id AS artist_id,
           tracks.name AS track,
           tracks.id AS track_id
    FROM spotify_artist sa
    LEFT JOIN LATERAL (
        SELECT st.name, st.artist_id, st.id
        FROM spotify_track st
        WHERE st.artist_id = sa.id
        ORDER BY st.artist_id
    ) AS tracks ON sa.id = tracks.artist_id
		)
		SELECT artist_id, artist, track_id, track 
		FROM ArtistTracks;`)

	if err != nil {
		log.Printf("%v", err)
		return nil, Errors.NewApiError(errors.New("Failed to fetch artists from the database"), 500)
	}
	defer rows.Close()

	artists := make(map[string]*domain.Artist)

	for rows.Next() {
		var artistID, artistName string
		var trackID, trackName *string

		err := rows.Scan(&artistID, &artistName, &trackID, &trackName)

		if err != nil {
			log.Printf("%v", err)
			return nil, Errors.NewApiError(errors.New("Failed to fetch artists from the database"), 500)
		}

		artist, exists := artists[artistID]
		if !exists {
			artist = &domain.Artist{
				Id:     artistID,
				Name:   artistName,
				Tracks: []domain.Track{},
			}
			artists[artistID] = artist
		}

		 // Vérifier si trackID est nil avant de créer la structure Track
    if trackID != nil {
        track := domain.Track{
            Id:   *trackID,
            Name: *trackName,
        }
        artist.Tracks = append(artist.Tracks, track)
    }
	}
	defer db.Close()
	var artistList []domain.Artist
	for _, artist := range artists {
		artistList = append(artistList, *artist)
	}
	return artistList, nil
}

// AddArtist ajoute un nouvel artiste à la base de données.
func AddArtist(payload domain.Artist) error {
 log.Println("AddArtist")
	db := Connection()
	req :=`INSERT INTO spotify_artist (id , name) 
					VALUES ($1, $2)`
	_, err := db.Exec(req, payload.Id, payload.Name)

	defer db.Close()

	if err != nil {
		log.Printf("%v : %v", err, payload)
		return Errors.NewApiError(errors.New("Failed to insert artist in the database"), 422)
	}

	return nil
}

// AddArtistTrack ajoute une piste au catalogue d'un artiste.
func AddArtistTrack(payload domain.Track, artistId string) error {
 log.Println("AddArtistTrack")
	db := Connection()

	// Insertion d'un nouveau track si artist_id existe dans la table spotify_artist. 
	// Si artist_id n'existe pas, il n'y a aucune insertion.
	req := `INSERT INTO spotify_track (id, artist_id, name)
					SELECT $1, sa.id, $2
					FROM spotify_artist sa
					WHERE sa.id = $3;`

	_, err := db.Exec(req, payload.Id, payload.Name, artistId)

	defer db.Close()

	if err != nil {
		log.Printf("%v : %v", err, payload)
		return Errors.NewApiError(errors.New("Failed to insert track for the artist in the database"), 422)
	}	
	return nil
}