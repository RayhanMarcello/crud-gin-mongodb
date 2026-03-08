package repository

import (
	"context"
	"crud-gin-mongodb/models"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *models.Product) error
	// FindAllProduct(ctx context.Context) ([]models.Product, error)
}

type productRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(coll *mongo.Collection) ProductRepository {
	return &productRepository{
		collection: coll,
	}
}

func (r *productRepository) CreateProduct(ctx context.Context, product *models.Product) error {
	now := time.Now()
	product.CreatedAt = now

	res, err := r.collection.InsertOne(ctx, product)
	if err != nil {
		return err
	}
	if oID, ok := res.InsertedID.(bson.ObjectID); ok {
		product.ID = oID
	}

	return nil
}
