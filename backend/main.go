package main

import (
	"github.com/gofiber/fiber/v3"
	"new_dashboard_data/api"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Get("/weather", func(c fiber.Ctx) error {
		var weather = api.GetWeather()
		return c.JSON(weather)
	})
	app.Listen(":3023")
}
