package ports

import (
	"context"
	"shop/internal/core/domain"
	"shop/internal/core/models"
)

type OrderService interface {
	Create(ctx context.Context, o domain.Order, userId uint) (uint, error)
	Get(ctx context.Context, id uint) (domain.Order, error)
	GetOrdersByUserId(ctx context.Context, userId uint) ([]domain.Order, error)
	Delete(ctx context.Context, id uint, userId uint) error
	Update(ctx context.Context, o domain.Order, userId uint) (uint, error)
	GetAll(ctx context.Context) ([]domain.Order, error)
}

type OrderRepository interface {
	Create(ctx context.Context, order models.Order, products []uint) (uint, error)
	Get(ctx context.Context, id uint) ([]models.ProductOrder, error)
	GetOrdersByUserId(ctx context.Context, userId uint) ([]models.Order, error)
	Delete(ctx context.Context, order models.Order) error
	Update(ctx context.Context, order models.Order, products []uint) (uint, error)
	GetOrderUserId(ctx context.Context, orderId uint) (models.Order, error)
	GetAll(ctx context.Context) ([]models.Order, error)
}
