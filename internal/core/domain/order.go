package domain

type Order struct {
	ID       uint      `json:"id"`
	Products []Product `json:"products"`
}
