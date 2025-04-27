package config

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

type AppConfig struct {
	DB   *sql.DB
	Port string
	Env  string
}

func InitializeApp() (*AppConfig, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: No .env file found")
	}

	dbConfig, err := GetDBConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load database config: %v", err)
	}

	db, err := dbConfig.ConnectDB()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database after connection: %v", err)
	}

	return &AppConfig{
		DB:   db,
		Port: getEnvWithDefault("PORT", "8080"),
		Env:  getEnvWithDefault("ENV", "development"),
	}, nil
}
