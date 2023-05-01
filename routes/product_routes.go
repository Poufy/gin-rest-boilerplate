package routes

import (
	"gin-boilerplate/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.RouterGroup) {
	r.GET("/products", controllers.GetProducts)
}
