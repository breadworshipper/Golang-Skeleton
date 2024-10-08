package service

import (
	"context"
	"errors"
	"mm-pddikti-cms/internal/module/user/entity"
	"mm-pddikti-cms/internal/module/user/ports"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ ports.UserService = (*userService)(nil)

type userService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) ports.UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Profile(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	user, err := s.repo.FindUserByID(ctx, id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("User not found!")
		}
		return nil, err
	}

	return user, nil
}
