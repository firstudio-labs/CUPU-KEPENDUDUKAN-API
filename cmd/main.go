package main

import (
	"fmt"
	"github.com/firstudio-lab/JARITMAS-API/cfg"
	"github.com/firstudio-lab/JARITMAS-API/internal/handler"
	"github.com/firstudio-lab/JARITMAS-API/internal/repository"
	"github.com/firstudio-lab/JARITMAS-API/internal/usecase"
	"github.com/firstudio-lab/JARITMAS-API/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {
	logger.InitLogger()
	app := fiber.New()
	DEBE, _ := cfg.GetPool(cfg.GetConfig())
	validate := validator.New()

	userRepository := repository.NewUserRepositoryImpl(DEBE)
	authUsecase := usecase.NewAuthUsecaseImpl(userRepository, validate)
	authHandler := handler.NewAuthHandlerImpl(authUsecase)

	app.Get("/", func(c *fiber.Ctx) error {
		logger.Log.Debug("GEGEGE")
		return c.SendString("Hello, World!")
	})

	app.Post("/api/login", authHandler.Login)
	app.Post("/api/register", authHandler.Register)

	app.Listen(fmt.Sprintf(":%s", cfg.GetConfig().Server.Port))
}
