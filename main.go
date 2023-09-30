package main

import (
	"belajar-backend/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })
	database.Connect()





    log.Fatal(app.Listen(":3000"))
}