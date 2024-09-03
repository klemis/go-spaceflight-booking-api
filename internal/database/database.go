package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/klemis/go-spaceflight-booking-api/models"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// DBInterface defines the methods related to database operations.
type DBInterface interface {
	GetDestinationID(launchpadID string, launchDate time.Time) (models.Destination, error)
	GetLaunchpadID(destinationID models.Destination, launchDate time.Time) (string, error)
	InsertBooking(request models.BookingRequest, launchpadID string) (uint, error)
	GetBookings() ([]models.Booking, error)
	DeleteBooking(id int) error
}

// DB is a wrapper around sql.DB that implements DBInterface.
type DB struct {
	*sql.DB
}

// NewDB creates a new instance of DB.
func NewDB(db *sql.DB) *DB {
	return &DB{DB: db}
}

func (db *DB) GetBookings() ([]models.Booking, error) {
	var bookings []models.Booking
	query := `SELECT id, first_name, last_name, gender, birthday, launchpad_id, destination_id, launch_date FROM bookings;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal("failed to close rows in GetBookings query")
		}
	}(rows)

	for rows.Next() {
		var booking models.Booking
		err = rows.Scan(&booking.ID, &booking.FirstName, &booking.LastName, &booking.Gender, &booking.Birthday, &booking.LaunchpadID, &booking.DestinationID, &booking.LaunchDate)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}

	if len(bookings) == 0 {
		return nil, sql.ErrNoRows
	}

	return bookings, nil
}

func (db *DB) DeleteBooking(id int) error {
	query := `DELETE FROM bookings WHERE id = $1;`
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (db *DB) GetDestinationID(launchpadID string, launchDate time.Time) (models.Destination, error) {
	query := `SELECT destination_id FROM schedules WHERE launchpad_id = $1 AND day_of_week = $2;`
	row := db.QueryRow(query, launchpadID, launchDate.Weekday())

	var schedule models.Schedule
	err := row.Scan(&schedule.Destination)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("missing destination for the provided launchpad at this date")
		}

		return 0, err
	}

	return schedule.Destination, nil
}

func (db *DB) GetLaunchpadID(destinationID models.Destination, launchDate time.Time) (string, error) {
	query := `SELECT launchpad_id FROM schedules WHERE destination_id = $1 AND day_of_week = $2;`
	row := db.QueryRow(query, destinationID, launchDate.Weekday())

	var schedule models.Schedule
	err := row.Scan(&schedule.LaunchpadID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("missing launchpad for the provided destination at this date")
		}

		return "", err
	}

	return schedule.LaunchpadID, nil
}

func (db *DB) InsertBooking(request models.BookingRequest, launchpadID string) (uint, error) {
	query := `
        INSERT INTO bookings (first_name, last_name, gender, birthday, launchpad_id, destination_id, launch_date)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id`

	var id uint
	err := db.QueryRow(query,
		request.FirstName,
		request.LastName,
		request.Gender,
		request.Birthday,
		launchpadID,
		request.DestinationID,
		request.LaunchDate,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to insert booking: %w", err)
	}

	return id, nil
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
