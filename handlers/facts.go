package handlers

import (
	"github.com/AthThobari/learn-api-go-divrhino-trivia/database"
	"github.com/AthThobari/learn-api-go-divrhino-trivia/models"
	"github.com/gofiber/fiber/v2"
)

func ListFact(c *fiber.Ctx) error{
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)
	return c.Status(200).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)

	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)
	return c.Status(200).JSON(fact)
}