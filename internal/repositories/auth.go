package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"shop/internal/core/domain"
	"shop/internal/core/ports"
)

type AuthRepository struct {
	db *sqlx.DB
}

var _ ports.AuthRepository = &AuthRepository{}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) GetUserPassword(ctx context.Context, username string) (string, error) {
	var password string
	sqlRequest := `SELECT password FROM users WHERE username = $1`

	err := r.db.QueryRowxContext(
		ctx,
		sqlRequest,
		username,
	).Scan(&password)

	return password, err
}

func (r *AuthRepository) GetUserRole(ctx context.Context, username string) (domain.Role, error) {
	var role domain.Role
	sqlRequest := `SELECT role FROM users WHERE username = $1`

	err := r.db.QueryRowxContext(
		ctx,
		sqlRequest,
		username,
	).Scan(&role)

	return role, err
}
