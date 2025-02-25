package handler

import (
	"fmt"
	"github.com/firstudio-lab/JARITMAS-API/internal/dto"
	"github.com/firstudio-lab/JARITMAS-API/internal/usecase"
	"github.com/firstudio-lab/JARITMAS-API/pkg/helper"
	"github.com/firstudio-lab/JARITMAS-API/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type JobsHandler interface {
	CreateJob(ctx *fiber.Ctx) error
	UpdateJobById(ctx *fiber.Ctx) error
	DeleteJobById(ctx *fiber.Ctx) error
	GetJobs(ctx *fiber.Ctx) error
}

type JobsHandlerImpl struct {
	usecase.JobsUsecase
}

func NewJobsHandler(jobsUsecase usecase.JobsUsecase) *JobsHandlerImpl {
	return &JobsHandlerImpl{JobsUsecase: jobsUsecase}
}

func (h JobsHandlerImpl) CreateJob(ctx *fiber.Ctx) error {
	var body dto.JobReqCreate
	if err := ctx.BodyParser(&body); err != nil {
		logger.Log.Errorf("Fail to parse body %e", err)
		err := fmt.Errorf("%d:%v", http.StatusInternalServerError, " failed to parse json")
		return helper.WResponses(ctx, err, "", nil)
	}

	err := h.JobsUsecase.CreateJobs(ctx.Context(), body)
	if err != nil {
		return helper.WResponses(ctx, err, "", nil)
	}

	return ctx.Status(http.StatusCreated).JSON(helper.NoData{
		Status:  "CREATED",
		Message: "successfully create data jobs",
	})
}

func (h JobsHandlerImpl) UpdateJobById(ctx *fiber.Ctx) error {
	params := ctx.Params("id")
	atoi, err := strconv.Atoi(params)
	if err != nil {
		err := fmt.Errorf("%d:%v", http.StatusBadRequest, "id is in corect")
		return helper.WResponses(ctx, err, "", nil)
	}

	var body dto.JobReqUpdate
	if err := ctx.BodyParser(&body); err != nil {
		logger.Log.Errorf("Fail to parse body %e", err)
		err := fmt.Errorf("%d:%v", http.StatusInternalServerError, " failed to parse json")
		return helper.WResponses(ctx, err, "", nil)
	}

	err = h.JobsUsecase.UpdateJobs(ctx.Context(), atoi, body)
	if err != nil {
		return helper.WResponses(ctx, err, "", nil)
	}

	return ctx.Status(http.StatusCreated).JSON(helper.NoData{
		Status:  "OK",
		Message: fmt.Sprintf("successfully update jobs id %d", atoi),
	})
}

func (h JobsHandlerImpl) DeleteJobById(ctx *fiber.Ctx) error {
	params := ctx.Params("id")
	atoi, err := strconv.Atoi(params)
	if err != nil {
		err := fmt.Errorf("%d:%v", http.StatusBadRequest, "id is not suitable")
		return helper.WResponses(ctx, err, "", nil)
	}

	err = h.JobsUsecase.DeleteJobs(ctx.Context(), atoi)
	if err != nil {
		return helper.WResponses(ctx, err, "", nil)
	}

	return ctx.Status(http.StatusCreated).JSON(helper.NoData{
		Status:  "OK",
		Message: fmt.Sprintf("deleted jobs id %d", atoi),
	})
}

func (h JobsHandlerImpl) GetJobs(ctx *fiber.Ctx) error {
	data, err := h.JobsUsecase.GetAllJobs(ctx.Context())
	if err != nil {
		return helper.WResponses(ctx, err, "", nil)
	}

	return ctx.Status(http.StatusCreated).JSON(helper.UseData{
		Status:  "OK",
		Message: "successfully getting all jobs",
		Data:    data,
	})
}
