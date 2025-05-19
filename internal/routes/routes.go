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

	//Роуты для апи
	app.Post("/api/auth/login", handlers.Login(ctx.UserService, cfg))
	app.Post("/api/auth/register", handlers.Registration(ctx.UserService, cfg))
	app.Get("/api/auth/logout", middlewares.RequireAuth(cfg, false), handlers.Logout())
	app.Patch("/api/auth/profile", middlewares.RequireAuth(cfg, false), handlers.ChangeCredentials(ctx.UserService))
	app.Patch("/api/auth/change-password", middlewares.RequireAuth(cfg, false), handlers.ChangePassword(ctx.UserService))

	app.Get("/api/split-systems/:id", middlewares.RequireAuth(cfg, false), handlers.GetSplitSystem(ctx.SplitSystemService))
	app.Get("/api/split-systems", middlewares.RequireAuth(cfg, true), handlers.GetAllSplitSystems(ctx.SplitSystemService, ctx.UserService))
	app.Delete("/api/split-systems/:id", middlewares.RequireAuth(cfg, false), middlewares.RequireAdmin(ctx.UserService), handlers.DeleteSplitSystem(ctx.SplitSystemService))
	app.Post("/api/split-systems", middlewares.RequireAuth(cfg, false), middlewares.RequireAdmin(ctx.UserService), handlers.CreateSplitSystem(ctx.SplitSystemService))
	app.Patch("/api/split-systems/:id", middlewares.RequireAuth(cfg, false), middlewares.RequireAdmin(ctx.UserService), handlers.UpdateSplitSystem(ctx.SplitSystemService))

	app.Get("/api/favorites", middlewares.RequireAuth(cfg, false), handlers.GetFavorites(ctx.FavoritesService))
	app.Delete("/api/favorites/:id", middlewares.RequireAuth(cfg, false), handlers.DeleteFavoritesItem(ctx.FavoritesService))
	app.Post("/api/favorites", middlewares.RequireAuth(cfg, false), handlers.AddToFavorites(ctx.FavoritesService))

	app.Get("/api/cart", middlewares.RequireAuth(cfg, false), handlers.GetCart(ctx.CartService))
	app.Delete("/api/cart/:id", middlewares.RequireAuth(cfg, false), handlers.DeleteCartItem(ctx.CartService))
	app.Post("/api/cart", middlewares.RequireAuth(cfg, false), handlers.AddToCart(ctx.CartService))

	app.Post("/api/order", middlewares.RequireAuth(cfg, false), handlers.CreateOrder(cfg, ctx))
	app.Patch("/api/order/:id", middlewares.RequireAuth(cfg, false), middlewares.RequireAdmin(ctx.UserService), handlers.UpdateOrderStatus(cfg, ctx))
	app.Get("/api/order", middlewares.RequireAuth(cfg, false), middlewares.RequireAdmin(ctx.UserService), handlers.GetAll(ctx))

	app.Post("/api/review", middlewares.RequireAuth(cfg, false), handlers.CreateReview(ctx))

	app.Get("/api/articles/:id", middlewares.RequireAuth(cfg, false), handlers.GetArticle(ctx.ArticleService))
	app.Get("/api/articles/:id", middlewares.RequireAuth(cfg, false), handlers.GetAllArticles(ctx.ArticleService))
	app.Post("/api/articles", middlewares.RequireAuth(cfg, false), middlewares.RequireAdmin(ctx.UserService), handlers.CreateArticle(ctx.ArticleService))
	app.Delete("/api/articles/:id", middlewares.RequireAuth(cfg, false), middlewares.RequireAdmin(ctx.UserService), handlers.DeleteArticle(ctx.ArticleService))
	app.Patch("/api/articles/:id", middlewares.RequireAuth(cfg, false), middlewares.RequireAdmin(ctx.UserService), handlers.UpdateArticle(ctx.ArticleService))

	app.Patch("/api/auth/change-credentials", middlewares.RequireAuth(cfg, false), handlers.ChangeCredentials(ctx.UserService))
	app.Delete("/api/auth/delete-account", middlewares.RequireAuth(cfg, false), handlers.DeleteAccount(ctx.UserService))

	//Роуты для фронта
	app.Get("/", handlers.IndexPage(cfg))
	app.Get("/admin", middlewares.RequireAuth(cfg, false), middlewares.RequireAdmin(ctx.UserService), handlers.AdminPage(cfg, ctx))

	app.Get("/articles/:id", handlers.ArticlePage(cfg, ctx))

	app.Get("/cart", middlewares.RequireAuth(cfg, false), handlers.CartPage(cfg, ctx))
	app.Get("/catalog", handlers.CatalogPage(cfg, ctx))
	app.Get("/contact", handlers.ContactPage(cfg))
	app.Get("/products/:id", middlewares.RequireAuth(cfg, true), handlers.ProductPage(cfg, ctx))
	app.Get("/profile", middlewares.RequireAuth(cfg, false), handlers.ProfilePage(cfg, ctx))
	app.Get("/blog", middlewares.RequireAuth(cfg, true), handlers.BlogPage(cfg, ctx))

}
