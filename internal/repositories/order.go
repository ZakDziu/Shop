package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"shop/internal/core/models"
	"shop/internal/core/ports"
)

type OrderRepository struct {
	db *sqlx.DB
}

var _ ports.OrderRepository = &OrderRepository{}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(ctx context.Context, order models.Order, products []uint) (uint, error) {
	var id uint
	sqlRequestOrder := `INSERT INTO orders ( 
                   		user_id
                   	) 
					VALUES (
        				$1
        			) RETURNING order_id`
	sqlRequestProducts := `INSERT INTO product_order ( 
                           product_id, 
                           order_id
                   	) 
					VALUES (
        				$1, $2
        			)`
	err := r.db.QueryRowContext(ctx, sqlRequestOrder, order.UserId).Scan(&id)
	if err != nil {
		return 0, err
	}

	for _, p := range products {
		err = r.db.QueryRowContext(ctx, sqlRequestProducts, p, id).Err()
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (r *OrderRepository) Get(ctx context.Context, id uint) ([]models.ProductOrder, error) {
	result := make([]models.ProductOrder, 0)

	sqlRequest := `SELECT product_id, order_id FROM product_order WHERE order_id = $1`
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

func (r *OrderRepository) Delete(ctx context.Context, order models.Order) error {
	sqlRequest := `DELETE FROM orders WHERE order_id = $1`
	if order.UserId != 0 {
		sqlRequest += ` AND user_id = $2`
		_, err := r.db.ExecContext(ctx, sqlRequest, order.OrderId, order.UserId)
		return err
	}
	_, err := r.db.ExecContext(ctx, sqlRequest, order.OrderId)
	return err
}

func (r *OrderRepository) Update(ctx context.Context, order models.Order, products []uint) (uint, error) {
	order, err := r.GetOrderUserId(ctx, order.OrderId)
	if err != nil {
		return 0, err
	}
	err = r.Delete(ctx, order)
	if err != nil {
		return 0, err
	}

	return r.Create(ctx, order, products)
}

func (r *OrderRepository) GetOrdersByUserId(ctx context.Context, userId uint) ([]models.Order, error) {
	result := make([]models.Order, 0)

	sqlRequest := `SELECT order_id, user_id FROM orders WHERE user_id = $1`
	err := r.db.SelectContext(ctx,
		&result,
		sqlRequest,
		userId,
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *OrderRepository) GetOrderUserId(ctx context.Context, orderId uint) (models.Order, error) {
	var order models.Order

	sqlRequest := `SELECT order_id, user_id FROM orders WHERE order_id = $1`
	err := r.db.QueryRowxContext(
		ctx,
		sqlRequest,
		orderId,
	).StructScan(&order)
	if err != nil {
		return models.Order{}, err
	}

	return order, nil
}

func (r *OrderRepository) GetAll(ctx context.Context) ([]models.Order, error) {
	result := make([]models.Order, 0)

	sqlRequest := `SELECT order_id, user_id FROM orders`
	err := r.db.SelectContext(ctx,
		&result,
		sqlRequest,
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}
