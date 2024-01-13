package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"github.com/trioxtron/golang-service/api"
)

func setupRoutes(app *fiber.App) {
	counter := &api.Counter{}
	app.Static("/src/tailwind.css", "./src/tailwind.css")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Counter": counter.ReturnValue(),
		})
	})

	app.Get("/countervalue", func(c *fiber.Ctx) error {
		return counter.GetValue(c)
	})

	app.Post("/setvalue", func(c *fiber.Ctx) error {
		return counter.SetValue(c)
	})
	app.Post("/increase", func(c *fiber.Ctx) error {
		counter.Increase()
		return counter.GetValue(c)
	})
	app.Post("/decrease", func(c *fiber.Ctx) error {
		counter.Decrease()
		return counter.GetValue(c)
	})
}

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	setupRoutes(app)
	app.Listen(":3000")
}
