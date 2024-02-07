package dependency_provider

import (
	"context"

	"product-catalog-manager/internal/application/configuration"
	"product-catalog-manager/internal/application/product"
	"product-catalog-manager/internal/infra/mongodb_adapter"
	"product-catalog-manager/pkg/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
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

func (d *DependencyProvider) GetMongoClient() (*mongo.Client, context.Context) {
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
	return client, context
}

func (d *DependencyProvider) GetProductRepository() product.ProductRepository {
	client, context := d.GetMongoClient()
	collection := client.Database(d.GetConfig().DBDatabase).Collection(mongodb_adapter.ProductRepositoryCollection)
	return mongodb_adapter.NewProductRepository(context, collection)
}

// func (d *DependencyProvider) GetKafkaAdapter() message_broker.MessageBroker {
// 	return kafka_adapter.NewKafkaAdapter(d)
// }
