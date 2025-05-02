package middlewares

import (
	"SplitSystemShop/internal/config"
	"SplitSystemShop/internal/utils"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func RequireAuth(cfg *config.Config, onlyUserId bool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Cookies("token")
		if tokenString == "" && !onlyUserId {
			return fiber.NewError(fiber.StatusUnauthorized, "Missing token")
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		userId, err := utils.ParseAndValidateJWT(tokenString, cfg)
		if err != nil && !onlyUserId {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
		}
		if err == nil {
			c.Locals("userId", userId)
		}
		return c.Next()
	}
}
