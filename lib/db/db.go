package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	// Db, err = sql.Open("postgres", "dbname=fulcrum sslmode=disable")
	Db, err = sql.Open("postgres", "dbname=fulcrum")
	if err != nil {
		log.Fatal(err)
	}
}
