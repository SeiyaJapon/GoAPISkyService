package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Database *DatabaseConfig `mapstructure:"database"`
	Port     string          `mapstructure:"port"`
	Env      string          `mapstructure:"env"`
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning")
	}

	dbConfig, err := GetDBConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		Database: dbConfig,
		Port:     getEnvWithDefault("PORT", "8080"),
		Env:      getEnvWithDefault("ENV", "development"),
	}, nil
}

func getEnvWithDefault(key, defaultVaule string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultVaule
}
