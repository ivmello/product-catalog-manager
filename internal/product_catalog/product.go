package product_catalog

import "github.com/google/uuid"

type Product struct {
	ID          string
	Title       string
	Description string
	Price       float64
}

func NewProduct(title, description string, price float64) *Product {
	return &Product{
		ID:          uuid.New().String(),
		Title:       title,
		Description: description,
		Price:       price,
	}
}
