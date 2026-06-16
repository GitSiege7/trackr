package main

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/GitSiege7/trackr/database"
	_ "github.com/ncruces/go-sqlite3/driver"
)

func OpenDB() (*database.Queries, error) {
	CONFIG_PATH, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	DB_PATH := filepath.Join(CONFIG_PATH, "/trackr/times.db")

	db, err := sql.Open("sqlite3", DB_PATH)
	if err != nil {
		return nil, err
	}

	queries := database.New(db)

	return queries, nil
}
