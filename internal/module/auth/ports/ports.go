package ports

import (
	"context"
	"mm-pddikti-cms/internal/module/user/entity"
)

type AuthRepository interface {
	FindUserByUsernameOrEmail(ctx context.Context, username_or_email string) (*entity.User, error)
}

type AuthService interface {
	Login(ctx context.Context, username, password string) (string, string, error)
}
