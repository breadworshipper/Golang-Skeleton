package ports

import (
	"context"
	"mm-pddikti-cms/internal/module/user/entity"

	"github.com/google/uuid"
)

type UserRepository interface {
	FindUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
}

type UserService interface {
	Profile(ctx context.Context, id uuid.UUID) (*entity.User, error)
}
