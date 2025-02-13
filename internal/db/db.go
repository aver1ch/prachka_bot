package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"laundryBot/internal/errs"

	_ "github.com/lib/pq"
)

func ConnectToDB() (*sql.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", errs.ErrConnectionToDB, err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", errs.ErrConnectionToDB, err)
	}

	log.Println("Подключение к PostgreSQL установлено!")
	return db, nil
}

func InsertUserToDB(db *sql.DB, username, roomNumber string, isAuthorised bool) error {
	query := "INSERT INTO users (username, room_number, is_authorised) VALUES ($1, $2, $3)"
	_, err := db.Exec(query, username, roomNumber, isAuthorised)
	if err != nil {
		return fmt.Errorf("%w: %w", err, errs.ErrInsertingDataFromDB)
	}

	log.Printf("Пользователь %s с номером комнаты %s успешно добавлен в базу", username, roomNumber)
	return nil
}

func GetIsAuthorisedFromDB(db *sql.DB, username string) (bool, error) {
	var isAuthorised bool
	query := "SELECT is_authorised FROM users WHERE username = $1"
	err := db.QueryRow(query, username).Scan(&isAuthorised)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("%w: %w", err, errs.ErrPullingDataFromDB)
	}

	return isAuthorised, nil
}
