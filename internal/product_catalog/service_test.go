package product_catalog_test

import (
	"errors"
	"testing"

	"product-catalog-manager/internal/infra/database"
	"product-catalog-manager/internal/product_catalog"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductService(t *testing.T) {
	t.Run("CreateProduct", func(t *testing.T) {
		t.Run("should create product successfully", func(t *testing.T) {
			input := product_catalog.CreateProductInput{
				Title:       "Product 1",
				Description: "Description 1",
				Price:       10.0,
			}
			productRepository := new(database.ProductRepositoryMock)
			productRepository.On("Save", mock.Anything).Return(nil)
			productService := product_catalog.NewProductService(productRepository)
			_, err := productService.CreateProduct(input)
			assert.Nil(t, err)
		})

		t.Run("should return error when product repository fails to save", func(t *testing.T) {
			input := product_catalog.CreateProductInput{
				Title:       "Product 1",
				Description: "Description 1",
				Price:       10.0,
			}
			productRepository := new(database.ProductRepositoryMock)
			productRepository.On("Save", mock.Anything).Return(errors.New("error on save product"))
			productService := product_catalog.NewProductService(productRepository)
			_, err := productService.CreateProduct(input)
			assert.Equal(t, "error on save product", err.Error())
		})
	})

	t.Run("ListProducts", func(t *testing.T) {
		t.Run("should list a list of products", func(t *testing.T) {
			productRepository := new(database.ProductRepositoryMock)
			productRepository.On("FindAll").Return([]product_catalog.Product{
				{ID: "1", Title: "Product 1", Description: "Description 1", Price: 10.0},
				{ID: "2", Title: "Product 2", Description: "Description 2", Price: 20.0},
			}, nil)
			productService := product_catalog.NewProductService(productRepository)
			products, err := productService.ListProducts()
			assert.Nil(t, err)
			assert.Equal(t, 2, len(products))
			assert.Equal(t, "1", products[0].ID)
			assert.Equal(t, "Product 1", products[0].Title)
			assert.Equal(t, "Description 1", products[0].Description)
			assert.Equal(t, 10.0, products[0].Price)
			assert.Equal(t, "2", products[1].ID)
			assert.Equal(t, "Product 2", products[1].Title)
			assert.Equal(t, "Description 2", products[1].Description)
			assert.Equal(t, 20.0, products[1].Price)
		})
	})
}
