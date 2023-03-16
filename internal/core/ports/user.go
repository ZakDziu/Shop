package ports

import (
	"context"
	"shop/internal/core/domain"
)

type UserService interface {
	Create(ctx context.Context, user domain.User) (int, error)
	Update(ctx context.Context, user domain.User) error
	GetIDByName(ctx context.Context, name string) (uint, error)
	Delete(ctx context.Context, id uint) error
	GetAll(ctx context.Context) ([]domain.User, error)
}

type UserRepository interface {
	Create(ctx context.Context, user domain.User) (id int, err error)
	Update(ctx context.Context, user domain.User) error
	GetIDByName(ctx context.Context, name string) (uint, error)
	Delete(ctx context.Context, id uint) error
	GetAll(ctx context.Context) ([]domain.User, error)
}
