package database

import (
	"context"
	"encoding/json"

	"product-catalog-manager/internal/product_catalog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ProductRepositoryCollection = "product"

type ProductRepository struct {
	collection *mongo.Collection
	context    context.Context
}

func NewProductRepository(
	context context.Context,
	collection *mongo.Collection,
) product_catalog.ProductRepository {
	return &ProductRepository{
		context:    context,
		collection: collection,
	}
}

func (r *ProductRepository) Save(product *product_catalog.Product) error {
	filter := bson.M{
		"identifier": product.ID,
	}
	update := bson.M{
		"$set": product,
	}
	upsert := true
	opts := options.Update().SetUpsert(upsert)
	_, err := r.collection.UpdateOne(r.context, filter, update, opts)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) FindAll() ([]product_catalog.Product, error) {
	var result bson.M
	cur, err := r.collection.Find(r.context, bson.D{{}}, options.Find())
	if err != nil {
		return nil, err
	}
	defer cur.Close(r.context)
	var products []product_catalog.Product
	for cur.Next(r.context) {
		err = cur.Decode(&result)
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}
		product, err := r.castBSONToProduct(result)
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
	}
	return products, err
}

func (r *ProductRepository) castBSONToProduct(result bson.M) (*product_catalog.Product, error) {
	bytes, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	product := new(product_catalog.Product)
	err = json.Unmarshal(bytes, &product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
