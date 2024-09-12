package main

import (
	"github.com/AthThobari/learn-api-go-divrhino-trivia/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFact)

	app.Post("/fact", handlers.CreateFact)
}