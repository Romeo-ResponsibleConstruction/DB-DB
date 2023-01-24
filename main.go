package main

import (
	"DB-DB/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	println("Hello, World!")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	err := app.Listen("127.0.0.1:8000")
	if err != nil {
		println(err) // simple error handling
		return
	}
}
