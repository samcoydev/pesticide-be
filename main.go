package main

import (
	"fmt"
	"pesticide/database"
	"pesticide/models/ticket"
	"pesticide/router"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	app.Use(middleware.Recover())
	app.Use(cors.New())

	initDatabase()
	router.SetupRoutes(app)
	app.Listen(3000)
}

func initDatabase() {
	var err error

	database.DBConn, err = gorm.Open(sqlite.Open("tickets.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	fmt.Println("Connection Opened to Ticket Database")
	database.DBConn.AutoMigrate(&ticket.Ticket{})
	fmt.Println("Ticket Database Migrated")
}
