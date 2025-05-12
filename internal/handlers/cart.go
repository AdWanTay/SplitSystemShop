package handlers

import (
	"SplitSystemShop/internal/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func DeleteCartItem(s *services.CartService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		splitSystemID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Неверный id товара"})
		}
		userID := c.Locals("userId").(uint)
		if err = s.RemoveFromCart(c.Context(), userID, uint(splitSystemID)); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Ошибка удаления товара"})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "товар успешно удален"})
	}
}

func AddToCart(s *services.CartService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type request struct {
			SplitSystemId uint `json:"split_system_id"`
		}

		body := &request{}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Неверный id товара"})
		}

		userID := c.Locals("userId").(uint)
		if err := s.AddToCart(c.Context(), userID, body.SplitSystemId); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Ошибка добавления товара"})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "товар успешно добавлен"})
	}
}

func GetCart(service *services.CartService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("userId").(uint)
		cart, err := service.GetCart(c.Context(), userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to get cart",
			})
		}
		return c.JSON(fiber.Map{
			"items": cart,
		})
	}
}
