package presenter

import (
	"github.com/gofiber/fiber/v2"
	"simple-key-value/pkg/entities"
)

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func KeyValueSuccessResponse(data *entities.KeyValue) *fiber.Map {
	keyValue := KeyValue{
		Key:   data.Key,
		Value: data.Value,
	}
	return &fiber.Map{
		"status": true,
		"data":   keyValue,
		"error":  nil,
	}
}

func KeyValueErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
