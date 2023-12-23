package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"net/http"
	"simple-key-value/api/presenter"
	"simple-key-value/pkg/entities"
	"simple-key-value/pkg/keyvalue"
)

func CreateKey(service keyvalue.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var requestBody entities.KeyValue
		err := ctx.BodyParser(&requestBody)
		log.Infow("input=", requestBody)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.KeyValueErrorResponse(err))
		}
		result, err := service.CreateKey(&requestBody)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.KeyValueErrorResponse(err))
		}
		return ctx.JSON(presenter.KeyValueSuccessResponse(result))
	}
}
