package routes

import (
	"github.com/fitrah-firdaus/simple-key-value/api/handlers"
	"github.com/fitrah-firdaus/simple-key-value/pkg/keyvalue"
	"github.com/gofiber/fiber/v2"
)

func KeyValueRouter(app fiber.Router, service keyvalue.Service) {
	app.Post("/kv", handlers.CreateOrUpdateKey(service))
	app.Get("/kv/:key", handlers.GetKey(service))
	app.Delete("/kv/:key", handlers.DeleteKey(service))
}
