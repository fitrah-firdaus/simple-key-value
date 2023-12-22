package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Health Check Success"))
	})
	log.Fatal(app.Listen(":9090"))
}
