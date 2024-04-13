package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/DGclasher/keylang/cmd/api"
	"github.com/DGclasher/keylang/config"
	"github.com/DGclasher/keylang/db"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Envs.DBHost, strconv.Itoa(config.Envs.DBPort), config.Envs.DBUser, config.Envs.DBPass, config.Envs.DBName)

	db, err := db.NewConnection(psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	initStorage(db)
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	err = createTables(db)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection established")
}

func createTables(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS licenses (
			id SERIAL PRIMARY KEY,
			user_email TEXT NOT NULL,
			product_id TEXT NOT NULL,
			license_key TEXT NOT NULL
		);
	`)
	if err != nil {
		return err
	}
	return nil
}
