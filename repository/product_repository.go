package repository

import (
	"context"
	"crud-gin-mongodb/models"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *models.Product) error
	FindAllProduct(ctx context.Context) ([]models.Product, error)
}

type productRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(collection *mongo.Collection) ProductRepository {
	return &productRepository{
		collection: collection,
	}
}

func (p *productRepository) CreateProduct(ctx context.Context, product *models.Product) error {
	now := time.Now()
	product.CreatedAt := now
}
