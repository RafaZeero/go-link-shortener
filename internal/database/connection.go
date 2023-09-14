package database

import (
	"database/sql"

	"github.com/RafaZeero/go-link-shortener/internal/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	var err error
	DB, err = sql.Open("postgres", config.DatabaseURL)

	return err
}

func Close() {
	if DB == nil {
		return
	}
	DB.Close()
}
