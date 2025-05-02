package routes

import (
	"SplitSystemShop/internal/config"
	"SplitSystemShop/internal/context"
	"SplitSystemShop/internal/handlers"
	"SplitSystemShop/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, cfg *config.Config, ctx *context.AppContext) {
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		c.Set("Pragma", "no-cache")
		c.Set("Expires", "0")
		return c.Next()
	})
	app.Static("/web", "./web", fiber.Static{CacheDuration: 0})

	//Роуты для апи
	app.Post("/api/auth/login", handlers.Login(ctx.UserService, cfg))
	app.Post("/api/auth/register", handlers.Registration(ctx.UserService, cfg))
	app.Get("/api/auth/logout", handlers.Logout())
	app.Patch("/api/auth/profile", middlewares.RequireAuth(cfg, false), handlers.ChangeCredentials(ctx.UserService))
	app.Patch("/api/auth/change-password", middlewares.RequireAuth(cfg, false), handlers.ChangePassword(ctx.UserService))
	//app.Patch("/api/auth/change-name", middlewares.RequireAuth(cfg, false), handlers.ChangeName(userService))
	//app.Patch("/api/auth/change-phone", middlewares.RequireAuth(cfg, false), handlers.ChangePhoneNumber(userService))
	//app.Delete("/api/auth/delete-account", middlewares.RequireAuth(cfg, false), handlers.DeleteAccount(userService))

	//Роуты для фронта
	//app.Get("/", handlers.IndexPage(cfg, courseService, userService))

}
