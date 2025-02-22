package helper

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type webResponseNoData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type webResponseData struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func WebResponses(ctx *fiber.Ctx, statusCode int, message string, data interface{}) error {
	switch statusCode {
	case http.StatusCreated:
		return ctx.Status(statusCode).JSON(webResponseNoData{
			Status:  "CREATED",
			Message: message,
		})

	case http.StatusOK:
		if data == nil {
			return ctx.Status(statusCode).JSON(webResponseNoData{
				Status:  "OK",
				Message: message,
			})
		} else {
			return ctx.Status(statusCode).JSON(webResponseData{
				Status:  "OK",
				Message: message,
				Data:    data,
			})
		}

	case http.StatusBadRequest:
		return ctx.Status(statusCode).JSON(webResponseNoData{
			Status:  "BAD REQUEST",
			Message: message,
		})
	case http.StatusNotFound:
		return ctx.Status(statusCode).JSON(webResponseNoData{
			Status:  "NOT FOUND",
			Message: message,
		})

	default:
		return ctx.Status(statusCode).JSON(webResponseNoData{
			Status:  "UKNOWN",
			Message: message,
		})
	}
}
