package main

import (
	"go/go-fiber/config"
	"go/go-fiber/internal/home"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()

	app := fiber.New()
	log.SetLevel(log.Level(logConfig.Level))
	app.Use(recover.New())

	home.NewHandler(app)

	app.Listen(":3000")
}
