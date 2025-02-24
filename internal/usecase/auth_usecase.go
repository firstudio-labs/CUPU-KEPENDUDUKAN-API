package usecase

import (
	"context"
	"fmt"
	"github.com/firstudio-lab/JARITMAS-API/internal/dto"
	"github.com/firstudio-lab/JARITMAS-API/internal/entity"
	"github.com/firstudio-lab/JARITMAS-API/internal/repository"
	"github.com/firstudio-lab/JARITMAS-API/pkg/helper"
	"github.com/firstudio-lab/JARITMAS-API/pkg/logger"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"strings"
)

type AuthUsecase interface {
	Register(ctx context.Context, request dto.RegisterRequest) error
	Login(ctx context.Context, request dto.LoginRequest) error
}

type AuthUsecaseImpl struct {
	repository.UserRepository
	*validator.Validate
}

func NewAuthUsecase(userRepository repository.UserRepository, validate *validator.Validate) *AuthUsecaseImpl {
	return &AuthUsecaseImpl{UserRepository: userRepository, Validate: validate}
}

func (a AuthUsecaseImpl) Register(ctx context.Context, request dto.RegisterRequest) error {
	if err := a.Validate.Struct(&request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errorMessages []string
		for _, validationError := range validationErrors {
			errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' is invalid: %s", validationError.Field(), validationError.Tag()))
		}

		return fmt.Errorf("validation failed: %s", strings.Join(errorMessages, ", "))
	}

	hashedPass, err := helper.ArgonGeneratePassword(request.Password)
	if err != nil {
		logger.Log.Debug("Hashed Pass Error")
	}

	newUser := entity.User{
		Model:    gorm.Model{},
		Username: request.Username,
		Password: hashedPass,
	}
	if err := a.UserRepository.CreateUser(ctx, newUser); err != nil {
		return err
	}

	return nil
}

func (a AuthUsecaseImpl) Login(ctx context.Context, request dto.LoginRequest) error {

	if err := a.Validate.Struct(&request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errorMessages []string
		for _, validationError := range validationErrors {
			errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' is invalid: %s", validationError.Field(), validationError.Tag()))
		}

		return fmt.Errorf("validation failed: %s", strings.Join(errorMessages, ", "))
	}

	user, err := a.UserRepository.FindUserByUsername(ctx, request.Username)
	if err != nil {
		return err
	}

	result, err := helper.ArgonComparePassword(user.Password, request.Password)
	if err != nil {
		logger.Log.Debug("Compare Pass Error", err)
		return err
	}

	if result == false {
		return fmt.Errorf("invalid credentisals")
	}

	return nil
}
