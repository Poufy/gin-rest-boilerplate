package routes

import (
	"gin-boilerplate/controllers"
	"gin-boilerplate/middleware"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.RouterGroup, userController *controllers.UserController) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("", userController.CreateUser)
		userRoutes.GET("", middleware.AuthMiddleware(), userController.GetUsers)
		userRoutes.GET("/:id", middleware.AuthMiddleware(), userController.GetUser)
		userRoutes.POST("/login", userController.Login)
	}
}
