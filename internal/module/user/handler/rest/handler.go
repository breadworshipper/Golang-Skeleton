package handler

import (
	"context"
	"mm-pddikti-cms/internal/module/user/ports"
	"mm-pddikti-cms/pkg/response"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserHandler struct {
	userService ports.UserService
}

func NewUserHandler(userService ports.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Profile(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	id := ctx.Locals("user_id").(uuid.UUID)

	user, err := h.userService.Profile(context, id)
	if err != nil {
		return response.SendResponse(ctx, response.ResponseParams{
			StatusCode: fiber.StatusNotFound,
			Message:    err.Error(),
		})
	}

	return response.SendResponse(ctx, response.ResponseParams{
		StatusCode: fiber.StatusAccepted,
		Data:       user,
	})
}
