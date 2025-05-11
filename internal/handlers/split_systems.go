package handlers

import (
	"SplitSystemShop/internal/dto"
	"SplitSystemShop/internal/models"
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

func GetAllSplitSystems(splitSystemService *services.SplitSystemService, userService *services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		splitSystems, err := splitSystemService.GetAllSplitSystems(c.Context())
		userID := c.Locals("userId")
		var cart []models.SplitSystem
		if userID != nil {
			cart, err = userService.GetCart(c.Context(), userID.(uint))
		}

		response := dto.CatalogResponse{}
		response.New(cart, splitSystems)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Ошибка получения товаров"})
		}
		return c.Status(fiber.StatusOK).JSON(response)
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
