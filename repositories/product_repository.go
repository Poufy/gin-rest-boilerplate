package repositories

import (
	"database/sql"

	"gin-boilerplate/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (ur *ProductRepository) CreateProduct(product *models.Product) (*models.Product, error) {
	result, err := ur.db.Exec("INSERT INTO products (name) VALUES ($1)", product.Name)

	if err != nil {
		return nil, err
	}
	// Get the ID of the newly inserted product
	productID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	product.ID = int(productID)

	return product, nil
}

func (ur *ProductRepository) GetProducts() ([]models.Product, error) {
	rows, err := ur.db.Query("SELECT id, name FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (ur *ProductRepository) GetProduct(productID string) (*models.Product, error) {
	var product models.Product
	err := ur.db.QueryRow("SELECT id, name FROM products WHERE id = $1", productID).Scan(&product.ID, &product.Name)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
