package handler

import (
	"context"
	"mm-pddikti-cms/internal/adapter"
	"mm-pddikti-cms/internal/module/auth/ports"
	"mm-pddikti-cms/internal/module/auth/entity"
	"mm-pddikti-cms/pkg/response"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService ports.AuthService
}

func NewAuthHandler(authService ports.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	loginRequest := &entity.LoginRequest{}
	if err := ctx.BodyParser(loginRequest); err != nil {
		return response.SendResponse(ctx, response.ResponseParams{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Invalid Input",
		})
	}

	validationErrors := adapter.Adapters.Validator.Validate(loginRequest)
	if len(validationErrors) > 0 {
		return response.SendResponse(ctx, response.ResponseParams{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Validation Failed",
			Errors:     validationErrors,
		})
	}

	accessToken, refreshToken, err := h.authService.Login(context, loginRequest.EmailOrUsername, loginRequest.Password)
	if err != nil {
		return response.SendResponse(ctx, response.ResponseParams{
			StatusCode: fiber.StatusUnauthorized,
			Message:    err.Error(),
		})
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(time.Hour * time.Duration(72)),
		Secure:   true,
		HTTPOnly: true,
		SameSite: "Lax",
	})

	ctx.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(time.Hour * time.Duration(120)),
		Secure:   true,
		HTTPOnly: true,
		SameSite: "Lax",
	})

	return response.SendResponse(ctx, response.ResponseParams{
		StatusCode: fiber.StatusOK,
		Message:    "Login Success!",
	})
}
