package config

import (
	"fmt"
	"os"
)

type (
	database struct {
		Username string
		Password string
		Host     string
		Port     string
	}
	config struct {
		Dsn string
	}
)

func NewConfig() *config {
	return &config{
		Dsn: getDsn(),
	}
}

func getEnv() {
}

func getDsn() string {
	db := database{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%s",
		db.Username,
		db.Password,
		db.Host,
		db.Port,
	)
}
