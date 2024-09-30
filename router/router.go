package router

import "github.com/gofiber/fiber/v2"


func Api(app *fiber.App) {
    api := app.Group("/api")
    api.Get("/hello", func(c *fiber.Ctx) error {
        return c.SendString("Hello world")
    })
}
