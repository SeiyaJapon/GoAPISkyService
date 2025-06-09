package db

import (
	"GoAPISkyService/internal/shared_context/config"
	"database/sql"
	"log"
)

func InitDB() (*sql.DB, error) {
	dbConfig, err := config.InitializeApp()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	log.Println("Database connection established successfully")

	return dbConfig.DB, nil
}
