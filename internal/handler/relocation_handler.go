package handler

import (
	"fmt"
	"github.com/firstudio-lab/JARITMAS-API/internal/dto"
	"github.com/firstudio-lab/JARITMAS-API/internal/entity"
	"github.com/firstudio-lab/JARITMAS-API/pkg/helper"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type RelocationHandler interface {
	AddRelocation(c *gin.Context)
	UpdateRelocation(c *gin.Context)
	ApproveRelocation(c *gin.Context)
	GetPerPage(c *gin.Context)
	DeleteRelocation(c *gin.Context)
}

type RelocationHandlerImpl struct {
	*validator.Validate
	*gorm.DB
}

func NewRelocationHandler(validate *validator.Validate, DB *gorm.DB) *RelocationHandlerImpl {
	return &RelocationHandlerImpl{Validate: validate, DB: DB}
}

func (r RelocationHandlerImpl) ApproveRelocation(c *gin.Context) {

	//TODO implement me
	panic("implement me")
}

func (r RelocationHandlerImpl) AddRelocation(c *gin.Context) {
	var body dto.RelocationRequest // letter we make dto for fast
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, helper.NoData{Status: "error", Message: "Failed to parse JSON"})
		return
	}
	// VALIDATION request body
	if err := r.Validate.Struct(&body); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errorMessages []string
		for _, validationError := range validationErrors {
			errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' is invalid: %s %s", validationError.Field(), validationError.Tag(), validationError.Param()))
		}
		errValidate := fmt.Sprintf("validation failed: %s", strings.Join(errorMessages, ", "))
		vErr := fmt.Errorf("%d:%v", http.StatusBadRequest, errValidate)
		helper.ErrResponses(c, vErr)
		return
	}

	// @TRANSACTION
	err := r.DB.WithContext(c.Request.Context()).Transaction(func(tx *gorm.DB) error {
		//checking kk exists
		if err := r.DB.Where("no_kk = ?", body.KKRequest).First(&entity.Citizen{}).Error; err != nil {
			return fmt.Errorf("%d:%s", http.StatusNotFound, "KK not found")
		}

		//checking family move
		if len(body.NIKFamilyMove) != 0 {
			for _, v := range body.NIKFamilyMove {
				var citizen entity.Citizen
				if err := tx.Where("nik = ?", v).First(&citizen).Error; err != nil {
					return fmt.Errorf("%d:%s", http.StatusNotFound, fmt.Sprintf("Family with NIK %v not found", v))
				}
			}
		} else {
			return fmt.Errorf("%d:%s", http.StatusBadRequest, "Family Move is required atleast insert 1 nik")
		}

		//checking family stay
		if len(body.NIKFamilyStay) != 0 {
			for _, v := range body.NIKFamilyStay {
				var citizen entity.Citizen
				if err := tx.Where("nik = ?", v).First(&citizen).Error; err != nil {
					return fmt.Errorf("%d:%s", http.StatusNotFound, fmt.Sprintf("Family with NIK %v not found", v))
				}
			}
		}

		//MAPPING FIRST
		requestToEntity := dto.RelocationRequestToEntity(body)
		err := tx.Create(&requestToEntity).Error
		if err != nil {
			return fmt.Errorf("%d:%s", http.StatusInternalServerError, "Failed to create relocation")
		}

		return nil
	})

	if err != nil {
		helper.ErrResponses(c, err)
		return
	}

	c.JSON(http.StatusCreated, helper.NoData{Status: "CREATED", Message: "Successfully created relocation"})
}

func (r RelocationHandlerImpl) UpdateRelocation(c *gin.Context) {
	params := c.Params.ByName("id")
	atoi, err := strconv.Atoi(params)
	if err != nil {
		err = fmt.Errorf("%d:%s", http.StatusBadRequest, "Failed to parse ID")
		helper.ErrResponses(c, err)
		return
	}

	var body entity.Relocation // letter we make dto for fast
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, helper.NoData{Status: "error", Message: "Failed to parse JSON"})
		return
	}

	if err = r.DB.WithContext(c.Request.Context()).Model(&entity.Relocation{}).Where("id = ?", atoi).Updates(&body).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helper.NoData{Status: "error", Message: "Failed to update relocation"})
	}

	c.JSON(http.StatusCreated, helper.NoData{Status: "CREATED", Message: "Successfully created relocation"})

}

func (r RelocationHandlerImpl) GetPerPage(c *gin.Context) {
	page := c.Query("page")
	pageFind, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		err = fmt.Errorf("%d:%s", http.StatusBadRequest, fmt.Sprintf("make sure page param number %v", pageFind))
		helper.ErrResponses(c, err)
		return
	}

	var totalItem int64
	if err = r.DB.WithContext(c.Request.Context()).Model(&entity.Relocation{}).Count(&totalItem).Error; err != nil {
		err = fmt.Errorf("%d:%s", http.StatusInternalServerError, "try again later")
		helper.ErrResponses(c, err)
		return
	}

	itemPerPage := 10
	pagination := dto.NewPagination(totalItem, int(pageFind), itemPerPage)
	offset := (pageFind - 1) * int64(itemPerPage)

	/// TAKE DATA FROM DATABASE WITH PAGINATION AND NOT SOFT DELETE
	var relocations []entity.Relocation
	if err = r.DB.WithContext(c.Request.Context()).Where("deleted_at IS NULL").Limit(itemPerPage).Offset(int(offset)).Find(&relocations).Error; err != nil {
		err = fmt.Errorf("%d:%s", http.StatusInternalServerError, "try again later")
		helper.ErrResponses(c, err)
		return
	}

	c.JSON(http.StatusOK, helper.UseData{Status: "OK", Message: "Successfully get data",
		Data: struct {
			Pagination  *dto.Pagination     `json:"pagination"`
			Relocations []entity.Relocation `json:"relocations"`
		}{pagination, relocations},
	})
}

func (r RelocationHandlerImpl) DeleteRelocation(c *gin.Context) {
	params := c.Params.ByName("id")
	atoi, err := strconv.Atoi(params)
	if err != nil {
		err = fmt.Errorf("%d:%s", http.StatusBadRequest, "Failed to parse ID")
		helper.ErrResponses(c, err)
		return
	}

	//DOING SOFT DELETE
	if err := r.DB.WithContext(c.Request.Context()).Model(&entity.Relocation{}).Where("id = ?", atoi).Update("deleted_at", time.Now().UnixNano()); err != nil {
		err := fmt.Errorf("%d:%s", http.StatusInternalServerError, "Failed to delete relocation")
		helper.ErrResponses(c, err)
		return
	}

	c.JSON(http.StatusOK, helper.NoData{Status: "OK", Message: "Successfully delete relocation"})
}
