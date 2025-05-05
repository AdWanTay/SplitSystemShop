package handlers

import (
	"SplitSystemShop/internal/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func DeleteFavoritesItem(s *services.FavoritesService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		splitSystemID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Неверный id товара"})
		}
		userID := c.Locals("userId").(uint)
		if err = s.RemoveFromFavorites(c.Context(), userID, uint(splitSystemID)); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Ошибка удаления товара"})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "товар успешно удален"})
	}
}

func AddToFavorites(s *services.FavoritesService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type request struct {
			SplitSystemId uint `json:"split_system_id"`
		}

		body := &request{}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Неверный id товара"})
		}

		userID := c.Locals("userId").(uint)
		if err := s.AddToFavorites(c.Context(), userID, body.SplitSystemId); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Ошибка добавления товара"})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "товар успешно добавлен"})
	}
}

func GetFavorites(service *services.FavoritesService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("userId").(uint)
		items, err := service.GetFavorites(c.Context(), userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Ошибка загрузки товаров",
			})
		}
		return c.JSON(fiber.Map{
			"items": items,
		})
	}
}
