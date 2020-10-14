package main

import (
	"fmt"
	"pesticide/database"
	"pesticide/router"
	"pesticide/ticket"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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
	router.SetupRoutes(app)
	app.Listen(3000)
}
