package config

import (
	"fmt"
	"os"
)

var DatabaseURL string

func init() {
	host := envOrFatal("DB_HOST")
	port := envOrFatal("DB_PORT")
	user := envOrFatal("DB_USER")
	password := envOrFatal("DB_PASSWORD")
	dbname := envOrFatal("DB_NAME")
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	DatabaseURL = psqlconn
}

func envOrFatal(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("missing required environment variable " + key)
	}
	return value
}
