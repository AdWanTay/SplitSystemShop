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
		status := c.Query("status")
		orderID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		}

		order, err := ctx.OrderService.UpdateOrderStatus(c.Context(), uint(orderID), status)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
		}

		err = utils.SendOrderStatusUpdateNotification(order.User.Email, order, cfg)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "ок"})
	}
}

func GetAll(ctx *context.AppContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		all, err := ctx.OrderService.GetAll(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": all})
	}
}
