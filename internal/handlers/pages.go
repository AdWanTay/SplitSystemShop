package handlers

import (
	"SplitSystemShop/internal/config"
	"SplitSystemShop/internal/context"
	"SplitSystemShop/internal/dto"
	"SplitSystemShop/internal/models"
	"SplitSystemShop/internal/services"
	"SplitSystemShop/internal/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
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

func AdminPage(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return Render(c, "admin-panel", fiber.Map{}, cfg)
	}
}
func ArticlePage(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return Render(c, "article", fiber.Map{}, cfg)
	}
}

func BlogPage(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return Render(c, "blog", fiber.Map{}, cfg)
	}
}

func CartPage(cfg *config.Config, cartService *services.CartService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("userId").(uint)
		response, err := cartService.LoadCartModuleData(c.Context(), userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
		}
		return Render(c, "profile", fiber.Map{"response": response}, cfg)
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
		fmt.Println(user)
		return Render(c, "profile", fiber.Map{
			"response": cartModuleData,
			"userData": user,
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
