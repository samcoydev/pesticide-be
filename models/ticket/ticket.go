package ticket

import (
	"pesticide/database"
	log "pesticide/logHandler"
	"time"

	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	Timestamp         time.Time `json:"timestamp"`
	AssignedUserEmail string    `json:"assigneduseremail"`
	Creator           string    `json:"creator"`
	PriorityLevel     int       `json:"prioritylevel"`
	Deadline          time.Time `json:"deadline"`
	Status            int       `json:"status"`
}

var fromName string = "[ticket.go]"

func GetTickets(c *fiber.Ctx) {
	log.Debug(fromName, "Get tickets")
	db := database.DBConn
	var tickets []Ticket
	db.Find(&tickets)
	c.JSON(tickets)
}

func GetTicket(c *fiber.Ctx) {
	log.Debug(fromName, "Get ticket")
	id := c.Params("id")
	db := database.DBConn
	var ticket Ticket
	db.Find(&ticket, id)
	c.JSON(ticket)
}

func NewTicket(c *fiber.Ctx) {
	log.Debug(fromName, "New Ticket")
	db := database.DBConn
	ticket := new(Ticket)
	if err := c.BodyParser(ticket); err != nil {
		c.Status(401).Send(err)
		return
	}
	db.Create(&ticket)
	c.JSON(ticket)
}

func DeleteTicket(c *fiber.Ctx) {
	log.Debug(fromName, "Delete ticket")
	id := c.Params("id")
	db := database.DBConn

	var ticket Ticket
	db.First(&ticket, id)
	if ticket.Title == "" {
		c.Status(401).Send("No ticket found with ID")
		return
	}
	db.Delete(&ticket)
	c.Send("Ticket Successfully deleted")
}

func UpdateTicket(c *fiber.Ctx) {
	log.Debug(fromName, "Update ticket")

	id := c.Params("id")
	db := database.DBConn

	var ticket Ticket
	db.Find(&ticket, id)

	newTicket := new(Ticket)
	if err := c.BodyParser(newTicket); err != nil {
		c.Status(401).Send(err)
		return
	}

	ticket = *newTicket

	db.Save(&ticket)
	c.JSON(ticket)
}
