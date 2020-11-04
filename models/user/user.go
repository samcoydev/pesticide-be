package user

import (
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
	"pesticide/models/roles"

	"pesticide/database"
	log "pesticide/logHandler"
	"pesticide/models/ticket"
)

type User struct {
	gorm.Model
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	FirstName string     `json:"firstname"`
	LastName  string     `json:"lastname"`
	Email     string     `json:"email"`
	Role      roles.Role `json:"role"`
	Token     string     `json:"token"`
}

var fromName string = "[user.go]"

func GetUsers(ctx *fiber.Ctx) {
	log.Debug(fromName, "Get Users")
	db := database.DBConn
	var users []User
	db.Find(&users)
	ctx.JSON(users)
}

func GetAssignedTickets(ctx *fiber.Ctx) {
	log.Debug(fromName, "Get users assigned tickets")
	db := database.DBConn
	id := ctx.Params("id")
	var tickets []ticket.Ticket
	var user User

	log.Debug(fromName, id)

	db.Find(&user, id)

	db.Table("tickets").Where("assigned_username = ?", user.Username).Find(&tickets)
	ctx.JSON(tickets)
}
