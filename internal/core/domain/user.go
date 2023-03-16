package domain

type Role string

const (
	Seller Role = "Seller"
	Buyer  Role = "Buyer"
	Admin  Role = "Admin"
)

type User struct {
	ID       uint   `json:"id" db:"id"`
	Name     string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Phone    string `json:"phone" db:"phone"`
	Role     Role   `json:"role" db:"role"`
}
