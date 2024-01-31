package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func Connect(config MongoDBConfig) (*mongo.Client, context.Context, error) {
	connectionContext := context.Background()
	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/?retryWrites=false",
		config.User,
		config.Password,
		config.Host,
		config.Port,
	)
	client, err := mongo.Connect(connectionContext, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	return client, connectionContext, nil
}

func Ping(connectionContext context.Context, client *mongo.Client) error {
	if err := client.Ping(connectionContext, readpref.Primary()); err != nil {
		return err
	}
	return nil
}
