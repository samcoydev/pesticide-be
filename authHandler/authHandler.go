package authhandler

import (
	"pesticide/database"
	log "pesticide/logHandler"
	models "pesticide/models/user"

	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
)

var fromName string = "[authHandler.go]"

func Register(ctx *fiber.Ctx) {
	db := database.DBConn
	user := new(models.User)

	// Create a user object from the posted data in "ctx"
	if err := ctx.BodyParser(user); err != nil {
		log.Debug(fromName, "Error parsing")
		ctx.Status(401).Send(err)
		return
	}

	encryptedPassword, err := encryptPassword(user.Password)
	if err != nil {
		log.Debug(fromName, "Error encrypting your password")
		ctx.Status(401).Send(err)
		return
	}

	user.Password = encryptedPassword

	db.Create(&user)
}

func Authenticate(ctx *fiber.Ctx) {
	user := new(models.User)

	// Create a user object from the posted data in "ctx"
	if err := ctx.BodyParser(user); err != nil {
		ctx.Status(401).Send(err)
		return
	}

	dbUser, err := findUserByUsername(user.Username)
	if err != nil {
		log.Debug(fromName, "Cannot find user: "+user.Username)
		ctx.Status(401).Send(err)
		return
	}

	if err := verifyPassword(dbUser.Password, user.Password); err != nil {
		log.Debug(fromName, "Passwords dont match for user: "+user.Username)
		ctx.Status(401).Send(err)
		return
	}

	log.Debug(fromName, "User logged in!")
	ctx.JSON(user)
}

func findUserByUsername(username string) (models.User, error) {
	db := database.DBConn
	var dbUser models.User

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
