package controllers

import (
	"gin-boilerplate/services"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	products := services.GetProducts()
	c.JSON(200, products)
}
