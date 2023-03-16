package ports

import (
	"context"
	"shop/internal/core/domain"
	"shop/internal/core/models"
)

type ProductService interface {
	Create(ctx context.Context, product domain.Product, userId uint) (uint, error)
	Get(ctx context.Context, productId uint) (domain.Product, error)
	Update(ctx context.Context, product domain.Product, userId uint) error
	Delete(ctx context.Context, product domain.Product, userId uint) error
	GetAllByUser(ctx context.Context, userId uint) ([]domain.Product, error)
	GetAll(ctx context.Context) ([]domain.Product, error)
}

type ProductRepository interface {
	Create(ctx context.Context, product models.Product) (uint, error)
	Get(ctx context.Context, id uint) (domain.Product, error)
	Update(ctx context.Context, product models.Product) error
	Delete(ctx context.Context, product models.Product) error
	GetAllByUser(ctx context.Context, id uint) ([]domain.Product, error)
	GetAll(ctx context.Context) ([]domain.Product, error)
}
