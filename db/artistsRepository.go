package database

import (
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

func AddArtist(payload domain.Artist) {
 log.Println("AddArtist")
	db := Connection()
	if db == nil {
    log.Fatal("Database connection is nil")
	}	
	//verifier s'il existe ?
	//verifier format de id
	req :=`INSERT INTO spotify_artist (id , name) 
					VALUES ($1, $2)`
	_, err := db.Exec(req, payload.Id, payload.Name)

	if err != nil {
		log.Fatal(err)
	}
}