package db

import (
	"database/sql"
	"fmt"
	"gin-boilerplate/config"
	"log"

	_ "github.com/lib/pq"
)

func NewDB() (*sql.DB, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Cfg.DB.User,
		config.Cfg.DB.Password,
		config.Cfg.DB.Host,
		config.Cfg.DB.Port,
		config.Cfg.DB.Name,
	)

	fmt.Println("Connecting to database with connection string:", connectionString)
	// Connect to your database and perform any necessary setup here
	db, err := sql.Open("postgres", "postgres://user:password@localhost:5432/mydatabase?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
		return nil, err
	}

	return db, nil
}
