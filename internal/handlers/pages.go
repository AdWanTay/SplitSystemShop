package handlers

import (
	"SplitSystemShop/internal/config"
	"github.com/gofiber/fiber/v2"
)

func Render(c *fiber.Ctx, template string, data fiber.Map, cfg *config.Config) error {
	if data == nil {
		data = fiber.Map{}
	}

	//tokenString := c.Cookies("token")
	//userId, err := utils.ParseAndValidateJWT(tokenString, cfg)
	//
	//if err == nil {
	//	data["firstName"], data["lastName"], _ = userService.GetFirstNameAndLastName(c.Context(), userId)
	//}
	//
	//data["isAuthenticated"] = err == nil

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

func CartPage(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return Render(c, "cart", fiber.Map{}, cfg)
	}
}

func CatalogPage(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return Render(c, "catalog", fiber.Map{}, cfg)
	}
}

func ContactPage(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return Render(c, "contact", fiber.Map{}, cfg)
	}
}

func ProfilePage(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return Render(c, "profile", fiber.Map{}, cfg)
	}
}

func ProductPage(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return Render(c, "product", fiber.Map{}, cfg)
	}
}
