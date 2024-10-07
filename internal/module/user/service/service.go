package service

import (
	"pddikti-cms/internal/module/user/entity"
	"pddikti-cms/internal/module/user/ports"
)


var _ ports.UserService = &userService{}

type userService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *userService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Login(req *entity.LoginRequest) (*entity.LoginResponse, error) {
	return nil, nil
}

func (s *userService) Register(req *entity.RegisterRequest) error {
	return nil
}
