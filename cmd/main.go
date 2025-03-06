package main

import (
	"fmt"
	"github.com/firstudio-lab/JARITMAS-API/cfg"
	"github.com/firstudio-lab/JARITMAS-API/internal/handler"
	"github.com/firstudio-lab/JARITMAS-API/internal/repository"
	"github.com/firstudio-lab/JARITMAS-API/internal/usecase"
	"github.com/firstudio-lab/JARITMAS-API/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func CustomCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("CustomCORSMiddleware called")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-API-Key")
		if c.Request.Method == "OPTIONS" {
			fmt.Println("Handling OPTIONS method")
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	}
}

func APIKeyMiddleware(APIKEY string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("APIKeyMiddleware called")
		if c.Request.Method == "OPTIONS" {
			fmt.Println("Skipping API key check for OPTIONS method")
			c.Next()
			return
		}
		key := c.GetHeader("X-API-Key")
		if key != APIKEY {
			fmt.Println("Invalid API key provided")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "ERROR",
				"message": "Invalid API key",
			})
			return
		}
		c.Next()
	}
}

func main() {
	logger.InitLogger()
	DEBE, _ := cfg.GetPool(cfg.GetConfig())
	validate := validator.New()
	APIKEY := "KORIE"

	r := gin.Default()
	r.Use(CustomCORSMiddleware())
	r.Use(APIKeyMiddleware(APIKEY))

	citizensRepository := repository.NewCitizensRepository()
	citizensUsecase := usecase.NewCitizensUsecase(citizensRepository, validate, DEBE)
	citizensHandler := handler.NewCitizensHandler(citizensUsecase)

	jobsRepository := repository.NewJobsRepository(DEBE)
	jobsUsecase := usecase.NewJobsUsecase(validate, jobsRepository)
	jobsHandler := handler.NewJobsHandler(jobsUsecase)

	// Setup routes for Citizens
	r.GET("/api/citizens/:nik", citizensHandler.FindCitizenByNIK)
	r.GET("/api/citizens", citizensHandler.FindCitizenPage) // ?page
	r.POST("/api/citizens", citizensHandler.CreateCitizen)
	r.PUT("/api/citizens/:nik", citizensHandler.UpdateCitizenByNIK)
	r.DELETE("/api/citizens/:nik", citizensHandler.DeleteCitizenByNIK)
	r.GET("/api/citizens-family/:kk", citizensHandler.FindAllMemberByKK)
	r.GET("/api/all-citizens", citizensHandler.FindAllCitizens)
	r.GET("/api/citizens-search/:namePattern", citizensHandler.FindSimilarName)

	// Setup routes for Jobs
	r.DELETE("/api/jobs/:id", jobsHandler.DeleteJobById)
	r.GET("/api/jobs", jobsHandler.GetJobs)
	r.POST("/api/jobs", jobsHandler.CreateJob)
	r.PUT("/api/jobs/:id", jobsHandler.UpdateJobById)
	r.GET("/api/jobs/:id", jobsHandler.GetJobById)
	r.GET("/api/jobs-search/:namePattern", jobsHandler.GetSimilarJobsName)

	// Start the Gin server
	port := cfg.GetConfig().Server.Port
	if port == "" {
		port = "3000"
	}
	r.Run(fmt.Sprintf(":%s", port))
}
