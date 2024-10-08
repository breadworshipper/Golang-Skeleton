package jwthandler

import (
	"mm-pddikti-cms/internal/module/user/entity"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type CustomClaims struct {
	UserId   uuid.UUID   `json:"id"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Role     entity.Role `json:"role"`
	jwt.RegisteredClaims
}

type CostumClaimsPayload struct {
	UserId                 uuid.UUID   `json:"id"`
	Username               string      `json:"username"`
	Email                  string      `json:"email"`
	Role                   entity.Role `json:"role"`
	AccessTokenExpiration  time.Time   `json:"access_token_expiration"`
	RefreshTokenExpiration time.Time   `json:"refresh_token_expiration"`
}
