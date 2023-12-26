package configuration

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type appInit struct {
}

type AppInit interface {
	InitMySQL(config Config) *sql.DB
	InitGormMySQL(config Config) *gorm.DB
	InitMongoDB(config Config) *mongo.Collection
	InitRedis(config Config) RedisCache
	InitFiberApp() *fiber.App
}

func (a *appInit) InitMySQL(config Config) *sql.DB {
	db := NewMySQLDatabase(config)
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"github://fitrah-firdaus/simple-key-value/database/migration",
		"mysql", driver)
	if err != nil {
		log.Error(err)
	}
	log.Info(driver)
	log.Info(m)
	err = m.Up()
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return db
}

func (a *appInit) InitGormMySQL(config Config) *gorm.DB {
	return NewGormMySQL(config)
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
