package routes

import (
	"DB-DB/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.GetDashboard)
	app.Get("/tickets", controllers.GetTickets)
	app.Get("/img/:fp", controllers.GetPicture)
	app.Get("/failedchecks/:id/:quantity", controllers.GetFailedChecks)
	app.Post("/newTicket", controllers.AddTicket)
	app.Delete("/deleteTicket/:id", controllers.DeleteTicket)
}
