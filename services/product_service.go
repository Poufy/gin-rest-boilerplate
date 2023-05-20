package services

import (
	"gin-boilerplate/models"
	"gin-boilerplate/repositories"
)

type ProductService struct {
	productRepository *repositories.ProductRepository
}

func NewProductService(productRepository *repositories.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (us *ProductService) CreateProduct(product *models.Product) (*models.Product, error) {
	// Perform any necessary validations or transformations on the product data
	// ...

	// Call the ProductRepository to create the product
	createdProduct, err := us.productRepository.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return createdProduct, nil
}

func (us *ProductService) GetProducts() ([]models.Product, error) {
	// Call the ProductRepository to fetch products
	products, err := us.productRepository.GetProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (us *ProductService) GetProduct(productID string) (*models.Product, error) {
	// Call the ProductRepository to fetch a specific product
	product, err := us.productRepository.GetProduct(productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}
