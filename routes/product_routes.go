package routes

import (
	"gin-boilerplate/controllers"

	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(router *gin.RouterGroup, productController *controllers.ProductController) {
	productRoutes := router.Group("/products")
	{
		productRoutes.POST("", productController.CreateProduct)
		productRoutes.GET("", productController.GetProducts)
		productRoutes.GET("/:id", productController.GetProduct)
	}
}
