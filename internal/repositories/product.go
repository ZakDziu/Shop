package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"shop/internal/core/domain"
	"shop/internal/core/models"
	"shop/internal/core/ports"
)

type ProductRepository struct {
	db *sqlx.DB
}

var _ ports.ProductRepository = &ProductRepository{}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(ctx context.Context, product models.Product) (uint, error) {
	var id uint
	sqlRequest := `INSERT INTO products ( 
                   		user_id, 
                   		name,
                      	description,
                      	price
                   	) 
					VALUES (
        				:user_id, 
        				:name, 
        				:description, 
        				:price
        			) RETURNING product_id`

	stmt, err := r.db.PrepareNamed(sqlRequest)
	if err != nil {
		return 0, err
	}

	if err = stmt.GetContext(ctx, &id, product); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ProductRepository) Update(ctx context.Context, product models.Product) error {
	sqlRequest := `UPDATE products
					SET name = $1,
					    description = $2,
					    price = $3
					WHERE product_id = $4`
	if product.UserId != 0 {
		sqlRequest += ` AND user_id = $5`
		_, err := r.db.ExecContext(ctx, sqlRequest, product.Name, product.Description, product.Price, product.ID, product.UserId)
		return err
	}
	_, err := r.db.ExecContext(ctx, sqlRequest, product.Name, product.Description, product.Price, product.ID)
	return err

}

func (r *ProductRepository) Delete(ctx context.Context, product models.Product) error {
	sqlRequest := `DELETE FROM products WHERE product_id = $1`
	if product.UserId != 0 {
		sqlRequest += ` AND user_id = $2`
		_, err := r.db.ExecContext(ctx, sqlRequest, product.ID, product.UserId)
		return err
	}
	_, err := r.db.ExecContext(ctx, sqlRequest, product.ID)
	return err

}

func (r *ProductRepository) Get(ctx context.Context, id uint) (domain.Product, error) {
	var product domain.Product
	sqlRequest := `SELECT product_id, name, description, price FROM products WHERE product_id = $1`

	err := r.db.QueryRowxContext(
		ctx,
		sqlRequest,
		id,
	).StructScan(&product)

	return product, err
}

func (r *ProductRepository) GetAllByUser(ctx context.Context, id uint) ([]domain.Product, error) {
	result := make([]domain.Product, 0)

	sqlRequest := `SELECT product_id, name, description, price FROM products WHERE user_id = $1`
	err := r.db.SelectContext(ctx,
		&result,
		sqlRequest,
		id,
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *ProductRepository) GetAll(ctx context.Context) ([]domain.Product, error) {
	result := make([]domain.Product, 0)

	sqlRequest := `SELECT product_id, name, description, price FROM products`
	err := r.db.SelectContext(ctx,
		&result,
		sqlRequest,
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}
