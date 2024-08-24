package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"saurabhkanawade/jwt/database"
	"saurabhkanawade/jwt/routes"
)

func main() {
	fmt.Println("Welcome to jwt auth...........")
	database.DBConn()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: false,
	}))

	routes.Routes(app)

	err := app.Listen(":9090")
	if err != nil {
		log.Fatalf("Error found while starting the server %s", err)
	}
}
