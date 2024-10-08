package middleware

import (
	"mm-pddikti-cms/pkg/jwthandler"
	"mm-pddikti-cms/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Get the access_token cookie
	cookie := c.Cookies("access_token")

	// If the cookie is not set, return an unauthorized status
	if cookie == "" {
		log.Error().Msg("middleware::AuthMiddleware - Unauthorized [Cookie not set]")
		return response.SendResponse(c, response.ResponseParams{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Unauthorized",
		})
	}

	// Parse the JWT string and store the result in `claims`
	claims, err := jwthandler.ParseTokenString(cookie)
	if err != nil {
		log.Error().Err(err).Msg("middleware::AuthMiddleware - Error while parsing token")
		return response.SendResponse(c, response.ResponseParams{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Bad request",
		})
	}

	c.Locals("user_id", claims.UserId)
	c.Locals("role", claims.Role)

	// If the token is valid, pass the request to the next handler
	return c.Next()
}
