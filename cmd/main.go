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
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	logger.InitLogger()
	app := fiber.New()
	DEBE, _ := cfg.GetPool(cfg.GetConfig())
	validate := validator.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: false,
	}))

	userRepository := repository.NewUserRepositoryImpl(DEBE)
	authUsecase := usecase.NewAuthUsecaseImpl(userRepository, validate)
	authHandler := handler.NewAuthHandlerImpl(authUsecase)

	app.Post("/api/login", authHandler.Login)
	app.Post("/api/register", authHandler.Register)
	protected := app.Group("/", cfg.JWTAuthMiddleware)
	protected.Get("/", func(c *fiber.Ctx) error {
		logger.Log.Debug("GEGEGE")
		return c.SendString("Hello, World!")
	})

	app.Listen(fmt.Sprintf(":%s", cfg.GetConfig().Server.Port))
}
