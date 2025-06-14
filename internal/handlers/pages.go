package handlers

import (
	"SplitSystemShop/internal/config"
	"SplitSystemShop/internal/context"
	"SplitSystemShop/internal/dto"
	"SplitSystemShop/internal/models"
	"SplitSystemShop/internal/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"html/template"
	"math"
	"strconv"
)

func Render(c *fiber.Ctx, template string, data fiber.Map, cfg *config.Config) error {
	if data == nil {
		data = fiber.Map{}
	}

	tokenString := c.Cookies("token")
	_, err := utils.ParseAndValidateJWT(tokenString, cfg)

	data["isAuthenticated"] = err == nil

	return c.Render(template, data)
}

func IndexPage(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return Render(c, "index", fiber.Map{}, cfg)
	}
}

func AdminPage(cfg *config.Config, ctx *context.AppContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		brands, err := ctx.BrandService.GetAll(c.Context())
		types, err := ctx.TypeService.GetAll(c.Context())
		modes, err := ctx.ModeService.GetAll(c.Context())
		energyClasses, err := ctx.EnergyClassService.GetAll(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
		}

		return Render(c, "admin-panel", fiber.Map{
			"brands":         brands,
			"types":          types,
			"modes":          modes,
			"energy_classes": energyClasses,
		}, cfg)
	}
}
func ArticlePage(cfg *config.Config, appContext *context.AppContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idParam := c.Params("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Некорректный ID")
		}

		article, err := appContext.ArticleService.GetByID(c.Context(), uint(id))
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Статья не найдена")
		}

		// Можно добавить 2-3 других статьи для "Вам может быть интересно"
		related, _ := appContext.ArticleService.GetRandomExcept(c.Context(), uint(id), 3)
		return Render(c, "article", fiber.Map{
			"article": fiber.Map{
				"ID":          article.ID,
				"Title":       article.Title,
				"Description": article.Description,
				"ImageURL":    article.ImageURL,
				"Content":     template.HTML(article.Content),
			},
			"related": related,
		}, cfg)
	}
}

func BlogPage(cfg *config.Config, appContext *context.AppContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var articles []models.Article
		articles, _ = appContext.ArticleService.GetAll(c.Context())

		// По умолчанию считаем, что не админ
		isAdmin := false

		// Без паники, если пользователь не залогинен
		if rawID := c.Locals("userId"); rawID != nil {
			fmt.Println("тест")

			if userID, ok := rawID.(uint); ok {
				user, err := appContext.UserService.GetUser(c.Context(), userID)
				fmt.Println("2")
				if err == nil && user.Role == "admin" {
					isAdmin = true
				}
			}
		}

		fmt.Println(isAdmin)

		return Render(c, "blog", fiber.Map{
			"articles": articles,
			"isAdmin":  isAdmin,
		}, cfg)
	}
}

func CartPage(cfg *config.Config, appContext *context.AppContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("userId").(uint)
		response, err := appContext.CartService.LoadCartModuleData(c.Context(), userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
		}
		return Render(c, "cart", fiber.Map{
			"hasProcessingOrders": appContext.UserService.HasProcessingOrders(c.Context(), userID),
			"response":            response,
		}, cfg)
	}
}

func CatalogPage(cfg *config.Config, ctx *context.AppContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		brands, err := ctx.BrandService.GetAll(c.Context())
		types, err := ctx.TypeService.GetAll(c.Context())
		modes, err := ctx.ModeService.GetAll(c.Context())
		energyClasses, err := ctx.EnergyClassService.GetAll(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
		}

		return Render(c, "catalog", fiber.Map{
			"brands":         brands,
			"types":          types,
			"modes":          modes,
			"energy_classes": energyClasses,
		}, cfg)
	}
}

func ContactPage(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return Render(c, "contact", fiber.Map{}, cfg)
	}
}

func ProfilePage(cfg *config.Config, appContext *context.AppContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("userId").(uint)
		cartModuleData, err := appContext.CartService.LoadCartModuleData(c.Context(), userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
		}

		user, err := appContext.UserService.GetUser(c.Context(), userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
		}
		return Render(c, "profile", fiber.Map{
			"hasProcessingOrders": appContext.UserService.HasProcessingOrders(c.Context(), userID),
			"response":            cartModuleData,
			"userData":            user,
		}, cfg)
	}
}

func ProductPage(cfg *config.Config, appContext *context.AppContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		}
		splitSystem, err := appContext.SplitSystemService.GetSplitSystem(c.Context(), uint(id))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
		}
		splitSystem.AverageRating = math.Round(splitSystem.AverageRating*10) / 10
		for i := range splitSystem.Reviews {
			if splitSystem.Reviews[i].User == nil {
				splitSystem.Reviews[i].User = &models.User{FirstName: "Профиль удален"}
			} else {
				original := splitSystem.Reviews[i].User
				copied := *original

				runes := []rune(copied.LastName)
				if len(runes) > 0 {
					postfix := string(runes[0]) + "."
					copied.FirstName += " " + postfix
				}
				splitSystem.Reviews[i].User = &copied
			}
		}
		inCart := false
		inFavorites := false
		userId := c.Locals("userId")
		if userId != nil {
			inCart = appContext.UserService.IsInCart(c.Context(), userId.(uint), uint(id))
			inFavorites = appContext.UserService.IsInFavorites(c.Context(), userId.(uint), uint(id))
		}

		response := dto.CatalogItem{
			SplitSystem: *splitSystem,
			InCart:      inCart,
			InFavorites: inFavorites,
		}
		return Render(c, "product", fiber.Map{"info": response}, cfg)
	}
}
