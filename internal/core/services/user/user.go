package user

import (
	"context"
	"errors"
	"shop/internal/core/domain"

	"shop/internal/core/ports"
)

type UserService struct {
	userRepo ports.UserRepository
}

var _ ports.UserService = &UserService{}

func NewUserService(userRepo ports.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Create(ctx context.Context, user domain.User) (int, error) {
	id, err := s.userRepo.GetIDByName(ctx, user.Name)
	if id != 0 || err != nil {
		return 0, errors.New("user with with username exist")
	}
	return s.userRepo.Create(ctx, user)
}

func (s *UserService) Update(ctx context.Context, user domain.User) error {
	return s.userRepo.Update(ctx, user)
}

func (s *UserService) GetIDByName(ctx context.Context, name string) (uint, error) {
	return s.userRepo.GetIDByName(ctx, name)
}

func (s *UserService) Delete(ctx context.Context, id uint) error {
	return s.userRepo.Delete(ctx, id)
}

func (s *UserService) GetAll(ctx context.Context) ([]domain.User, error) {
	return s.userRepo.GetAll(ctx)
}
