package configuration

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
)

type appInit struct {
}

type AppInit interface {
	InitMySQL(config Config) *sql.DB
	InitMongoDB(config Config) *mongo.Collection
	InitRedis(config Config) RedisCache
	InitFiberApp() *fiber.App
}

func (a *appInit) InitMySQL(config Config) *sql.DB {
	return NewMySQLDatabase(config)
}

func (a *appInit) InitMongoDB(config Config) *mongo.Collection {
	database, _ := NewMongoDatabase(config)

	collection := database.Collection(config.Get("MONGO_COLLECTION"))
	return collection
}

func (a *appInit) InitRedis(config Config) RedisCache {
	return NewRedisCache(config)
}

func (a *appInit) InitFiberApp() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Health Check Success"))
	})
	return app
}

func NewAppInit() AppInit {
	return &appInit{}
}
