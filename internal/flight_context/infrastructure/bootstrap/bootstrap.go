package bootstrap

import (
	"GoAPISkyService/internal/flight_context/infrastructure/bootstrap/db"
	"database/sql"
)

type AppDependencies struct {
	DB *sql.DB
}

func InitDependencies() (*AppDependencies, error) {
	dbConfig, err := db.InitDB()
	if err != nil {
		return nil, err
	}

	return &AppDependencies{
		DB: dbConfig,
	}, nil
}
