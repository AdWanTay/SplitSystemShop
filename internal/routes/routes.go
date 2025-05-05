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

	app.Get("/api/split-system/:id", middlewares.RequireAuth(cfg, false), handlers.GetSplitSystem(ctx.SplitSystemService))
	app.Get("/api/split-system", middlewares.RequireAuth(cfg, false), handlers.GetAllSplitSystems(ctx.SplitSystemService))
	app.Delete("/api/split-system/:id", middlewares.RequireAuth(cfg, false), middlewares.RequireAdmin(ctx.UserService), handlers.DeleteSplitSystem(ctx.SplitSystemService))

	app.Get("/api/favorites", middlewares.RequireAuth(cfg, false), handlers.GetFavorites(ctx.FavoritesService))
	app.Delete("/api/favorites/:id", middlewares.RequireAuth(cfg, false), handlers.DeleteFavoritesItem(ctx.FavoritesService))
	app.Post("/api/favorites", middlewares.RequireAuth(cfg, false), handlers.AddToFavorites(ctx.FavoritesService))

	app.Get("/api/cart", middlewares.RequireAuth(cfg, false), handlers.GetCart(ctx.CartService))
	app.Delete("/api/cart/:id", middlewares.RequireAuth(cfg, false), handlers.DeleteCartItem(ctx.CartService))
	app.Post("/api/cart", middlewares.RequireAuth(cfg, false), handlers.AddToCart(ctx.CartService))

	//app.Patch("/api/auth/change-name", middlewares.RequireAuth(cfg, false), handlers.ChangeName(userService))
	//app.Patch("/api/auth/change-phone", middlewares.RequireAuth(cfg, false), handlers.ChangePhoneNumber(userService))
	//app.Delete("/api/auth/delete-account", middlewares.RequireAuth(cfg, false), handlers.DeleteAccount(userService))

	//Роуты для фронта
	app.Get("/", handlers.IndexPage(cfg))
	app.Get("/admin", handlers.AdminPage(cfg))
	app.Get("/article", handlers.ArticlePage(cfg))
	app.Get("/cart", handlers.CartPage(cfg))
	app.Get("/catalog", handlers.CatalogPage(cfg))
	app.Get("/contact", handlers.ContactPage(cfg))
	app.Get("/products", handlers.ProductPage(cfg))
	app.Get("/profile", handlers.ProfilePage(cfg))
	app.Get("/blog", handlers.BlogPage(cfg))

}
