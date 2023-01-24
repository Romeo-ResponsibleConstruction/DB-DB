package controllers

import "github.com/gofiber/fiber/v2"

func GetDashboard(c *fiber.Ctx) error {
	return c.JSON(c.App().Stack())
}
