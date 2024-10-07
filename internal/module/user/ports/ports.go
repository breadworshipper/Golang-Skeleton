package ports

import (
	"pddikti-cms/internal/module/user/entity"
)

type UserRepository interface {
	EmailExist(email string) bool
	Register(user *entity.User) error
	GetUserByEmail(email string) (*entity.User, error)
	GetUserById(id string) (*entity.User, error)
}

type UserService interface {
	Login(req *entity.LoginRequest) (*entity.LoginResponse, error)
	Register(req *entity.RegisterRequest) error
}
