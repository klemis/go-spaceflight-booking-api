package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
)

func main() {
	//dbConnectionString := "host=localhost port=5432 user=admin password=admin dbname=bookings sslmode=disable"
	databaseURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("failed to close database connection: %v", err)
		}
	}(db)

	migrationsDir := "internal/database/migrations"

	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		log.Fatalf("failed to read migrations directory: %v", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			migrationScript, err := os.ReadFile(filepath.Join(migrationsDir, file.Name()))
			if err != nil {
				log.Fatalf("failed to read migration file %s: %v", file.Name(), err)
			}

			_, err = db.Exec(string(migrationScript))
			if err != nil {
				log.Fatalf("failed to execute migration %s: %v", file.Name(), err)
			}

			log.Printf("Successfully executed migration: %s", file.Name())
		}
	}

	log.Println("All migrations executed successfully.")
}
