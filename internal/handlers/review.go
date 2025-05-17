package handlers

import (
	"SplitSystemShop/internal/context"
	"SplitSystemShop/internal/dto"
	"github.com/gofiber/fiber/v2"
)

func CreateReview(ctx *context.AppContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input dto.NewReviewRequest
		userID := c.Locals("userId").(uint)

		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}
		review, err := ctx.ReviewService.Create(c.Context(), input, userID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		user, err := ctx.UserService.GetUser(c.Context(), userID)
		if err != nil {
			return err
		}
		runes := []rune(user.LastName)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Отзыв успешно добавлен",
			"item": fiber.Map{
				"comment": review.Comment,
				"rating":  review.Rating,
				"name":    user.FirstName + " " + string(runes[0]) + ".",
			}})
	}
}
