package main

import (
	"log"
	"simple-key-value/api/routes"
	"simple-key-value/configuration"
	"simple-key-value/pkg/keyvalue"
)

func main() {

	appInit := configuration.NewAppInit()
	config := configuration.New()

	collection := appInit.InitMongoDB(config)
	redisCache := appInit.InitRedis(config)

	keyValueRepository := keyvalue.NewRepo(collection)
	keyValueService := keyvalue.NewService(keyValueRepository, redisCache)

	app := appInit.InitFiberApp()
	api := app.Group("/api")
	routes.KeyValueRouter(api, keyValueService)

	log.Fatal(app.Listen(":9090"))
}
