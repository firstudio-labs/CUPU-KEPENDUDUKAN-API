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

	userRepository := repository.NewUserRepository(DEBE)
	authUsecase := usecase.NewAuthUsecase(userRepository, validate)
	authHandler := handler.NewAuthHandler(authUsecase)

	citizensRepository := repository.NewCitizensRepository()
	citizensUsecase := usecase.NewCitizensUsecase(citizensRepository, validate, DEBE)
	citizensHandler := handler.NewCitizensHandler(citizensUsecase)

	jobsRepository := repository.NewJobsRepository(DEBE)
	jobsUsecase := usecase.NewJobsUsecase(validate, jobsRepository)
	jobsHandler := handler.NewJobsHandler(jobsUsecase)

	app.Post("/api/login", authHandler.Login)
	app.Post("/api/register", authHandler.Register)
	protected := app.Group("/", cfg.JWTAuthMiddleware)
	protected.Get("/", func(c *fiber.Ctx) error {
		logger.Log.Debug("GEGEGE")
		return c.SendString("Hello, World!")
	})

	protected.Get("/api/citizens/:nik", citizensHandler.FindCitizenByNIK)
	protected.Get("/api/citizens", citizensHandler.FindCitizenPage)         // unclear
	protected.Post("/api/citizens", citizensHandler.CreateCitizen)          // unclear
	protected.Put("/api/citizens/:nik", citizensHandler.UpdateCitizenByNIK) // unclear
	protected.Delete("/api/citizens/:nik", citizensHandler.DeleteCitizenByNIK)

	protected.Delete("/api/jobs/:id", jobsHandler.DeleteJobById)
	protected.Get("/api/jobs", jobsHandler.GetJobs)
	protected.Post("/api/jobs", jobsHandler.CreateJob)
	protected.Put("/api/jobs/:id", jobsHandler.UpdateJobById)

	app.Listen(fmt.Sprintf(":%s", cfg.GetConfig().Server.Port))
}
