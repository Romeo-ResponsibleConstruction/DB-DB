package controllers

import (
	"DB-DB/database"
	"DB-DB/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func GetDashboard(c *fiber.Ctx) error {
	var tickets []models.DeliveryTicket

	database.DB.Find(&tickets)

	return c.JSON(tickets)
}

func AddTicket(c *fiber.Ctx) error {
	// Add a new ticket entry to the database
	// format of ticket not finalised, format of request is draft final

	c.Accepts("application/json")

	// TODO assign id programatically, rather than requiring it in the json input!

	// parse json
	ticket := new(models.DeliveryTicket)
	if err := c.BodyParser(&ticket); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// store ticket to database
	result := database.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&ticket)
	// nothing affected => conflict => show error
	if result.RowsAffected == 0 {
		c.Status(fiber.StatusForbidden)
		return c.JSON(fiber.Map{
			"message": "could not add ticket",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
