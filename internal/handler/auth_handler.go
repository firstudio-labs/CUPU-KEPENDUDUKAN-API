package handler

import (
	"fmt"
	"github.com/firstudio-lab/JARITMAS-API/cfg"
	"github.com/firstudio-lab/JARITMAS-API/internal/dto"
	"github.com/firstudio-lab/JARITMAS-API/internal/usecase"
	"github.com/firstudio-lab/JARITMAS-API/pkg/helper"
	"github.com/firstudio-lab/JARITMAS-API/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

type AuthHandler interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
}

type AuthHandlerImpl struct {
	usecase.AuthUsecase
}

func NewAuthHandlerImpl(authUsecase usecase.AuthUsecase) *AuthHandlerImpl {
	return &AuthHandlerImpl{AuthUsecase: authUsecase}
}

func (h AuthHandlerImpl) Register(ctx *fiber.Ctx) error {
	var body dto.RegisterRequest
	if err := ctx.BodyParser(&body); err != nil {
		logger.Log.Errorf("Fail to parse body %e", err)
		return helper.WebResponses(ctx, http.StatusInternalServerError, "FAILED TO PARSE BODY", nil)
	}

	if err := h.AuthUsecase.Register(ctx.Context(), body); err != nil {
		return helper.WebResponses(ctx, http.StatusBadRequest, fmt.Sprintf("%e", err), nil)
	}

	return helper.WebResponses(ctx, http.StatusCreated, "successfully created new user", nil)
}

func (h AuthHandlerImpl) Login(ctx *fiber.Ctx) error {
	var body dto.LoginRequest
	if err := ctx.BodyParser(&body); err != nil {
		logger.Log.Errorf("Fail to parse body %e", err)
		return helper.WebResponses(ctx, http.StatusInternalServerError, "FAILED TO PARSE BODY", nil)
	}

	if err := h.AuthUsecase.Login(ctx.Context(), body); err != nil {
		return helper.WebResponses(ctx, http.StatusBadRequest, fmt.Sprintf("%s", err.Error()), nil)
	}
	//SETTING GENERATE JWT
	expTime := time.Now().Add(time.Minute * 35) // << KADALUARSA DALAM 35 minute
	claims := cfg.JWTClaim{
		Username: body.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "koriebruh.akaJamal",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenValue, err := tokenAlgo.SignedString([]byte(cfg.JWT_KEY))
	if err != nil {
		logger.Log.Errorf("Failed to generate JWT token: %v", err)
		return helper.WebResponses(ctx, http.StatusInternalServerError, "FAILED TO PARSE BODY", nil)
	}

	return helper.WebResponses(ctx, http.StatusOK, "Login Success", map[string]interface{}{
		"token":    tokenValue,
		"lifetime": expTime,
	})
}
