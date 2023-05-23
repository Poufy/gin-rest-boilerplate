package controllers

import (
	"log"
	"net/http"
	"time"

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
		log.Default().Println(err)
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
		log.Default().Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) Login(c *gin.Context) {
	// Parse the request body into the login credentials struct
	var loginCredentials models.LoginCredentials
	if err := c.ShouldBindJSON(&loginCredentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the UserService to authenticate the user and generate an authentication token
	token, err := uc.userService.Login(loginCredentials.Email, loginCredentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Set the token expiration time in the response headers
	expirationTime := time.Now().Add(15 * time.Minute) // Token expires in 15 minutes
	c.Header("Authorization", token)
	c.Header("Expires-At", expirationTime.Format(time.RFC3339))

	c.JSON(http.StatusOK, gin.H{"token": token})
}
