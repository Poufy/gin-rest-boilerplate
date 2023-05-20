package routes

import (
	"gin-boilerplate/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.RouterGroup, userController *controllers.UserController) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("", userController.CreateUser)
		userRoutes.GET("", userController.GetUsers)
		userRoutes.GET("/:id", userController.GetUser)
	}
}
