package keyvalue

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"simple-key-value/pkg/entities"
	"time"
)

type Repository interface {
	CreateKey(value *entities.KeyValue) (*entities.KeyValue, error)
	GetKey(key string) (*entities.KeyValue, error)
	UpdateKey(value *entities.KeyValue) (*entities.KeyValue, error)
	DeleteKey(key string) error
}

type repository struct {
	Collection *mongo.Collection
}

func (r repository) CreateKey(value *entities.KeyValue) (*entities.KeyValue, error) {
	value.ID = primitive.NewObjectID()
	value.CreatedAt = time.Now()
	value.UpdatedAt = time.Now()
	_, err := r.Collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, err
	}
	return value, nil
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

func (r repository) UpdateKey(value *entities.KeyValue) (*entities.KeyValue, error) {
	value.UpdatedAt = time.Now()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": value.ID}, bson.M{"$set": value})
	if err != nil {
		return nil, err
	}
	return value, nil
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
