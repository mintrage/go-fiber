package main

import (
	"go/go-fiber/config"
	"go/go-fiber/internal/home"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func main() {
	config.Init()
	dbConf := config.NewDatabaseConfig()
	log.Println(dbConf)

	app := fiber.New()
	app.Use(recover.New())

	home.NewHandler(app)

	app.Listen(":3000")
}
