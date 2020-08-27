package main

import (
	"pesticide/ticket"

	"github.com/gofiber/fiber"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	app.Get("/api/v1/ticket", ticket.GetTickets)
	app.Get("/api/v1/ticket/:id", ticket.GetTicket)
	app.Post("/api/v1/ticket", ticket.NewTicket)
	app.Delete("/api/v1/ticket/:id", ticket.DeleteTicket)
}

func main() {
	app := fiber.New()

	setupRoutes(app)
	app.Listen(3000)
}
