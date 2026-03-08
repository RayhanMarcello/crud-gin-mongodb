package service

import (
	"context"
	"crud-gin-mongodb/dto"
	"crud-gin-mongodb/models"
	"crud-gin-mongodb/repository"
	"time"
)

type ProductService interface {
	CreateService(ctx context.Context, req dto.CreateProdReq) (*models.Product, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{
		repo: repo,
	}
}

func (s *productService) CreateService(ctx context.Context, req dto.CreateProdReq) (*models.Product, error) {
	product := &models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	if err := s.repo.CreateProduct(ctx, product); err != nil {
		return nil, err
	}
	return product, nil
}
