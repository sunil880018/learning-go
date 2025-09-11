package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitPostgres(url string) *sql.DB {
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	return db
}
