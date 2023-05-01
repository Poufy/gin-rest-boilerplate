package routes

import (
	"gin-boilerplate/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.AuthMiddleware())

	api := router.Group("/api")
	{
		UserRoutes(api)
		ProductRoutes(api)
	}

	return router
}
