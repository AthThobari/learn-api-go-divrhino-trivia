package main

import (
	"github.com/AthThobari/learn-api-go-divrhino-trivia/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)



	app.Listen(":3000")
}
