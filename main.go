package main

import (
	"fmt"
	"pesticide/database"
	"pesticide/ticket"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
	app.Get("/api/v1/createfaketicket", ticket.NewFakeTicket)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("tickets.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&ticket.Ticket{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()

	// handle panics and don't kill the server!
	app.Use(middleware.Recover())
	app.Use(cors.New())

	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
}
