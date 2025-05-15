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

	engine.AddFunc("stars", func(rating int) string {
		var stars string
		for i := 1; i <= 5; i++ {
			if i <= rating {
				stars += "★"
			} else {
				stars += "☆"
			}
		}
		return stars
	})

	engine.AddFunc("formatPrice", func(price int) string {
		s := fmt.Sprintf("%d", price)
		n := len(s)

		if n <= 3 {
			return s
		}

		var b strings.Builder
		pre := n % 3
		if pre > 0 {
			b.WriteString(s[:pre])
			if n > pre {
				b.WriteString(" ")
			}
		}

		for i := pre; i < n; i += 3 {
			b.WriteString(s[i : i+3])
			if i+3 < n {
				b.WriteString(" ")
			}
		}
		return b.String()
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
