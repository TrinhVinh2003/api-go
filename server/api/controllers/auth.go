package controllers

import (
	"project/vnexpress/api/models"
	"project/vnexpress/config/driver/database"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "vinhcute01"

func Signup(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"message": "Don't sign up",
		})
	}

	var existingUser models.User
	database.DBConn.Where("email = ?", data["email"]).First(&existingUser)
	if existingUser.Id != 0 {
		c.JSON(fiber.Map{
			"message": "user already exist",
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	users := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}
	database.DBConn.Create(&users)
	return c.JSON(users)

}
func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.JSON(fiber.Map{
			"message": "Please enter login",
		})

	}
	if data["email"] == "" || data["password"] == "" {
		c.JSON(fiber.Map{
			"message": "Password or email empty",
		})
	}
	var user models.User

	database.DBConn.Where("email = ?", data["email"]).First(&user)
	if user.Id == 0 {
		c.JSON(fiber.Map{
			"message": "Not found user",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	claim := jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claim)

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		c.JSON(fiber.Map{
			"message": "Could not generate token",
		})
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
	})

}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Logout success",
	})
}
