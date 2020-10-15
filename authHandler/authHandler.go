package authHandler

import (
	"fmt"
	"pesticide/database"

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
	u := new(User)

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
	user := new(User)
	var storedUser User

	// Unpack http request data
	if err := c.BodyParser(user); err != nil {
		fmt.Println("Error parsing")
		c.Status(503).Send(err)
		return
	}

	// Get object from database
	rows, err := db.Debug().Model(&User{}).Where("username = ?", user.Username).Select("Username, Password").Rows()
	if err != nil {
		fmt.Println("Ran into issue")
		return
	}
	for rows.Next() {
		db.ScanRows(rows, &storedUser)
	}

	if err := VerifyPassword(storedUser.Password, user.Password); err != nil {
		fmt.Println("Passwords dont match")
	}

	fmt.Println("User logged in!")
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
