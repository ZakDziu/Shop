package auth

import (
	"context"
	"errors"
	"os"
	"shop/internal/core/domain"
	"shop/internal/core/ports"
)

type AuthService struct {
	authRepo ports.AuthRepository
	userRepo ports.UserRepository
}

var _ ports.AuthService = &AuthService{}

func NewAuthService(authRepo ports.AuthRepository, userRepo ports.UserRepository) *AuthService {
	return &AuthService{authRepo: authRepo, userRepo: userRepo}
}

func (s *AuthService) CheckUserCredentialsAndRole(ctx context.Context, username, password string, expectedRole domain.Role) (uint, error) {
	if username == os.Getenv("SUPER_USERNAME") && password == os.Getenv("SUPER_PASSWORD") {
		return 0, nil
	}
	pass, err := s.authRepo.GetUserPassword(ctx, username)
	if pass != password || err != nil {
		return 0, errors.New("can't find user")
	}
	role, err := s.authRepo.GetUserRole(ctx, username)
	if role != expectedRole {
		return 0, errors.New("user can't use this handle")
	}
	id, err := s.userRepo.GetIDByName(ctx, username)
	if role != expectedRole {
		return 0, errors.New("error in get id")
	}
	return id, nil
}
