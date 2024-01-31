package database

import (
	"product-catalog-manager/internal/product_catalog"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	product_catalog.ProductRepository
	mock.Mock
}

func (m *ProductRepositoryMock) Save(product *product_catalog.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *ProductRepositoryMock) FindAll() ([]product_catalog.Product, error) {
	args := m.Called()
	return args.Get(0).([]product_catalog.Product), args.Error(1)
}
