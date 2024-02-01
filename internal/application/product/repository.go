package product

type ProductRepository interface {
	Save(product *Product) error
	FindAll() ([]Product, error)
}
