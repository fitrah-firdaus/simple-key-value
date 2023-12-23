package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"simple-key-value/api/routes"
	"simple-key-value/configuration"
	"simple-key-value/pkg/keyvalue"
)

func main() {

	config := configuration.New()
	database, client := configuration.NewMongoDatabase(config)

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	collection := database.Collection(config.Get("MONGO_COLLECTION"))
	keyValueRepository := keyvalue.NewRepo(collection)
	keyValueService := keyvalue.NewService(keyValueRepository)

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Health Check Success"))
	})
	api := app.Group("/api")
	routes.KeyValueRouter(api, keyValueService)
	log.Fatal(app.Listen(":9090"))
}
