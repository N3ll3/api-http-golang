package db

import (
	"database/sql"
	"fmt"
	"os"
)

var db *sql.DB

func Connection() {
 var (
  host     = os.Getenv("DBHOST")
  port     = os.Getenv("DBPORT")
  user     = os.Getenv("DBUSER")
  password = os.Getenv("DBPASS")
  dbname   = os.Getenv("DBNAME")
	sslmode  = os.Getenv("DBSSLMODE")
)
	 postgresqlDbInfo := fmt.Sprintf("host=%s port=%s user=%s "+
    "password=%s dbname=%s sslmode=%s",
    host, port, user, password, dbname, sslmode) 
	db, err := sql.Open("postgres", postgresqlDbInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()
  err = db.Ping()
  if err != nil {
    panic(err)
  }
  fmt.Println("Established a successful connection!")
}