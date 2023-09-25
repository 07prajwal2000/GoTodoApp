package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
)

func main() {
	htmlEngine := handlebars.New("./templates", ".hbs")
	htmlEngine.ShouldReload = true

	app := fiber.New(fiber.Config{
		Views: htmlEngine,
	})

	app.Static("/assets", "./static", fiber.Static{
		CacheDuration: time.Hour * 6,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", &fiber.Map{
			"hello": "world 1",
		})
	})

	app.Listen(":3000")
}
