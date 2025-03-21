package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/firstudio-lab/JARITMAS-API/cfg"
	"github.com/firstudio-lab/JARITMAS-API/internal/handler"
	"github.com/firstudio-lab/JARITMAS-API/internal/repository"
	"github.com/firstudio-lab/JARITMAS-API/internal/usecase"
	"github.com/firstudio-lab/JARITMAS-API/pkg/logger"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

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
			c.AbortWithStatus(http.StatusNoContent)
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

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(CustomCORSMiddleware())
	r.Use(APIKeyMiddleware(APIKEY))

	citizensRepository := repository.NewCitizensRepository()
	citizensUsecase := usecase.NewCitizensUsecase(citizensRepository, validate, DEBE)
	citizensHandler := handler.NewCitizensHandler(citizensUsecase)

	jobsRepository := repository.NewJobsRepository(DEBE)
	jobsUsecase := usecase.NewJobsUsecase(validate, jobsRepository)
	jobsHandler := handler.NewJobsHandler(jobsUsecase)

	countryHandler := handler.NewCountryHandler(DEBE)
	relocationHandler := handler.NewRelocationHandler(validate, DEBE)

	// Setup routes for Citizens
	v1 := r.Group("/api")
	{
		v1.GET("/citizens/:nik", citizensHandler.FindCitizenByNIK)
		v1.GET("/citizens", citizensHandler.FindCitizenPage) // ?page
		v1.POST("/citizens", citizensHandler.CreateCitizen)
		v1.PUT("/citizens/:nik", citizensHandler.UpdateCitizenByNIK)
		v1.DELETE("/citizens/:nik", citizensHandler.DeleteCitizenByNIK)
		v1.GET("/citizens-family/:kk", citizensHandler.FindAllMemberByKK)
		v1.GET("/all-citizens", citizensHandler.FindAllCitizens)
		v1.GET("/citizens-search/:namePattern", citizensHandler.FindSimilarName)

		// Setup routes for Jobs
		v1.DELETE("/jobs/:id", jobsHandler.DeleteJobById)
		v1.GET("/jobs", jobsHandler.GetJobs)
		v1.POST("/jobs", jobsHandler.CreateJob)
		v1.PUT("/jobs/:id", jobsHandler.UpdateJobById)
		v1.GET("/jobs/:id", jobsHandler.GetJobById)
		v1.GET("/jobs-search/:namePattern", jobsHandler.GetSimilarJobsName)

		//
		v1.GET("/provinces", countryHandler.GetProvince)
		v1.GET("/districts/:province-code", countryHandler.GetDistrictByProvinceCode)
		v1.GET("/sub-districts/:district-code", countryHandler.GetSubDistrictByDistrictCode)
		v1.GET("/villages/:sub-district-code", countryHandler.GetVillageBySUbDistrictCode)

		v1.GET("/province", countryHandler.ProvincesPagination)
		v1.GET("/districts", countryHandler.DistrictsPagination)
		v1.GET("/sub-districts", countryHandler.SubDistrictsPagination)
		v1.GET("/villages", countryHandler.VillagesPagination)

		v1.GET("/relocations", relocationHandler.GetPerPage)
		v1.POST("/relocations", relocationHandler.AddRelocation)
		v1.PUT("/relocations/:id", relocationHandler.UpdateRelocation)
		v1.DELETE("/relocations/:id", relocationHandler.DeleteRelocation)
		v1.POST("/relocations/:id/approved", relocationHandler.ApproveRelocation) // SAMSEK

	}

	// START SERVER
	port := cfg.GetConfig().Server.Port
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r.Handler(),
	}

	go func() {
		slog.Info("Listening And Server HTTP on ", slog.String("port", port))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("Server failed: %s", err)
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("Server shutdown gracefully")
	}
}
