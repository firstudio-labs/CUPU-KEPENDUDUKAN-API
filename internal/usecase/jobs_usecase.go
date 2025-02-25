package usecase

import (
	"context"
	"fmt"
	"github.com/firstudio-lab/JARITMAS-API/internal/dto"
	"github.com/firstudio-lab/JARITMAS-API/internal/entity"
	"github.com/firstudio-lab/JARITMAS-API/internal/repository"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

type JobsUsecase interface {
	GetAllJobs(ctx context.Context) ([]entity.Job, error)
	CreateJobs(ctx context.Context, request dto.JobReqCreate) error
	UpdateJobs(ctx context.Context, idjobs int, request dto.JobReqUpdate) error
	DeleteJobs(ctx context.Context, idjobs int) error
}
type JobsUsecaseImpl struct {
	*validator.Validate
	repository.JobsRepository
}

func NewJobsUsecase(validate *validator.Validate, jobsRepository repository.JobsRepository) *JobsUsecaseImpl {
	return &JobsUsecaseImpl{Validate: validate, JobsRepository: jobsRepository}
}

func (u JobsUsecaseImpl) GetAllJobs(ctx context.Context) ([]entity.Job, error) {
	jobs, err := u.JobsRepository.FindAllJobs(ctx)
	if err != nil {
		return nil, fmt.Errorf("%d:%w", http.StatusInternalServerError, err)
	}
	return jobs, nil
}

func (u JobsUsecaseImpl) CreateJobs(ctx context.Context, request dto.JobReqCreate) error {
	if err := u.Validate.Struct(&request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errorMessages []string
		for _, validationError := range validationErrors {
			errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' is invalid: %s %s", validationError.Field(), validationError.Tag(), validationError.Param()))
		}

		errValidate := fmt.Sprintf("validation failed: %s", strings.Join(errorMessages, ", "))
		return fmt.Errorf("%s", errValidate)

	}

	err := u.JobsRepository.ExistJobCode(ctx, request.Code)
	if err != nil {
		if err := u.JobsRepository.CreateJobs(ctx, entity.Job{Code: request.Code, Name: request.Name}); err != nil {
			return fmt.Errorf("%d:%w", http.StatusBadRequest, err)
		}
		return nil
	}

	return fmt.Errorf("%d:%w", http.StatusBadRequest, err)
}

func (u JobsUsecaseImpl) UpdateJobs(ctx context.Context, idjobs int, request dto.JobReqUpdate) error {
	if err := u.Validate.Struct(&request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errorMessages []string
		for _, validationError := range validationErrors {
			errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' is invalid: %s %s", validationError.Field(), validationError.Tag(), validationError.Param()))
		}
		errValidate := fmt.Sprintf("validation failed: %s", strings.Join(errorMessages, ", "))
		return fmt.Errorf("%s", errValidate)
	}

	err := u.JobsRepository.ExistJobCode(ctx, request.Code)
	if err != nil {
		//KARENA TIDAK EXIST KITA CREATE
		if err := u.JobsRepository.UpdateJobById(ctx, idjobs, entity.Job{Code: request.Code, Name: request.Name}); err != nil {
			return fmt.Errorf("%d:%w", http.StatusBadRequest, err)
		}
		return nil
	}
	return fmt.Errorf("%d:%w", http.StatusBadRequest, err)
}

func (u JobsUsecaseImpl) DeleteJobs(ctx context.Context, idjobs int) error {
	if err := u.JobsRepository.DeleteJobById(ctx, idjobs); err != nil {
		return fmt.Errorf("%d:%w", http.StatusBadRequest, err)
	}

	return nil
}
