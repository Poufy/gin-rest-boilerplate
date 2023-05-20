package controllers

import (
	"net/http"

	"gin-boilerplate/models"
	"gin-boilerplate/services"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService *services.ProductService
}

// The productService is passed to the ProductController constructor to establish the dependency between the controller and the service.
// This dependency injection is good for testing and also for reusability.
func NewProductController(productService *services.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (uc *ProductController) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdProduct, err := uc.productService.CreateProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}

func (uc *ProductController) GetProducts(c *gin.Context) {
	products, err := uc.productService.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (uc *ProductController) GetProduct(c *gin.Context) {
	productID := c.Param("id")

	product, err := uc.productService.GetProduct(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product"})
		return
	}

	c.JSON(http.StatusOK, product)
}
