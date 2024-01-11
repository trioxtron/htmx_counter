package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/trioxtron/golang-service/api"
)

func setupRoutes(app *fiber.App) {
    counter := &api.Counter{}
    app.Static("/", "./src/index.html")
    app.Static("/src/tailwind.css", "./src/tailwind.css")
    app.Get("/countervalue", func(c *fiber.Ctx) error {
        return counter.GetValue(c)
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
	app := fiber.New()
    app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
    }))
	setupRoutes(app)
	app.Listen(":3000")

}
