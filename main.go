package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"simple-key-value/configuration"
	"simple-key-value/pkg/keyvalue"
)

func main() {

	config := configuration.New()
	database := configuration.NewMongoDatabase(config)
	collection := database.Collection(config.Get("MONGO_COLLECTION"))
	keyValueRepository := keyvalue.NewRepo(collection)
	keyvalue.NewService(keyValueRepository)

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Health Check Success"))
	})
	log.Fatal(app.Listen(":9090"))
}
