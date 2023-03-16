package product

import (
	"context"
	"shop/internal/core/domain"
	"shop/internal/core/models"
	"shop/internal/core/ports"
)

type ProductService struct {
	productRepo ports.ProductRepository
}

var _ ports.ProductService = &ProductService{}

func NewProductService(productRepo ports.ProductRepository) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (s *ProductService) Create(ctx context.Context, product domain.Product, userId uint) (uint, error) {
	return s.productRepo.Create(ctx, parseToModel(product, userId))
}

func (s *ProductService) Get(ctx context.Context, productId uint) (domain.Product, error) {
	return s.productRepo.Get(ctx, productId)
}

func (s *ProductService) Update(ctx context.Context, product domain.Product, userId uint) error {
	return s.productRepo.Update(ctx, parseToModel(product, userId))
}

func (s *ProductService) Delete(ctx context.Context, product domain.Product, userId uint) error {
	return s.productRepo.Delete(ctx, parseToModel(product, userId))
}

func (s *ProductService) GetAllByUser(ctx context.Context, userId uint) ([]domain.Product, error) {
	return s.productRepo.GetAllByUser(ctx, userId)
}

func (s *ProductService) GetAll(ctx context.Context) ([]domain.Product, error) {
	return s.productRepo.GetAll(ctx)
}

func parseToModel(p domain.Product, userId uint) models.Product {
	return models.Product{
		ID:          p.ID,
		Name:        p.Name,
		UserId:      userId,
		Description: p.Description,
		Price:       p.Price,
	}
}
