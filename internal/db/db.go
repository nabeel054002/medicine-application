package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

var DB *sql.DB

// InitDB initializes the database and creates the necessary tables
func InitDB() error {
	// Open a connection to the SQLite database (this will create the database file if it doesn't exist)
	var err error
	DB, err = sql.Open("sqlite3", "./coupons.db")
	if err != nil {
		return fmt.Errorf("could not open database: %v", err)
	}

	// Create the coupons table if it doesn't already exist
	query := `
	CREATE TABLE IF NOT EXISTS coupons (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		coupon_code TEXT UNIQUE NOT NULL,
		expiry_date TEXT NOT NULL,
		usage_type TEXT NOT NULL,
		applicable_medicine_ids TEXT,
		applicable_categories TEXT,
		min_order_value REAL,
		valid_time_window TEXT,
		terms_and_conditions TEXT,
		discount_type TEXT,
		discount_value REAL,
		max_usage_per_user INTEGER
	);
	`
	_, err = DB.Exec(query)
	if err != nil {
		return fmt.Errorf("could not create table: %v", err)
	}

	log.Println("Database initialized and table created")
	return nil
}

// CloseDB closes the database connection
func CloseDB() {
	err := DB.Close()
	if err != nil {
		log.Fatalf("could not close the database: %v", err)
	}
}
