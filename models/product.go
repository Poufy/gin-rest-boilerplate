package models

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var Products = []Product{
	{ID: 1, Name: "Product 1", Price: 100},
	{ID: 2, Name: "Product 2", Price: 200},
}
