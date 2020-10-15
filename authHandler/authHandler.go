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

	encryptedPassword, err := encryptPassword(user.Password)
	if err != nil {
		fmt.Println("Error encrypting your password")
		ctx.Status(503).Send(err)
		return
	}

	user.Password = encryptedPassword

	db.Create(&user)
}

func Authenticate(ctx *fiber.Ctx) {
	user := new(User)

	// Create a user object from the posted data in "ctx"
	if err := ctx.BodyParser(user); err != nil {
		ctx.Status(503).Send(err)
		return
	}

	dbUser, err := findUserByUsername(user.Username)
	if err != nil {
		fmt.Println("Cannot find user: ", user.Username)
		return
	}

	if err := verifyPassword(dbUser.Password, user.Password); err != nil {
		fmt.Println("Passwords dont match for user: ", user.Username)
		return
	}

	fmt.Println("User logged in!")
}

func findUserByUsername(username string) (User, error) {
	db := database.DBConn
	var dbUser User

	result := db.Table("users").Where("username = ?", username).First(&dbUser)
	if result.Error != nil {
		return dbUser, result.Error
	}

	result.Scan(&dbUser)
	return dbUser, result.Error
}

func encryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), err
}

func verifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
