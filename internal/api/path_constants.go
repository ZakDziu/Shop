package api

const (
	CreateUser  = "/api/create-user"
	UpdateUser  = "/api/update-user"
	DeleteUser  = "/api/delete-user"
	GetAllUsers = "/api/get-all-users"

	CreateProduct        = "/api/create-product"
	UpdateProduct        = "/api/update-product"
	DeleteProduct        = "/api/delete-product"
	GetAllProductsByUser = "/api/get-all-products-by-user"
	GetAllProducts       = "/api/get-all-products"

	CreateOrder       = "/api/create-order"
	GetOrdersByUserId = "/api/get-orders-by-user-id"
	DeleteOrder       = "/api/delete-order"
	UpdateOrder       = "/api/update-order"
	GetAllOrders      = "/api/get-all-orders"
)
