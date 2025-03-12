package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"go-etl/config"
)

var DB *sql.DB

func ConnectDB() {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.AppConfig.DBHost, config.AppConfig.DBUser, config.AppConfig.DBPassword, config.AppConfig.DBName)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	fmt.Println("Connected to PostgreSQL!")
}

// SaveRawData inserts raw JSON into PostgreSQL
func SaveRawData(jsonData string) error {
	_, err := DB.Exec("INSERT INTO raw_data (data) VALUES ($1)", jsonData)
	return err
}
