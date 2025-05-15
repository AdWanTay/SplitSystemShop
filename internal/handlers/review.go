package handlers

import (
	"SplitSystemShop/internal/dto"
	"SplitSystemShop/internal/services"
	"github.com/gofiber/fiber/v2"
)

func CreateReview(reviewService *services.ReviewService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input dto.NewReviewRequest
		userID := c.Locals("userId").(uint)
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}

		err := reviewService.Create(c.Context(), input, userID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		}
		return nil
	}
}
