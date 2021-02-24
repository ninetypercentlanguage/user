package db

import (
	"database/sql"
	"log"

	// package init registers postgres driver
	_ "github.com/lib/pq"
)

func InitDB(connection string) *sql.DB {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
