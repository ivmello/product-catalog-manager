package dependency_provider

import (
	"product-catalog-manager/internal/configuration"
	"product-catalog-manager/internal/infra/database"
	"product-catalog-manager/internal/product_catalog"
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

func (d *DependencyProvider) GetProductService() product_catalog.ProductService {
	return product_catalog.NewProductService(d.GetProductRepository())
}

func (d *DependencyProvider) GetProductRepository() product_catalog.ProductRepository {
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
