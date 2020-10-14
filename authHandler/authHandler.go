package authHandler

import (
	"fmt"
	"pesticide/database"
	"pesticide/user"

	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber"
)

func Register(c *fiber.Ctx) {
	fmt.Println("Someone registered")
	db := database.DBConn
	u := new(user.User)

	if err := c.BodyParser(u); err != nil {
		c.Status(503).Send(err)
		return
	}

	// Encrypt password here
	if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 8); err != nil {
		c.Status(503).Send(err)
		return
	} else {
		u.Password = string(hashedPassword)
	}

	// Add a row with user object
	db.Create(&u)
	c.JSON(u)
}

func Authenticate(c *fiber.Ctx) {
	fmt.Println("Someone is logging in")
	u := new(user.User)

	if err := c.BodyParser(u); err != nil {
		c.Send(err)
		return
	}

	// This is what we get in the http post request
	fmt.Println(u.Username)
	fmt.Println(u.Password)
}
