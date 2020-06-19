package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/template/html"
	"github.com/markbates/pkger"
)

func main() {
	engine := html.NewFileSystem(pkger.Dir("/views"), ".html")
	app := fiber.New(&fiber.Settings{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) {
		c.Render("index", fiber.Map{
			"Title": "Fiber tests!",
		})
	})

	app.Get("/layouts", func(c *fiber.Ctx) {
		c.Render("index", fiber.Map{
			"Title": "Fiber layout tests!",
		}, "layouts/main")
	})

	app.Listen(3000)
}
