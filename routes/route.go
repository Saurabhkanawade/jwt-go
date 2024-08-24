package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"saurabhkanawade/jwt/controller"
)

func Routes(c *fiber.App) {
	log.Info("getting routes...")
	c.Get("/getUser", controller.User)
	c.Post("/register", controller.Register)
	c.Post("/login", controller.Login)
	c.Post("/logout", controller.Logout)
}
