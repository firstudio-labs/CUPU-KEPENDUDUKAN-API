package handler

import (
	"fmt"
	"github.com/firstudio-lab/JARITMAS-API/internal/dto"
	"github.com/firstudio-lab/JARITMAS-API/internal/usecase"
	"github.com/firstudio-lab/JARITMAS-API/pkg/helper"
	"github.com/firstudio-lab/JARITMAS-API/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type JobsHandler interface {
	CreateJob(c *gin.Context)
	UpdateJobById(c *gin.Context)
	DeleteJobById(c *gin.Context)
	GetJobs(c *gin.Context)
	GetJobById(c *gin.Context)
	GetSimilarJobsName(c *gin.Context)
}

type JobsHandlerImpl struct {
	usecase.JobsUsecase
}

func NewJobsHandler(jobsUsecase usecase.JobsUsecase) *JobsHandlerImpl {
	return &JobsHandlerImpl{JobsUsecase: jobsUsecase}
}

func (h JobsHandlerImpl) CreateJob(c *gin.Context) {
	var body dto.JobReqCreate
	if err := c.ShouldBindJSON(&body); err != nil {
		logger.Log.Errorf("Fail to parse body: %v", err)
		err = fmt.Errorf("%d:%s", http.StatusBadRequest, "failed to parse json")
		helper.ErrResponses(c, err)
		return
	}

	err := h.JobsUsecase.CreateJobs(c.Request.Context(), body)
	if err != nil {
		helper.ErrResponses(c, err)
		return
	}

	c.JSON(http.StatusCreated, helper.NoData{
		Status:  "CREATED",
		Message: "Successfully created job data",
	})
}
func (h JobsHandlerImpl) UpdateJobById(c *gin.Context) {
	id := c.Param("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.NoData{
			Status:  "error",
			Message: "ID is incorrect",
		})
		return
	}

	var body dto.JobReqUpdate
	if err := c.ShouldBindJSON(&body); err != nil {
		logger.Log.Errorf("Fail to parse body: %v", err)
		err = fmt.Errorf("%d:%s", http.StatusBadRequest, "failed to parse json")
		helper.ErrResponses(c, err)
		return
	}

	err = h.JobsUsecase.UpdateJobs(c.Request.Context(), atoi, body)
	if err != nil {
		helper.ErrResponses(c, err)
		return
	}

	c.JSON(http.StatusOK, helper.NoData{
		Status:  "OK",
		Message: fmt.Sprintf("Successfully updated job ID %d", atoi),
	})
}

func (h JobsHandlerImpl) DeleteJobById(c *gin.Context) {
	id := c.Param("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.NoData{
			Status:  "error",
			Message: "ID is not suitable",
		})
		return
	}

	err = h.JobsUsecase.DeleteJobs(c.Request.Context(), atoi)
	if err != nil {
		helper.ErrResponses(c, err)
		return
	}

	c.JSON(http.StatusOK, helper.NoData{
		Status:  "OK",
		Message: fmt.Sprintf("Deleted job ID %d successfully", atoi),
	})
}

func (h JobsHandlerImpl) GetJobs(c *gin.Context) {
	data, err := h.JobsUsecase.GetAllJobs(c.Request.Context())
	if err != nil {
		helper.ErrResponses(c, err)
		return
	}

	c.JSON(http.StatusOK, helper.UseData{
		Status:  "OK",
		Message: "Successfully retrieved all jobs",
		Data:    data,
	})
}

func (h JobsHandlerImpl) GetJobById(c *gin.Context) {
	id := c.Param("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.NoData{
			Status:  "error",
			Message: "ID is not suitable",
		})
		return
	}
	data, err := h.JobsUsecase.GetJobsById(c.Request.Context(), atoi)
	if err != nil {
		helper.ErrResponses(c, err)
		return
	}

	c.JSON(http.StatusOK, helper.UseData{
		Status:  "OK",
		Message: fmt.Sprintf("Successfully retrieved jobs where id %d", atoi),
		Data:    data,
	})
}

func (h JobsHandlerImpl) GetSimilarJobsName(c *gin.Context) {
	namePattern := c.Param("namePattern")

	jobs, err := h.JobsUsecase.GetJobsSimilarName(c.Request.Context(), namePattern)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.NoData{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, helper.UseData{
		Status:  "OK",
		Message: "Successfully retrieved the jobs similar name",
		Data:    jobs,
	})
}
