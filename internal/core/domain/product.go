package domain

import "github.com/shopspring/decimal"

type Product struct {
	ID          uint            `json:"id" db:"product_id"`
	Name        string          `json:"productName" db:"name"`
	Description string          `json:"description" db:"description"`
	Price       decimal.Decimal `json:"price" db:"price"`
}
