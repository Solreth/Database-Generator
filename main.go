package main

import (
	"Basic-Backend/database"
	"log"

	// "os"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// os.Setenv("CGO_ENABLED", "0")
	database.ConnectDb()
	app := fiber.New()
	app.Mount("/", Routes())
	log.Fatal(app.Listen(":3000"))
}
