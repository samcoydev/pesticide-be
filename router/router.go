package router

import (
	"pesticide/authhandler"
	"pesticide/models/ticket"
	"pesticide/models/user"

	"github.com/gofiber/fiber"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/v1/ticket", ticket.GetTickets)
	app.Get("/api/v1/ticket/:id", ticket.GetTicket)
	app.Post("/api/v1/ticket", ticket.NewTicket)
	app.Put("/api/v1/ticket/:id", ticket.UpdateTicket)
	app.Delete("/api/v1/ticket/:id", ticket.DeleteTicket)

	app.Get("/api/v1/users", user.GetUsers)
	app.Get("api/v1/users/:id/tickets", user.GetAssignedTickets)
	app.Post("/api/v1/users/register", authhandler.Register)
	app.Post("/api/v1/users/authenticate", authhandler.Authenticate)
}
