package handlers

import (
	"SplitSystemShop/internal/dto"
	"SplitSystemShop/internal/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func CreateArticle(articleService *services.ArticleService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req dto.NewArticleRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Некорректные данные"})
		}
		article, err := articleService.Create(c.Context(), req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusCreated).JSON(article)
	}
}

func GetArticle(articleService *services.ArticleService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Неверный ID"})
		}
		article, err := articleService.GetByID(c.Context(), uint(id))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Статья не найдена"})
		}
		return c.JSON(article)
	}
}

func GetAllArticles(articleService *services.ArticleService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		articles, err := articleService.GetAll(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(articles)
	}
}

func DeleteArticle(articleService *services.ArticleService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Неверный ID"})
		}
		if err := articleService.Delete(c.Context(), uint(id)); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.SendStatus(fiber.StatusNoContent)
	}
}

func UpdateArticle(articleService *services.ArticleService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Неверный ID"})
		}
		var req dto.NewArticleRequest
		if err = c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Некорректные данные"})
		}
		updatedArticle, err := articleService.Update(c.Context(), uint(id), req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(updatedArticle)
	}
}
