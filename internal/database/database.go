package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

// InitDB initializes the database connection.
func InitDB(dbConnectionString string) (*DB, error) {
	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return nil, err
	}
	// Verify connection.
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
		return nil, err
	}
	return &DB{db}, nil
}

// CloseDB closes the database connection.
func (db *DB) CloseDB() {
	if err := db.DB.Close(); err != nil {
		log.Fatalf("failed to close database connection: %v", err)
	}
}
