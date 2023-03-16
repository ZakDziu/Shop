package application

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"shop/internal/api"
	"shop/internal/core/services/auth"
	"shop/internal/core/services/order"
	"shop/internal/core/services/product"
	"shop/internal/core/services/user"
	"shop/internal/repositories"
)

func Start() {
	initEnv()
	db = NewPostgreSQLDBConnection()

	userRepo := repositories.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	authRepo := repositories.NewAuthRepository(db)
	authService := auth.NewAuthService(authRepo, userRepo)
	productRepo := repositories.NewProductRepository(db)
	productService := product.NewProductService(productRepo)
	orderRepo := repositories.NewOrderRepository(db)
	orderService := order.NewOrderService(orderRepo, productRepo)

	mux := api.NewMux(userService, authService, productService, orderService)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}

func initEnv() {
	if _, err := os.Stat(".env"); err == nil {
		var fileEnv map[string]string
		fileEnv, _ = godotenv.Read()

		for key, val := range fileEnv {
			if len(os.Getenv(key)) == 0 {
				_ = os.Setenv(key, val)
			}
		}
	}
}
