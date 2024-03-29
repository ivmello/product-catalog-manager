package product

import "fmt"

type ProductService interface {
	CreateProduct(input CreateProductInput) (*CreateProductOutput, error)
	ListProducts() ([]ListProductsOutput, error)
	HandleMessage(msg []byte) error
	SendMessage(msg []byte) error
}

type service struct {
	productRepository ProductRepository
}

func NewProductService(productRepository ProductRepository) ProductService {
	return &service{
		productRepository: productRepository,
	}
}

func (s *service) CreateProduct(input CreateProductInput) (*CreateProductOutput, error) {
	product := NewProduct(input.Title, input.Description, input.Price)
	err := s.productRepository.Save(product)
	if err != nil {
		return nil, err
	}
	output := &CreateProductOutput{
		ID: product.ID,
	}
	return output, nil
}

func (s *service) ListProducts() ([]ListProductsOutput, error) {
	products, err := s.productRepository.FindAll()
	if err != nil {
		return nil, err
	}
	output := make([]ListProductsOutput, 0, len(products))
	for _, product := range products {
		output = append(output, ListProductsOutput(product))
	}
	return output, nil
}

func (s *service) HandleMessage(msg []byte) error {
	fmt.Println("Received message: ", string(msg))
	return nil
}

func (s *service) SendMessage(msg []byte) error {
	// implementar
	return nil
}
