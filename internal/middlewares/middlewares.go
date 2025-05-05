package middlewares

import (
	"SplitSystemShop/internal/config"
	"SplitSystemShop/internal/services"
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

		userID, err := utils.ParseAndValidateJWT(tokenString, cfg)
		if err != nil && !onlyUserId {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
		}
		if err == nil {
			c.Locals("userId", userID)
		}
		return c.Next()
	}
}

func RequireAdmin(service *services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("userId").(uint)
		role, err := service.GetUserRole(c.Context(), userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Ошибка получения роли"})
		}

		if role != "admin" {
			return c.Status(fiber.StatusMethodNotAllowed).JSON(fiber.Map{"message": "Нет прав"})
		}

		return c.Next()
	}
}
