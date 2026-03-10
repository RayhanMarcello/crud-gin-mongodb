package repository

import (
	"context"
	"crud-gin-mongodb/models"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *models.Product) error
	FindAllProduct(ctx context.Context) ([]models.Product, error)
	FindById(ctx context.Context, id string) (*models.Product, error)
	UpdateByID(ctx context.Context, id string, product *models.Product) (*models.Product, error)
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

func (r *productRepository) FindAllProduct(ctx context.Context) ([]models.Product, error) {
	var products []models.Product

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil

}

func (r *productRepository) FindById(ctx context.Context, id string) (*models.Product, error) {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var product models.Product
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) UpdateByID(ctx context.Context, id string, product *models.Product) (*models.Product, error) {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id!")
	}

	update := bson.M{
		"$set": bson.M{
			"name":        product.Name,
			"description": product.Description,
			"price":       product.Price,
			"stock":       product.Stock,
			"updated_at":  product.UpdatedAt,
		},
	}
	result, err := r.collection.UpdateOne(ctx, bson.M{
		"_id": objID,
	}, update)

	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, errors.New("product not found")
	}
	return product, nil

}
