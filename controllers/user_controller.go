package controllers

import (
	"net/http"

	"gin-boilerplate/models"
	"gin-boilerplate/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

// The userService is passed to the UserController constructor to establish the dependency between the controller and the service.
// This dependency injection is good for testing and also for reusability.
func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := uc.userService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.userService.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (uc *UserController) GetUser(c *gin.Context) {
	userID := c.Param("id")

	user, err := uc.userService.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	c.JSON(http.StatusOK, user)
}
