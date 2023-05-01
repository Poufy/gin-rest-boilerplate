package db

import (
	"fmt"
	"gin-boilerplate/config"
	// Import your database driver here
)

func InitDB() {
	// Replace this function with your actual database initialization logic
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Cfg.DB.Host,
		config.Cfg.DB.Port,
		config.Cfg.DB.User,
		config.Cfg.DB.Password,
		config.Cfg.DB.Name,
	)

	fmt.Println("Connecting to database with connection string:", connectionString)
	// Connect to your database and perform any necessary setup here
}
