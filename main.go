package main

import (
	"DB-DB/database"
	"DB-DB/methods"
	"DB-DB/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	println("Starting database...")

	database.Connect()

	println("Starting API...")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	err := app.Listen(methods.StringFromFile("dashboardaddress.txt"))
	if err != nil {
		println(err.Error()) // simple error handling
		return
	}
}
