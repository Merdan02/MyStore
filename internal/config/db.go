package config

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var DB *sql.DB

func ConnectDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	ConnStr := fmt.Sprintf("host =%s port=%s user=%s dbname=%s password=%s", host, port, user, dbname, password)

	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("✅ Successfully connected to DB")

	DB = db

	return DB, nil

}
