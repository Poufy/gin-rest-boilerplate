package routes

import (
	"gin-boilerplate/controllers"
	"gin-boilerplate/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controllers.UserController, productController *controllers.ProductController) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.AuthMiddleware())

	api := router.Group("/api")
	{
		SetupUserRoutes(api, userController)
		SetupProductRoutes(api, productController)
	}

	return router
}
