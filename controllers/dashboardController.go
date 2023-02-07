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

	// TODO parse weight
	weight := 33 // test value

	// add ticket
	ticket := models.DeliveryTicket{Id: 2, Weight: uint(weight)}

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
