package product_catalog

type CreateProductInput struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type CreateProductOutput struct {
	ID string `json:"id"`
}

type ListProductsOutput struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
