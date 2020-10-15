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
	//db := database.DBConn
	user := new(User)

	// Create a user object from the posted data in "ctx"
	if err := c.BodyParser(user); err != nil {
		c.Status(503).Send(err)
		return
	}

	dbUser, err := findUserByUsername(user.Username)
	if err != nil {
		fmt.Println("Cannot find use111")
	}

	if err := verifyPassword(dbUser.Password, user.Password); err != nil {
		fmt.Println("Passwords dont match")
	}

	fmt.Println("User logged in!")
}

func findUserByUsername(username string) (User, error) {
	db := database.DBConn
	var dbUser User
	//rows, err := db.Debug().Model(&User{}).Where("username = ?", username).Select("Username, Password").Rows()
	rows, err := db.Table("users").Where("username = ?", username).Select("Username, Password").Rows()
	for rows.Next() {
		db.ScanRows(rows, &dbUser)
	}
	return dbUser, err
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
