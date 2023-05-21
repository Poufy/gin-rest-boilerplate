package database

import (
	"database/sql"
	"log"
)

// Here we can define functions that will create the necessary tables and any initial schema setup.
func CreateTables(db *sql.DB) {
	// Create users table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			lastname VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL
		)
	`)
	if err != nil {
		log.Fatal("Failed to create users table:", err)
	}

	// Create products table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			price NUMERIC(10,2) NOT NULL
		)
	`)
	if err != nil {
		log.Fatal("Failed to create products table:", err)
	}

	// Add any additional table creation or schema setup here
	// ...
}

func RunMigrations(db *sql.DB) {
	// Call the functions for initial table creation or schema setup
	CreateTables(db)
	// ...
}
