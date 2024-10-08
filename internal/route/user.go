package route

import (
	"mm-pddikti-cms/internal/middleware"
	handler "mm-pddikti-cms/internal/module/user/handler/rest"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app fiber.Router, userHandler *handler.UserHandler) {

	group := app.Group("/user")

	group.Use(func(c *fiber.Ctx) error {
		return middleware.AuthMiddleware(c)
	})

	group.Get("/profile", userHandler.Profile)
}
