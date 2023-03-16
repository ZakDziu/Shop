package models

type Order struct {
	OrderId uint `db:"order_id"`
	UserId  uint `db:"user_id"`
}
