package helpers

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/GitSiege7/trackr/database"
	_ "github.com/ncruces/go-sqlite3/driver"
	"github.com/pressly/goose"
)

func migrateUp(db *sql.DB) error {
	goose.SetDialect("sqlite3")
	err := goose.Up(db, "../../sql/schema")
	if err != nil {
		return err
	}

	return nil
}

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

	err = migrateUp(db)
	if err != nil {
		return nil, err
	}

	queries := database.New(db)

	return queries, nil
}
