package repository

import (
	"context"
	"github.com/fitrah-firdaus/simple-key-value/pkg/entities"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Repository interface {
	CreateOrUpdateKey(value *entities.KeyValue) (*entities.KeyValue, error)
	GetKey(key string) (*entities.KeyValue, error)
	DeleteKey(key string) error
}

type repository struct {
	Collection *mongo.Collection
}

func (r repository) CreateOrUpdateKey(value *entities.KeyValue) (*entities.KeyValue, error) {
	value.ID = primitive.NewObjectID()
	value.CreatedAt = time.Now()
	value.UpdatedAt = time.Now()
	log.Info("value = %s", value)
	filter := bson.M{"key": value.Key}
	update := bson.M{
		"$set": bson.M{
			"key":       value.Key,
			"value":     value.Value,
			"updatedAt": value.UpdatedAt,
		},
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	result := r.Collection.FindOneAndUpdate(context.Background(), filter, update, &opt)
	if result.Err() != nil {
		log.Error(result.Err())
		return nil, result.Err()
	}
	decodeErr := result.Decode(&value)
	return value, decodeErr
}

func (r repository) GetKey(key string) (*entities.KeyValue, error) {
	var result entities.KeyValue
	filter := bson.D{{"key", key}}
	err := r.Collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r repository) DeleteKey(key string) error {
	_, err := r.Collection.DeleteOne(context.Background(), bson.M{"key": key})
	if err != nil {
		return err
	}
	return nil
}

func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}
