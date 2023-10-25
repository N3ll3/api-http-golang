package db

import "log"

type Artist struct {
	Id     string
	Name   string
	Tracks []Track
}

type Track struct {
	Id   string
	Name string
}

func GetArtists() []Artist {
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

	artists := make(map[string]*Artist)

	for rows.Next() {
		var artistID, artistName, trackID, trackName string
		err := rows.Scan(&artistID, &artistName, &trackID, &trackName)
		if err != nil {
			log.Fatal(err)
		}

		artist, exists := artists[artistID]
		if !exists {
			artist = &Artist{
				Id:     artistID,
				Name:   artistName,
				Tracks: []Track{},
			}
			artists[artistID] = artist
		}

		track := Track{
			Id:   trackID,
			Name: trackName,
		}
		artist.Tracks = append(artist.Tracks, track)
	}

	var artistList []Artist
	for _, artist := range artists {
		artistList = append(artistList, *artist)
	}

	return artistList
}