package configuration

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewMongoDatabase(config Config) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(config.Get("MONGO_URI")))
	/*defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()*/
	db := client.Database(config.Get("MONGO_DATABASE"))
	return db
}
