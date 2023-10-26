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

func Connection() *sql.DB {
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
  // defer db.Close()
  err = db.Ping()
  if err != nil {
    log.Fatal(err)
  }
  log.Println("Established a successful connection!")
	return db;
}

