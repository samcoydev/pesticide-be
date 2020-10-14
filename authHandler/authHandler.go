package authHandler

import (
	"fmt"
	"pesticide/database"
	"pesticide/user"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/gofiber/fiber"
)

type User struct {
	gorm.Model
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Token     string `json:"token"`
}

func Register(c *fiber.Ctx) {
	fmt.Println("Someone registered!")
	db := database.DBConn
	u := new(user.User)

	// Unpack http request data
	if err := c.BodyParser(u); err != nil {
		c.Status(503).Send(err)
		return
	}

	// Encrypt password here
	if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 8); err != nil {
		c.Status(503).Send(err)
		return
	} else {
		// Change password to the hashed version
		u.Password = string(hashedPassword)

		// Add a row with user object
		db.Create(&u)
		c.JSON(u)
	}
}

func Authenticate(c *fiber.Ctx) {
	fmt.Println("Someone is logging in")
	db := database.DBConn
	u := new(user.User)
	storedUser := new(user.User)

	// Unpack http request data
	if err := c.BodyParser(u); err != nil {
		c.Status(503).Send(err)
		return
	}

	fmt.Println(string(u.Username))

	// Finds the user object in database that matches the hhtp requests id
	if storedUser := db.Find(&u, string(u.Username)); storedUser != nil {
		fmt.Println("Unauthorized")
		c.Send(storedUser)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(u.Password)); err != nil {
		fmt.Println("Wrong password")
		c.Send(err)
		return
	}

	fmt.Println("User logged in!", u.Username)
}
