package helper

import (
	"fmt"
	"github.com/firstudio-lab/JARITMAS-API/pkg/logger"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

// ExtractHTTPCodeAndMessage parses an error formatted as "<code>:<message>" and returns the code and message
func ExtractHTTPCodeAndMessage(err error) (int, string) {
	// Misalkan format error adalah "%d:%e"
	parts := strings.Split(err.Error(), ":")
	if len(parts) != 2 {
		// Jika format error tidak sesuai, log error dan kembalikan status BadRequest
		logger.Log.Debug(err)
		return http.StatusBadRequest, fmt.Sprintf("%s", err)
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

func ErrResponses(ctx *gin.Context, err error) {
	code, msgErr := ExtractHTTPCodeAndMessage(err)
	ctx.JSON(code, NoData{
		Status:  "ERROR",
		Message: msgErr,
	})
	fmt.Println(code, msgErr)
	slog.Info("RESULT => ", slog.Int("CODE", code), slog.String("message", msgErr))
}
