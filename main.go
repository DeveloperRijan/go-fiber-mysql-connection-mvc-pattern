package main

import (
	"GoFiberMySQLApp/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(context *fiber.Ctx) error {
		//init configs
		msg := config.InitConfigs()
		return context.SendString(msg)
	})

	app.Listen(":8080")
}
