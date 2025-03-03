package main

import (
	"fmt"
	"github.com/firstudio-lab/JARITMAS-API/cfg"
	"github.com/firstudio-lab/JARITMAS-API/internal/handler"
	"github.com/firstudio-lab/JARITMAS-API/internal/repository"
	"github.com/firstudio-lab/JARITMAS-API/internal/usecase"
	"github.com/firstudio-lab/JARITMAS-API/pkg/helper"
	"github.com/firstudio-lab/JARITMAS-API/pkg/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

// Middleware to validate the API key
func Middleware(APIKEY string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Getting API key from the request header
		key := c.GetHeader("X-API-Key")
		if key != APIKEY {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.NoData{
				Status:  "ERROR",
				Message: "Invalid API key",
			})
			return
		}
		c.Next()
	}
}

func main() {
	logger.InitLogger()
	r := gin.Default()

	DEBE, _ := cfg.GetPool(cfg.GetConfig()) // Assuming error handling is done within the function
	validate := validator.New()
	APIKEY := "KORIE"

	//r.Use(gin.Recovery())
	// CORS middleware setup for Gin
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true, // or use AllowOrigins to specify exact origins
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-API-Key"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(Middleware(APIKEY))

	citizensRepository := repository.NewCitizensRepository()
	citizensUsecase := usecase.NewCitizensUsecase(citizensRepository, validate, DEBE)
	citizensHandler := handler.NewCitizensHandler(citizensUsecase)

	jobsRepository := repository.NewJobsRepository(DEBE)
	jobsUsecase := usecase.NewJobsUsecase(validate, jobsRepository)
	jobsHandler := handler.NewJobsHandler(jobsUsecase)

	// Setup routes for Citizens
	r.GET("/api/citizens/:nik", citizensHandler.FindCitizenByNIK)
	r.GET("/api/citizens", citizensHandler.FindCitizenPage) // Assuming this is implemented correctly in your handler
	r.POST("/api/citizens", citizensHandler.CreateCitizen)
	r.PUT("/api/citizens/:nik", citizensHandler.UpdateCitizenByNIK)
	r.DELETE("/api/citizens/:nik", citizensHandler.DeleteCitizenByNIK)

	// Setup routes for Jobs
	r.DELETE("/api/jobs/:id", jobsHandler.DeleteJobById)
	r.GET("/api/jobs", jobsHandler.GetJobs)
	r.POST("/api/jobs", jobsHandler.CreateJob)
	r.PUT("/api/jobs/:id", jobsHandler.UpdateJobById)

	// Start the Gin server
	port := cfg.GetConfig().Server.Port
	if port == "" {
		port = "3000"
	}
	r.Run(fmt.Sprintf(":%s", port))
}
