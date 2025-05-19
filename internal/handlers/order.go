package handlers

import (
	"SplitSystemShop/internal/config"
	"SplitSystemShop/internal/context"
	"SplitSystemShop/internal/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func CreateOrder(cfg *config.Config, ctx *context.AppContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("userId").(uint)
		order, err := ctx.OrderService.CreateOrderByUserCart(c.Context(), userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
		}
		user, err := ctx.UserService.GetUser(c.Context(), userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
		}

		err = utils.SendNewOrderNotification(user.Email, order, cfg)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "ок"})
	}
}

func UpdateOrderStatus(cfg *config.Config, ctx *context.AppContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		orderID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		}

	}
}
