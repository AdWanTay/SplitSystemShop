package handlers

import (
	"SplitSystemShop/internal/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

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

func GetAllSplitSystems(service *services.SplitSystemService) fiber.Handler {
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
