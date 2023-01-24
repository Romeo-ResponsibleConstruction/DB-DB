package routes

import (
	"DB-DB/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.GetDashboard)
}
