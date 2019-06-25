package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func connectDataBase() *sql.DB {
	database, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	return database
}
