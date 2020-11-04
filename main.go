package main

import (
	"pesticide/database"
	log "pesticide/logHandler"
	"pesticide/models/ticket"
	"pesticide/models/user"
	"pesticide/router"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var fromName = "[main.go]"

func main() {
	app := fiber.New()
	log.InitLog(fromName, "Logs initialized")

	app.Use(middleware.Recover())
	app.Use(cors.New())

	initDatabase()
	router.SetupRoutes(app)
	app.Listen(3000)
}

func initDatabase() {
	var err error

	database.DBConn, err = gorm.Open(sqlite.Open("pesticide.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	log.Debug(fromName, "Connection Opened to the Database")
	database.DBConn.AutoMigrate(&ticket.Ticket{})
	log.Debug(fromName, "Ticket Table Migrated")
	database.DBConn.AutoMigrate(&user.User{})
	log.Debug(fromName, "User Table Migrated")

}
