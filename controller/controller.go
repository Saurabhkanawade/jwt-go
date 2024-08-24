package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"saurabhkanawade/jwt/database"
	"saurabhkanawade/jwt/models"
	"strconv"
	"time"
)

func User(c *fiber.Ctx) error {
	log.Info("User called...")
	cookies := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookies, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("id = ? ", claims.Issuer).First(&user)

	return c.JSON(models.UserModel{
		Name:  user.Name,
		Email: user.Email,
	})
}

func Register(c *fiber.Ctx) error {
	log.Info("Register called...")

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := &models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	database.DB.Create(&user)

	return c.JSON(user)
}

const SecretKey = "secret"

func Login(c *fiber.Ctx) error {
	log.Info("Login called...")
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found, please register new !",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Clould not login",
		})
	}

	cookies := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 12),
		HTTPOnly: true,
	}

	c.Cookie(&cookies)

	return c.JSON(fiber.Map{
		"message": "successfully login to application",
	})
}

func Logout(c *fiber.Ctx) error {
	log.Info("Logout called...")

	cookies := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookies)

	return c.JSON(fiber.Map{
		"message": "Successfully logged out from application",
	})
}
