package controllers

import (
	"gin-boilerplate/services"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := services.GetUsers()
	c.JSON(200, users)
}
