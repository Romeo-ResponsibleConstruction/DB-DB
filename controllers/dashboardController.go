package controllers

import "github.com/gofiber/fiber/v2"

func GetDashboard(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"hello":  "world",
		"server": c.Hostname(),
	})
}
