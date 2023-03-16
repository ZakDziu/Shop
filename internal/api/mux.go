package api

import (
	"net/http"
	"shop/internal/api/handlers"
	"shop/internal/core/services/auth"
	"shop/internal/core/services/order"
	"shop/internal/core/services/product"
	"shop/internal/core/services/user"
)

func NewMux(
	userService *user.UserService,
	authService *auth.AuthService,
	productService *product.ProductService,
	orderService *order.OrderService,
) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc(CreateUser, handlers.CreateUser(authService, userService))
	mux.HandleFunc(UpdateUser, handlers.UpdateUser(authService, userService))
	mux.HandleFunc(DeleteUser, handlers.DeleteUser(authService, userService))
	mux.HandleFunc(GetAllUsers, handlers.GetAllUsers(authService, userService))

	mux.HandleFunc(CreateProduct, handlers.CreateProduct(authService, productService))
	mux.HandleFunc(UpdateProduct, handlers.UpdateProduct(authService, productService))
	mux.HandleFunc(DeleteProduct, handlers.DeleteProduct(authService, productService))
	mux.HandleFunc(GetAllProductsByUser, handlers.GetAllProductsByUser(authService, productService))
	mux.HandleFunc(GetAllProducts, handlers.GetAllProducts(productService))

	mux.HandleFunc(CreateOrder, handlers.CreateOrder(authService, orderService))
	mux.HandleFunc(GetOrdersByUserId, handlers.GetOrdersByUserId(authService, orderService))
	mux.HandleFunc(DeleteOrder, handlers.DeleteOrder(authService, orderService))
	mux.HandleFunc(UpdateOrder, handlers.UpdateOrder(authService, orderService))
	mux.HandleFunc(GetAllOrders, handlers.GetAllOrders(authService, orderService))

	return mux
}
