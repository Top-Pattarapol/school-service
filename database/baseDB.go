package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func connect() *sql.DB {
	database, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	return database
}

func baseExec(db *sql.DB, query string, args ...interface{}) error {
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(args...)
	if err != nil {
		return err
	}
	return nil
}

func baseQuery(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	return rows, err
}

func baseQueryRow(db *sql.DB, query string, args ...interface{}) (*sql.Row, error) {
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(args...)
	return row, nil
}
