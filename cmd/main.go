package main

import (
	"go/go-fiber/internal/home"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	home.NewHandler(app)

	app.Listen(":3000")
}
