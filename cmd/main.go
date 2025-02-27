package main

import (
	"fmt"
	"github.com/firstudio-lab/JARITMAS-API/cfg"
	"github.com/firstudio-lab/JARITMAS-API/internal/handler"
	"github.com/firstudio-lab/JARITMAS-API/internal/repository"
	"github.com/firstudio-lab/JARITMAS-API/internal/usecase"
	"github.com/firstudio-lab/JARITMAS-API/pkg/helper"
	"github.com/firstudio-lab/JARITMAS-API/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Middleware(APIKEY string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Mendapatkan API key dari header request
		key := c.Get("X-API-Key")
		if key != APIKEY {
			return c.Status(fiber.StatusUnauthorized).JSON(helper.NoData{Status: "ERROR", Message: "Invalid API key"})
		}
		return c.Next()
	}
}

func main() {
	logger.InitLogger()
	app := fiber.New()
	DEBE, _ := cfg.GetPool(cfg.GetConfig())
	validate := validator.New()
	APIKEY := "KORIE"

	app.Use(Middleware(APIKEY))

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: false,
	}))

	citizensRepository := repository.NewCitizensRepository()
	citizensUsecase := usecase.NewCitizensUsecase(citizensRepository, validate, DEBE)
	citizensHandler := handler.NewCitizensHandler(citizensUsecase)

	jobsRepository := repository.NewJobsRepository(DEBE)
	jobsUsecase := usecase.NewJobsUsecase(validate, jobsRepository)
	jobsHandler := handler.NewJobsHandler(jobsUsecase)

	app.Get("/api/citizens/:nik", citizensHandler.FindCitizenByNIK)
	app.Get("/api/citizens", citizensHandler.FindCitizenPage)         // unclear
	app.Post("/api/citizens", citizensHandler.CreateCitizen)          // unclear
	app.Put("/api/citizens/:nik", citizensHandler.UpdateCitizenByNIK) // unclear
	app.Delete("/api/citizens/:nik", citizensHandler.DeleteCitizenByNIK)

	app.Delete("/api/jobs/:id", jobsHandler.DeleteJobById)
	app.Get("/api/jobs", jobsHandler.GetJobs)
	app.Post("/api/jobs", jobsHandler.CreateJob)
	app.Put("/api/jobs/:id", jobsHandler.UpdateJobById)

	app.Listen(fmt.Sprintf(":%s", cfg.GetConfig().Server.Port))
}
