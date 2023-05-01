package services

import "gin-boilerplate/models"

func GetUsers() []models.User {
	return models.Users
}
