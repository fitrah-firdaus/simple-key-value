package configuration

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		panic(err)
	}
	return db
}
