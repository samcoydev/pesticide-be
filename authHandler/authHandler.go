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

func Register(ctx *fiber.Ctx) {
	db := database.DBConn
	user := new(User)

	// Create a user object from the posted data in "ctx"
	if err := ctx.BodyParser(user); err != nil {
		fmt.Println("Error parsing")
		ctx.Status(503).Send(err)
		return
	}

	user.Password = encryptPassword(user.Password)

	db.Create(&user)
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

	if err := verifyPassword(storedUser.Password, user.Password); err != nil {
		fmt.Println("Passwords dont match")
	}

	fmt.Println("User logged in!")
}

func encryptPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		fmt.Println("error encrypting password")
	}
	return string(hashedPassword)
}

func verifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
