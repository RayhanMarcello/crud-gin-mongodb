package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoConfig struct {
	Client      *mongo.Client
	Database    *mongo.Database
	ProductColl *mongo.Collection
}

func NewMongoConfig() (*MongoConfig, error) {
	uri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB")
	productCollectionName := os.Getenv("MONGO_PRODUCTS_COLLECTION")

	_, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("cant connect db err : %w", err)
	}
	db := client.Database(dbName)

	return &MongoConfig{
		Client:      client,
		Database:    db,
		ProductColl: db.Collection(productCollectionName),
	}, nil

}
