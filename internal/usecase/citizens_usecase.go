package usecase

import (
	"context"
	"github.com/firstudio-lab/JARITMAS-API/internal/dto"
	"github.com/firstudio-lab/JARITMAS-API/internal/entity"
	"github.com/firstudio-lab/JARITMAS-API/internal/repository"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CitizensUsecase interface {
	FindCitizenByNIK(ctx context.Context, nik int) (entity.Citizen, error)
	FindCitizenPage(ctx context.Context, page int) ([]entity.Citizen, error)
	CreateCitizen(ctx context.Context, request dto.CitizenReqCreate) error
	UpdateCitizenByNIK(ctx context.Context, nik int, request dto.CitizenReqUpdate) error
	DeleteCitizenByNIK(ctx context.Context, nik int) error
}

type CitizensUsecaseImpl struct {
	repository.CitizensRepository
	*validator.Validate
	DB *gorm.DB
}

func (u CitizensUsecaseImpl) FindCitizenByNIK(ctx context.Context, nik int) (entity.Citizen, error) {
	citizenByNIK, err := u.CitizensRepository.GetCitizenByNIK(ctx, u.DB, nik)
	if err != nil {
		return entity.Citizen{}, err
	}

	return citizenByNIK, nil
}

func (u CitizensUsecaseImpl) FindCitizenPage(ctx context.Context, page int) ([]entity.Citizen, error) {
	perPage, err := u.CitizensRepository.GetAllCitizenPerPage(ctx, u.DB, page)
	if err != nil {
		return nil, err
	}

	return perPage, nil

}

func (u CitizensUsecaseImpl) CreateCitizen(ctx context.Context, request dto.CitizenReqCreate) error {
	//TODO implement me
	panic("implement me")
}

func (u CitizensUsecaseImpl) UpdateCitizenByNIK(ctx context.Context, nik int, request dto.CitizenReqUpdate) error {
	//TODO implement me
	panic("implement me")
}

func (u CitizensUsecaseImpl) DeleteCitizenByNIK(ctx context.Context, nik int) error {
	//TODO implement me
	panic("implement me")
}
