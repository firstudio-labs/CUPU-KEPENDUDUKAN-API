package handler

import (
	"github.com/firstudio-lab/JARITMAS-API/internal/dto"
	"github.com/firstudio-lab/JARITMAS-API/internal/usecase"
	"github.com/firstudio-lab/JARITMAS-API/pkg/helper"
	"github.com/firstudio-lab/JARITMAS-API/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CitizenHandler interface {
	FindCitizenByNIK(c *gin.Context)
	FindCitizenPage(c *gin.Context)
	CreateCitizen(c *gin.Context)
	UpdateCitizenByNIK(c *gin.Context)
	DeleteCitizenByNIK(c *gin.Context)
	FindAllMemberByKK(c *gin.Context)
	FindAllCitizens(c *gin.Context)
	FindSimilarName(c *gin.Context)
}

type CitizensHandlerImpl struct {
	usecase.CitizensUsecase
}

func NewCitizensHandler(citizensUsecase usecase.CitizensUsecase) *CitizensHandlerImpl {
	return &CitizensHandlerImpl{CitizensUsecase: citizensUsecase}
}

func (h CitizensHandlerImpl) FindCitizenPage(c *gin.Context) {
	page := c.Query("page")
	atoi, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.NoData{
			Status:  "error",
			Message: "page banggg",
		})
		return
	}

	citizenPage, err := h.CitizensUsecase.FindCitizenPage(c.Request.Context(), int(atoi))
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.NoData{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, helper.UseData{
		Status:  "OK",
		Message: "Successfully retrieved the citizen",
		Data:    citizenPage,
	})
}

func (h CitizensHandlerImpl) FindCitizenByNIK(c *gin.Context) {
	nik := c.Param("nik")
	atoi, err := strconv.ParseInt(nik, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.NoData{
			Status:  "error",
			Message: "NIK is not suitable",
		})
		return
	}

	Citizen, err := h.CitizensUsecase.FindCitizenByNIK(c.Request.Context(), atoi)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.NoData{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, helper.UseData{
		Status:  "OK",
		Message: "Successfully retrieved the citizen",
		Data:    Citizen,
	})
}

func (h CitizensHandlerImpl) CreateCitizen(c *gin.Context) {
	var body dto.CitizenReqCreate
	if err := c.ShouldBindJSON(&body); err != nil {
		logger.Log.Errorf("Fail to parse body: %v", err)
		c.JSON(http.StatusBadRequest, helper.NoData{
			Status:  "error",
			Message: "Failed to parse JSON",
		})
		return
	}

	if err := h.CitizensUsecase.CreateCitizen(c.Request.Context(), body); err != nil {
		c.JSON(http.StatusInternalServerError, helper.NoData{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, helper.NoData{
		Status:  "CREATED",
		Message: "Successfully created new Citizen",
	})
}

func (h CitizensHandlerImpl) UpdateCitizenByNIK(c *gin.Context) {
	var body dto.CitizenReqUpdate
	nik := c.Param("nik")
	atoi, err := strconv.ParseInt(nik, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.NoData{
			Status:  "error",
			Message: "NIK is not suitable",
		})
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		logger.Log.Errorf("Fail to parse body: %v", err)
		c.JSON(http.StatusBadRequest, helper.NoData{
			Status:  "error",
			Message: "Failed to parse JSON",
		})
		return
	}

	if err := h.CitizensUsecase.UpdateCitizenByNIK(c.Request.Context(), atoi, body); err != nil {
		c.JSON(http.StatusInternalServerError, helper.NoData{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, helper.NoData{
		Status:  "OK",
		Message: "Successfully updated Citizen",
	})
}

func (h CitizensHandlerImpl) DeleteCitizenByNIK(c *gin.Context) {
	nik := c.Param("nik")
	atoi, err := strconv.ParseInt(nik, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.NoData{
			Status:  "error",
			Message: "NIK is not suitable",
		})
		return
	}

	if err := h.CitizensUsecase.DeleteCitizenByNIK(c.Request.Context(), atoi); err != nil {
		c.JSON(http.StatusInternalServerError, helper.NoData{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, helper.NoData{
		Status:  "OK",
		Message: "Deleted Citizen successfully",
	})
}

func (h CitizensHandlerImpl) FindAllMemberByKK(c *gin.Context) {
	nik := c.Param("kk")
	atoi, err := strconv.ParseInt(nik, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.NoData{
			Status:  "error",
			Message: "kk is not suitable",
		})
		return
	}

	Citizen, err := h.CitizensUsecase.FindMemberByKK(c.Request.Context(), atoi)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.NoData{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, helper.UseData{
		Status:  "OK",
		Message: "Successfully retrieved the citizen",
		Data:    Citizen,
	})
}

func (h CitizensHandlerImpl) FindAllCitizens(c *gin.Context) {
	Citizen, err := h.CitizensUsecase.FindAllCitizens(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.NoData{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, helper.UseData{
		Status:  "OK",
		Message: "Successfully retrieved the citizen",
		Data:    Citizen,
	})
}

func (h CitizensHandlerImpl) FindSimilarName(c *gin.Context) {
	namePattern := c.Param("namePattern")

	Citizen, err := h.CitizensUsecase.FindNameSimilar(c.Request.Context(), namePattern)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.NoData{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, helper.UseData{
		Status:  "OK",
		Message: "Successfully retrieved the citizen",
		Data:    Citizen,
	})
}
