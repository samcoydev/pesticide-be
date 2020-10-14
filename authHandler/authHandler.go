package authHandler

import (
	"fmt"
	"pesticide/database"
	"pesticide/user"

	"github.com/gofiber/fiber"
)

func Register(c *fiber.Ctx) {
	fmt.Println("Someone registered")
	db := database.DBConn
	_user := new(user.User)
	if err := c.BodyParser(_user); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&_user)
	c.JSON(_user)
}

func Login(c *fiber.Ctx) {
	fmt.Println("Someone is logging in")
	c.Send(c.Body())
}
