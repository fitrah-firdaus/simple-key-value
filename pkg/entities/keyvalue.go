package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type KeyValue struct {
	ID        primitive.ObjectID `json:"id" bson:"_id, omitempty"`
	Key       string             `json:"key" bson:"key"`
	Value     string             `json:"value" bson:"value"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}
