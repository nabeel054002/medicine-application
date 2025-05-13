package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	var err error

	// Use a file-based DB (you can use ":memory:" for in-memory testing)
	DB, err = sql.Open("sqlite3", "./coupons.db")
	if err != nil {
		return fmt.Errorf("failed to open DB: %w", err)
	}

	// Optional: Enforce foreign keys
	_, err = DB.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return fmt.Errorf("failed to enable foreign keys: %w", err)
	}

	log.Println("Connected to SQLite database.")
	return nil
}
