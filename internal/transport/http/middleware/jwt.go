package middleware

import (
	"market/internal/utils/security"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JWT() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get(fiber.HeaderAuthorization)
		if auth == "" {
			return fiber.ErrUnauthorized
		}

		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			return fiber.ErrUnauthorized
		}

		token := parts[1]

		claims, err := security.ParseToken(token)
		if err != nil {
			return fiber.ErrUnauthorized
		}

		userIDRaw, ok := claims["user_id"]
		if !ok {
			return fiber.ErrUnauthorized
		}

		userID, ok := userIDRaw.(string)
		if !ok || userID == "" {
			return fiber.ErrUnauthorized
		}

		c.Locals("user_id", userID)

		return c.Next()
	}
}
