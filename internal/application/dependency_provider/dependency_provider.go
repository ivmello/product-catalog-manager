package dependency_provider

import (
	"product-catalog-manager/internal/application/configuration"
	"product-catalog-manager/internal/application/product"
	"product-catalog-manager/internal/infra/database"
	"product-catalog-manager/pkg/mongodb"
)

type DependencyProvider struct {
	config *configuration.Config
}

func New(config *configuration.Config) *DependencyProvider {
	return &DependencyProvider{config: config}
}

func (d *DependencyProvider) GetConfig() *configuration.Config {
	return d.config
}

func (d *DependencyProvider) GetProductService() product.ProductService {
	return product.NewProductService(d.GetProductRepository())
}

func (d *DependencyProvider) GetProductRepository() product.ProductRepository {
	client, context, err := mongodb.Connect(mongodb.MongoDBConfig{
		User:     d.GetConfig().DBUser,
		Password: d.GetConfig().DBPassword,
		Host:     d.GetConfig().DBHost,
		Port:     d.GetConfig().DBPort,
		Database: d.GetConfig().DBDatabase,
	})
	if err != nil {
		panic(err)
	}
	err = mongodb.Ping(context, client)
	if err != nil {
		panic(err)
	}
	collection := client.Database(d.GetConfig().DBDatabase).Collection(database.ProductRepositoryCollection)
	return database.NewProductRepository(context, collection)
}
