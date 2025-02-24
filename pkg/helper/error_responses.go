package helper

import (
	"github.com/firstudio-lab/JARITMAS-API/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"strings"
)

func ExtractHTTPCodeAndMessage(err error) (int, string) {
	if err == nil {
		return http.StatusOK, ""
	}
	// Misalkan format error adalah "%d:%e"
	parts := strings.Split(err.Error(), ":")
	if len(parts) != 2 {
		// Jika format error tidak sesuai
		logger.Log.Debug(err)
		return http.StatusInternalServerError, "unexpected error format"
	}
	// Parsing status code
	httpCode, parseErr := strconv.Atoi(parts[0])
	if parseErr != nil {
		return http.StatusInternalServerError, "failed to parse HTTP status code"
	}
	// Error message
	errorMessage := parts[1]
	return httpCode, errorMessage
}

type UseData struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type NoData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func WResponses(ctx *fiber.Ctx, err error, message string, data interface{}) error {
	code, msgErr := ExtractHTTPCodeAndMessage(err)
	if msgErr == "" || err == nil {
		if message == "" {
			ctx.Status(code).JSON(NoData{Status: "OK", Message: message})
		}
		return ctx.Status(code).JSON(UseData{Status: "OK", Message: message, Data: data})
	}
	return ctx.Status(code).JSON(NoData{
		Status:  "ERROR",
		Message: msgErr,
	})
}
