package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type KeyValue struct {
	ID        primitive.ObjectID `json:"id" bson:"_id, omitempty" gorm:"-:all"`
	Key       string             `json:"key" bson:"key" gorm:"primaryKey;column:keylog;type:varchar(255)"`
	Value     string             `json:"value" bson:"value" gorm:"type:text"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}
