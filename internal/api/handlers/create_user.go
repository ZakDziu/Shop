package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"shop/internal/core/domain"
	"shop/internal/core/services/auth"
	"shop/internal/core/services/user"
)

type CreateUserResponse struct {
	ID int `json:"id"`
}

func CreateUser(
	authService *auth.AuthService,
	userService *user.UserService,
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
		req := &domain.User{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(&requestErrorFmt{Error: err.Error()})

			return
		}
		id, err := userService.Create(ctx, *req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(&requestErrorFmt{Error: err.Error()})

			return
		}
		response := CreateUserResponse{ID: id}
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Panic(err)
		}
	}
}
