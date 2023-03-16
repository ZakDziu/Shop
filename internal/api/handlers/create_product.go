package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"shop/internal/core/domain"
	"shop/internal/core/services/auth"
	"shop/internal/core/services/product"
)

func CreateProduct(
	authService *auth.AuthService,
	productService *product.ProductService,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(&requestErrorFmt{Error: "no auth"})

			return
		}
		userId, err := authService.CheckUserCredentialsAndRole(ctx, username, password, domain.Seller)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(&requestErrorFmt{Error: err.Error()})

			return
		}
		req := &domain.Product{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(&requestErrorFmt{Error: err.Error()})

			return
		}
		id, err := productService.Create(ctx, *req, userId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(&requestErrorFmt{Error: err.Error()})

			return
		}
		pr, err := productService.Get(ctx, id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(&requestErrorFmt{Error: err.Error()})

			return
		}

		err = json.NewEncoder(w).Encode(pr)
		if err != nil {
			log.Panic(err)
		}
	}
}
