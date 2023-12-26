package configuration

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func NewMongoDatabase(config Config) (*mongo.Database, *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(config.Get("MONGO_URI")))
	/*defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()*/
	db := client.Database(config.Get("MONGO_DATABASE"))
	return db, client
}

func NewMySQLDatabase(config Config) *sql.DB {
	db, err := sql.Open("mysql", config.Get("MYSQL_URI"))
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return db
}

func NewGormMySQL(config Config) *gorm.DB {
	sqlDB := NewMySQLDatabase(config)
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return gormDB
}
