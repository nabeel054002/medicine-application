package db

import (
	"fmt"
	"log"
)

// confirm if this is valid schema or not
func CreateSchema() error {
	tables := []string{
		`CREATE TABLE IF NOT EXISTS coupons (
			coupon_code TEXT PRIMARY KEY,
			usage_type TEXT CHECK(usage_type IN ('one_time', 'multi_use', 'time_based')),
			max_usage_per_user INT,
			min_order_value DECIMAL,
			terms_and_conditions TEXT,
			expiry_date TIMESTAMP,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`,

		`CREATE TABLE IF NOT EXISTS categories (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL UNIQUE
		);`,

		`CREATE TABLE IF NOT EXISTS medicines (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			category_id INTEGER,
			FOREIGN KEY (category_id) REFERENCES categories(id)
		);`,

		`CREATE TABLE IF NOT EXISTS coupon_applicable_categories (
			coupon_code TEXT,
			category_id INTEGER,
			PRIMARY KEY (coupon_code, category_id),
			FOREIGN KEY (coupon_code) REFERENCES coupons(coupon_code),
			FOREIGN KEY (category_id) REFERENCES categories(id)
		);`,

		`CREATE TABLE IF NOT EXISTS coupon_applicable_medicines (
			coupon_code TEXT,
			medicine_id TEXT,
			PRIMARY KEY (coupon_code, medicine_id),
			FOREIGN KEY (coupon_code) REFERENCES coupons(coupon_code),
			FOREIGN KEY (medicine_id) REFERENCES medicines(id)
		);`,

		`CREATE TABLE IF NOT EXISTS coupon_usages (
			coupon_code TEXT,
			user_id TEXT,
			usage_count INT DEFAULT 0,
			PRIMARY KEY (coupon_code, user_id),
			FOREIGN KEY (coupon_code) REFERENCES coupons(coupon_code)
		);`,

		// for multiple time windows - hence separate primary key
		`CREATE TABLE IF NOT EXISTS time_windows (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			coupon_code TEXT,
			start_time TIMESTAMP,
			end_time TIMESTAMP,
			FOREIGN KEY (coupon_code) REFERENCES coupons(coupon_code)
		);`,
		
		`CREATE TABLE IF NOT EXISTS discounts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			coupon_code TEXT,
			discount_type TEXT CHECK(discount_type IN ('inventory', 'charges')),
			discount_value DECIMAL,
			FOREIGN KEY (coupon_code) REFERENCES coupons(coupon_code)
		);`,
	}

	for _, table := range tables {
		_, err := DB.Exec(table)
		if err != nil {
			return fmt.Errorf("error creating schema: %w", err)
		}
	}

	log.Println("Tables created successfully.")
	return nil
}
