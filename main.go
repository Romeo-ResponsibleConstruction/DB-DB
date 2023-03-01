package main

import (
	"DB-DB/database"
	"DB-DB/routes"
	"bytes"
	"io"
	"os"
	"strings"

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

	//get the dashboard address from file
	fi, err := os.Open("dashboardaddress.txt")
	if err != nil {
		panic(err)
	}
	// read to buffer
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
	}
	// buffer -> string
	buf = bytes.Trim(buf, "\x00")
	dashboardAddress := strings.TrimSpace(string(buf))
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	err = app.Listen(dashboardAddress)
	if err != nil {
		println(err) // simple error handling
		return
	}
}
