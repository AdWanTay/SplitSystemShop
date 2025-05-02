package app

import (
	"SplitSystemShop/internal/config"
	"SplitSystemShop/internal/context"
	"SplitSystemShop/internal/database"
	"SplitSystemShop/internal/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"strings"
)

func App(cfg *config.Config) error {
	db, err := database.GetConnection(cfg.Database)
	if err != nil {
		return fmt.Errorf("get database connection: %w", err)
	}

	ctx := context.InitServices(db)

	engine := html.New("./web/templates", ".html")
	engine.Reload(true)

	engine.AddFunc("split", func(s, sep string) []string {
		return strings.Split(s, sep)
	})

	app := fiber.New(fiber.Config{
		Views: engine,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			switch code {
			case fiber.StatusUnauthorized:
				return c.Status(code).Render("errors/401", fiber.Map{})
			case fiber.StatusNotFound:
				return c.Status(code).Render("errors/404", fiber.Map{})
			case fiber.StatusNotAcceptable:
				return c.Status(code).Render("errors/406", fiber.Map{})
			default:
				fmt.Println(err)
				return c.Status(code).Render("errors/500", fiber.Map{
					"error": err.Error(),
				})
			}
		},
	})

	routes.SetupRoutes(app, cfg, ctx)
	err = app.Listen(":" + cfg.Port)
	if err != nil {
		return fmt.Errorf("app listen: %w", err)
	}
	return nil
}
