package handlers

import (
	"SplitSystemShop/internal/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetCart(service *services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Locals("userId").(uint)
		cart, err := service.GetCart(c.Context(), userId)
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

func GetFavorites(service *services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Locals("userId").(uint)
		favorites, err := service.GetFavorites(c.Context(), userId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to get favorites",
			})
		}
		return c.JSON(fiber.Map{
			"items": favorites,
		})
	}
}

func GetSplitSystem(service *services.SplitSystemService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		splitSystemId, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "invalid param",
			})
		}
		splitSystem, err := service.GetSplitSystem(c.Context(), uint(splitSystemId))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to get split-system",
			})
		}
		return c.JSON(splitSystem)
	}
}

func GetAllSplitSystem(service *services.SplitSystemService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		splitSystems, err := service.GetAllSplitSystems(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to get split-system",
			})
		}
		return c.JSON(fiber.Map{
			"items": splitSystems,
		})
	}
}
