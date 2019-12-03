package database

import (
	"database/sql"
	"github.com/ipan97/go-dig-sample/config"
	_ "github.com/mattn/go-sqlite3"
)

func Connect(config *config.Config) (*sql.DB, error) {
	return sql.Open("sqlite3", config.DatabasePath)
}
