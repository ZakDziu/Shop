package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"shop/internal/core/domain"
	"shop/internal/core/services/auth"
	"shop/internal/core/services/order"
)

func GetAllOrders(
	authService *auth.AuthService,
	orderService *order.OrderService,

) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(&requestErrorFmt{Error: "no auth"})

			return
		}
		_, err := authService.CheckUserCredentialsAndRole(ctx, username, password, domain.Admin)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(&requestErrorFmt{Error: err.Error()})

			return
		}

		orders, err := orderService.GetAll(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(&requestErrorFmt{Error: err.Error()})

			return
		}

		err = json.NewEncoder(w).Encode(orders)
		if err != nil {
			log.Panic(err)
		}
	}
}
