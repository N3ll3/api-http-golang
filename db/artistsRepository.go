package database

import (
	apiError "api-http/Error"
	"api-http/domain"
	"log"
)




func GetArtists() []domain.Artist {
	log.Println("GetArtists")
	db := Connection()
	if db == nil {
    log.Fatal("Database connection is nil")
	}	
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
		SELECT artist_id, artist, track_id, track FROM ArtistTracks`)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	artists := make(map[string]*domain.Artist)

	for rows.Next() {
		var artistID, artistName, trackID, trackName string
		err := rows.Scan(&artistID, &artistName, &trackID, &trackName)
		if err != nil {
			log.Fatal(err)
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

		track := domain.Track{
			Id:   trackID,
			Name: trackName,
		}
		artist.Tracks = append(artist.Tracks, track)
	}
	defer db.Close()
	var artistList []domain.Artist
	for _, artist := range artists {
		artistList = append(artistList, *artist)
	}
	return artistList
}

func AddArtist(payload domain.Artist) (bool, apiError.ApiError) {
 log.Println("AddArtist")
	db := Connection()
	if db == nil {
    log.Fatal("Database connection is nil")
	}	

	req :=`INSERT INTO spotify_artist (id , name) 
					VALUES ($1, $2)`
	_, err := db.Exec(req, payload.Id, payload.Name)

	defer db.Close()

	if err != nil {
		return false, apiError.ApiError{
					Code :422,
					Message : "Conflit",
				}
	}

	return true, apiError.ApiError{}
}

func AddArtistTrack(payload domain.Track, artistId string) (bool, apiError.ApiError) {
 log.Println("AddArtistTrack")
	db := Connection()
	if db == nil {
    log.Println("Database connection is nil")
	}	
	
	// Insertion d'un nouveau track si artist_id existe dans la table spotify_artist. 
	// Si artist_id n'existe pas, il n'y a aucune insertion.
	req := `INSERT INTO spotify_track (id, artist_id, name)
					SELECT $1, sa.id, $2
					FROM spotify_artist sa
					WHERE sa.id = $3;`

	_, err := db.Exec(req, payload.Id, payload.Name, artistId)

	defer db.Close()

	if err != nil {
		log.Printf("%v", err)
		return false, apiError.ApiError{
					Code :422,
					Message : "Conflit",
				}
	}

	return true, apiError.ApiError{}
}