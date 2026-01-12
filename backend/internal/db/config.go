package db

import (
	"fmt"
	"os"
)

type Config struct {
	DatabaseURL string
}

func LoadConfig() (*Config, error) {
	variableEnvironnementBD := "MIMI_DATABASE_URL"
	dbURL := os.Getenv(variableEnvironnementBD)
	if dbURL == "" {
		return nil, fmt.Errorf("La variable \"%s\" n'existe pas sur votre poste", variableEnvironnementBD)
	}

	return &Config{
		DatabaseURL: dbURL,
	}, nil
}
