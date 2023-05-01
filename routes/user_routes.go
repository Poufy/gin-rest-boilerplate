package routes

import (
	"gin-boilerplate/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	r.GET("/users", controllers.GetUsers)
}
