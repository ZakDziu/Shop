package models

import "github.com/shopspring/decimal"

type Product struct {
	ID          uint            `db:"product_id"`
	Name        string          `db:"name"`
	UserId      uint            `db:"user_id"`
	Description string          `db:"description"`
	Price       decimal.Decimal `db:"price"`
}
