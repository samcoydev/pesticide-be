package main

import (
	"github.com/gofiber/fiber"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
}

func main() {
	app := fiber.New()

	setupRoutes(app)
	app.Listen(3000)
}
