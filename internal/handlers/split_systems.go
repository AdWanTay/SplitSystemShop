package handlers

import (
	"SplitSystemShop/internal/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetSplitSystem(service *services.SplitSystemService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		splitSystemID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Неверный id товара"})
		}

		splitSystem, err := service.GetSplitSystem(c.Context(), uint(splitSystemID))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Ошибка получения товара",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"item": splitSystem,
		})
	}
}

func GetAllSplitSystems(service *services.SplitSystemService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		splitSystems, err := service.GetAllSplitSystems(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Ошибка получения товаров"})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"items": splitSystems,
		})
	}
}

func DeleteSplitSystem(service *services.SplitSystemService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		splitSystemID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Неверный id товара"})
		}
		if service.DeleteSplitSystem(c.Context(), uint(splitSystemID)) != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Ошибка удаления товара"})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Товар успешно удален",
		})
	}
}
