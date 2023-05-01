package services

import "gin-boilerplate/models"

func GetProducts() []models.Product {
	return models.Products
}
