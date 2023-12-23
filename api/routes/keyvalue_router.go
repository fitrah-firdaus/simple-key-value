package routes

import (
	"github.com/gofiber/fiber/v2"
	"simple-key-value/api/handlers"
	"simple-key-value/pkg/keyvalue"
)

func KeyValueRouter(app fiber.Router, service keyvalue.Service) {
	app.Post("/kv", handlers.CreateKey(service))
}
