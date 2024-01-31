package product_catalog

type ProductRepository interface {
	Save(product *Product) error
	FindAll() ([]Product, error)
}
