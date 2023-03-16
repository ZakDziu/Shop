package ports

import (
	"context"
	"shop/internal/core/domain"
)

type AuthService interface {
	CheckUserCredentialsAndRole(ctx context.Context, username, password string, expectedRole domain.Role) (uint, error)
	HashPassword(password string) (string, error)
}

type AuthRepository interface {
	GetUserPassword(ctx context.Context, username string) (string, error)
	GetUserRole(ctx context.Context, username string) (domain.Role, error)
}
