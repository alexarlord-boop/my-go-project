package data

import (
	"time"
)

// structure for an API product
type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float32   `json:"price"`
	SKU         string    `json:"sku"`
	CreatedOn   time.Time `json:"-"`
	UpdatedOn   time.Time `json:"-"`
	DeletedOn   time.Time `json:"-"`
}

func GetProducts() []*Product {
	return productList
}

// slice of pointers to Product structs
var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now(),
		UpdatedOn:   time.Now(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now(),
		UpdatedOn:   time.Now(),
	},
}
