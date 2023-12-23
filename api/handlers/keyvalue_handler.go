package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"net/http"
	"simple-key-value/api/presenter"
	"simple-key-value/pkg/entities"
	"simple-key-value/pkg/keyvalue"
)

func CreateOrUpdateKey(service keyvalue.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var requestBody entities.KeyValue
		err := ctx.BodyParser(&requestBody)
		log.Infow("input=", requestBody)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.KeyValueErrorResponse(err))
		}
		result, err := service.CreateOrUpdateKey(&requestBody)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.KeyValueErrorResponse(err))
		}
		return ctx.JSON(presenter.KeyValueSuccessResponse(result))
	}
}

func GetKey(service keyvalue.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		key := ctx.Params("key")
		result, err := service.GetKey(key)
		if err != nil {
			log.Error(err)
		}
		if result != nil {
			return ctx.JSON(presenter.KeyValueSuccessResponse(result))
		}
		ctx.Status(http.StatusNotFound)
		return nil
	}
}

func DeleteKey(service keyvalue.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		key := ctx.Params("key")
		err := service.DeleteKey(key)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.KeyValueErrorResponse(err))
		}
		return ctx.JSON(&fiber.Map{
			"status": true,
			"data":   "Deleted successfully",
			"err":    nil,
		})
	}
}
