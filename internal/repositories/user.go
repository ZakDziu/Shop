package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"shop/internal/core/domain"
	"shop/internal/core/ports"
)

type UserRepository struct {
	db *sqlx.DB
}

var _ ports.UserRepository = &UserRepository{}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user domain.User) (id int, err error) {
	sqlRequest := `INSERT INTO users ( 
                   		username, 
                   		password, 
                   		role, 
                   		phone
                   		) 
					VALUES (
        				:username, 
        				:password, 
						:role,
        				:phone
        			) RETURNING id`

	stmt, err := r.db.PrepareNamed(sqlRequest)
	if err != nil {
		return 0, err
	}
	if err = stmt.GetContext(ctx, &id, user); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UserRepository) Update(ctx context.Context, user domain.User) error {
	sqlRequest := `UPDATE users 
					SET username = $1,
					    password = $2,
					    phone = $3,
					    role = $4
					WHERE id = $5`
	_, err := r.db.ExecContext(ctx, sqlRequest, user.Name, user.Password, user.Phone, user.Role, user.ID)
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id uint) error {
	sqlRequest := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, sqlRequest, id)
	return err
}

func (r *UserRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	result := make([]domain.User, 0)

	sqlRequest := `SELECT id, username, password, role, phone FROM users`
	err := r.db.SelectContext(ctx,
		&result,
		sqlRequest,
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *UserRepository) GetIDByName(ctx context.Context, name string) (uint, error) {
	var id uint
	sqlRequest := `SELECT id FROM users WHERE username = $1`

	err := r.db.QueryRowxContext(
		ctx,
		sqlRequest,
		name,
	).Scan(&id)
	return id, err
}
