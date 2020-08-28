package main

import (
	"fmt"
	"pesticide/database"
	"pesticide/ticket"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "tickets.db")
	if err != nil {
		panic("failed to connect to database")
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&ticket.Ticket{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()

	setupRoutes(app)
	app.Listen(3000)

	defer database.DBConn.Close()
}
