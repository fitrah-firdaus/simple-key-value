package main

import (
	"github.com/fitrah-firdaus/simple-key-value/api/routes"
	"github.com/fitrah-firdaus/simple-key-value/configuration"
	"github.com/fitrah-firdaus/simple-key-value/pkg/keyvalue"
	"github.com/fitrah-firdaus/simple-key-value/pkg/keyvalue/repository"
	"log"
)

func main() {

	appInit := configuration.NewAppInit()
	config := configuration.New()

	//collection := appInit.InitMongoDB(config)
	db := appInit.InitMySQL(config)
	redisCache := appInit.InitRedis(config)

	keyValueRepository := repository.NewMySQLRepo(db)
	keyValueService := keyvalue.NewService(keyValueRepository, redisCache)

	app := appInit.InitFiberApp()
	api := app.Group("/api")
	routes.KeyValueRouter(api, keyValueService)

	log.Fatal(app.Listen(":9090"))
}
