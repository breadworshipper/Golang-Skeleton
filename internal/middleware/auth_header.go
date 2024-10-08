package middleware

import (
	"mm-pddikti-cms/pkg/jwthandler"
	"mm-pddikti-cms/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func AuthBearer(c *fiber.Ctx) error {
	AccessToken := c.Get("Authorization")

	// If the cookie is not set, return an unauthorized status
	if AccessToken == "" {
		log.Error().Msg("middleware::AuthMiddleware - Unauthorized [Header not set]")
		return response.SendResponse(c, response.ResponseParams{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Unauthorized",
		})
	}

	// remove the Bearer prefix
	if len(AccessToken) > 7 {
		AccessToken = AccessToken[7:]
	}

	// Parse the JWT string and store the result in `claims`
	claims, err := jwthandler.ParseTokenString(AccessToken)
	if err != nil {
		log.Error().Err(err).Any("payload", AccessToken).Msg("middleware::AuthMiddleware - Error while parsing token")
		return response.SendResponse(c, response.ResponseParams{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Unauthorized",
		})
	}

	c.Locals("user_id", claims.UserId)
	c.Locals("role", claims.Role)

	// If the token is valid, pass the request to the next handler
	return c.Next()
}
