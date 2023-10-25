package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Connection() {
	log.Println("Preparation de la connexion DB")
	godotenv.Load(".env")
 var (
  host     = os.Getenv("DBHOST")
  port     = os.Getenv("DBPORT")
  user     = os.Getenv("DBUSER")
  password = os.Getenv("DBPASS")
  dbname   = os.Getenv("DBNAME")
	sslmode  = os.Getenv("DBSSLMODE")
	certifPath 	 = os.Getenv("CERTIFCAPATH")
	)
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%s user=%s "+
    "password=%s dbname=%s sslmode=%s sslrootcert=%s",
    host, port, user, password, dbname, sslmode, certifPath) 

	db, err := sql.Open("postgres", postgresqlDbInfo)
  if err != nil {
    log.Fatal(err)
  }
	if db == nil {
    log.Fatal("Database connection is nil")
	}	
  defer db.Close()
  err = db.Ping()
  if err != nil {
    log.Fatal(err)
  }
  log.Println("Established a successful connection!")
}

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
	log.Println("GetArtists")
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