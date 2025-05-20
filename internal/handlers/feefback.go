package handlers

import (
	"SplitSystemShop/internal/config"
	"SplitSystemShop/internal/dto"
	"SplitSystemShop/internal/utils"
	"github.com/gofiber/fiber/v2"
	"log"
)

func SendFeedback(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input dto.FeedbackRequest
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}
		go func() {
			err := utils.SendFeedback(input, cfg)
			if err != nil {
				log.Println("Ошибка отправки письма:", err)
			}
		}()

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Сообщение отправлено",
		})
	}
}
