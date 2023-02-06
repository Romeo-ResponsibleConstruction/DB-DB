package controllers

import (
	"DB-DB/database"
	"DB-DB/models"
	"github.com/gofiber/fiber/v2"
)

func GetDashboard(c *fiber.Ctx) error {
	var tickets []models.DeliveryTicket

	database.DB.Find(&tickets)

	return c.JSON(tickets)
}
