package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func CreateDBClient() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@postgres:5432/test_db?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
