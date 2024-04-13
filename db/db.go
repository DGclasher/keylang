package db

import (
	"database/sql"
	"log"
	 
	_ "github.com/lib/pq"
)

func NewConnection(conn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
