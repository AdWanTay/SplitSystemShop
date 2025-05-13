package handlers

import (
	"SplitSystemShop/internal/config"
	"SplitSystemShop/internal/context"
	"SplitSystemShop/internal/services"
	"SplitSystemShop/internal/utils"
	"github.com/gofiber/fiber/v2"
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

func CartPage(cfg *config.Config, userService *services.UserService) fiber.Handler {

	return func(c *fiber.Ctx) error {
		userID := c.Locals("userId").(uint)

		cart, err := userService.GetCart(c.Context(), userID)
		if err != nil {
			return err
		}

		favorites, err := userService.GetFavorites(c.Context(), userID)
		if err != nil {
			return err
		}

		return Render(c, "cart", fiber.Map{
			"cart": fiber.Map{
				"total": len(cart),
				"items": cart,
			},
			"favorites": fiber.Map{
				"total": len(favorites),
				"items": favorites,
			},
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

func ProfilePage(cfg *config.Config, userService *services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("userId").(uint)

		cart, err := userService.GetCart(c.Context(), userID)
		if err != nil {
			return err
		}

		favorites, err := userService.GetFavorites(c.Context(), userID)
		if err != nil {
			return err
		}

		return Render(c, "profile", fiber.Map{
			"cart": fiber.Map{
				"total": len(cart),
				"items": cart,
			},
			"favorites": fiber.Map{
				"total": len(favorites),
				"items": favorites,
			},
		}, cfg)
	}
}

func ProductPage(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return Render(c, "product", fiber.Map{}, cfg)
	}
}
