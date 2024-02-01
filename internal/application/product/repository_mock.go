package product

import (
	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	ProductRepository
	mock.Mock
}

func (m *ProductRepositoryMock) Save(product *Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *ProductRepositoryMock) FindAll() ([]Product, error) {
	args := m.Called()
	return args.Get(0).([]Product), args.Error(1)
}
