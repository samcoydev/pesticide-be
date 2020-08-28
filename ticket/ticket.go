package ticket

import (
	"pesticide/database"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Ticket struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Timestamp   string `json:"timestamp"`
}

func GetTickets(c *fiber.Ctx) {
	db := database.DBConn
	var tickets []Ticket
	db.Find(&tickets)
	c.JSON(tickets)
}

func GetTicket(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var ticket Ticket
	db.Find(&ticket, id)
	c.JSON(ticket)
}

func NewTicket(c *fiber.Ctx) {
	db := database.DBConn
	var ticket Ticket
	ticket.Title = "Test ticket"
	ticket.Description = "Testing our ticket system."
	ticket.Timestamp = "Aug 28th"
	db.Create(&ticket)
	c.JSON(ticket)
}

func DeleteTicket(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var ticket Ticket
	db.First(&ticket, id)
	if ticket.Title == "" {
		c.Status(500).Send("No ticket found with ID")
		return
	}
	db.Delete(&ticket)
	c.Send("Ticket Successfully deleted")
}
