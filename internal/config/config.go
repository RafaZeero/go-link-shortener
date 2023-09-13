package config

import (
	"os"
)

var DatabaseURL string

func init() {
	DatabaseURL = envOrFatal("DATABASE_URL")
}

func envOrFatal(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("missing required environment variable " + key)
	}
	return value
}
