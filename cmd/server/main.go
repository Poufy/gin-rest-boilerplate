package main

import (
	"gin-boilerplate/config"
	"gin-boilerplate/controllers"
	"gin-boilerplate/database"
	"gin-boilerplate/repositories"
	"gin-boilerplate/routes"
	"gin-boilerplate/services"
	"log"
)

func main() {
	config.LoadConfig()
	db, err := database.NewDB()
	// Connect to the database
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	// Run database migrations
	database.RunMigrations(db)

	// Initialize repositories
	userRepository := repositories.NewUserRepository(db)
	productRepository := repositories.NewProductRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepository)
	productService := services.NewProductService(productRepository)

	// Initialize controllers
	userController := controllers.NewUserController(userService)
	productController := controllers.NewProductController(productService)

	router := routes.SetupRouter(userController, productController)
	router.Run(config.Cfg.SERVER.Port)
}
