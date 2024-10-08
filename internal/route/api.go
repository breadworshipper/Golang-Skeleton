package route

import (
	// integlocalstorage "mm-pddikti-cms/internal/integration/localstorage"
	// m "mm-pddikti-cms/internal/middleware"

	auth "mm-pddikti-cms/internal/module/auth/handler/rest"
	user "mm-pddikti-cms/internal/module/user/handler/rest"

	"mm-pddikti-cms/pkg/response"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func SetupRoutes(app *fiber.App, authHandler *auth.AuthHandler, userHandler *user.UserHandler) {

	// add /api prefix to all routes
	// app.Get("/storage/private/:filename", m.ValidateSignedURL, storageFile)

	app.Get("/csrf-token", func(c *fiber.Ctx) error {
		csrfToken := c.Locals("csrf")
		if csrfToken == nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "CSRF token not found",
			})
		}
		return c.JSON(fiber.Map{
			"csrfToken": csrfToken,
		})
	})
	
	api := app.Group("/api")
	SetupAuthRoutes(api, authHandler)
	SetupUserRoutes(api, userHandler)

	// fallback route
	app.Use(func(c *fiber.Ctx) error {
		var (
			method = c.Method()                       // get the request method
			path   = c.Path()                         // get the request path
			query  = c.Context().QueryArgs().String() // get all query params
			ua     = c.Get("User-Agent")              // get the request user agent
			ip     = c.IP()                           // get the request IP
		)

		log.Trace().
			Str("method", method).
			Str("path", path).
			Str("query", query).
			Str("ua", ua).
			Str("ip", ip).
			Msg("Route not found.")
		return response.SendResponse(c, response.ResponseParams{
			StatusCode: fiber.StatusNotFound,
			Message:    "Route not found",
		})
	})
}

func storageFile(c *fiber.Ctx) error {
	var (
		fileName = c.Params("filename")
		filePath = filepath.Join("storage", "private", fileName)
	)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Error().Err(err).Any("url", filePath).Msg("handler::storageFile - File not found")
		return response.SendResponse(c, response.ResponseParams{
			StatusCode: fiber.StatusNotFound,
			Message:    "File not found",
		})
	}

	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Error().Err(err).Any("url", filePath).Msg("handler::storageFile - Failed to read file")
		return response.SendResponse(c, response.ResponseParams{
			StatusCode: fiber.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	return c.Send(fileBytes)
}
