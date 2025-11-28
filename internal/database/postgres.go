package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgresDB(databseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databseURL)

	if err != nil {
		return nil, fmt.Errorf("error opening db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to db: %w", err)
	}

	log.Println("DB connected successfully")

	return  db, nil;
}