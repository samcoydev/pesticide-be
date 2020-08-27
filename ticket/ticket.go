package ticket

import (
	"github.com/gofiber/fiber"
)

func GetTickets(c *fiber.Ctx) {
	c.Send("All Ticket")
}

func GetTicket(c *fiber.Ctx) {
	c.Send("Single Ticket")
}

func NewTicket(c *fiber.Ctx) {
	c.Send("New Ticket")
}

func DeleteTicket(c *fiber.Ctx) {
	c.Send("Delete Ticket")
}
