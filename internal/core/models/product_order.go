package models

type ProductOrder struct {
	OrderId   uint `db:"order_id"`
	ProductId uint `db:"product_id"`
}
