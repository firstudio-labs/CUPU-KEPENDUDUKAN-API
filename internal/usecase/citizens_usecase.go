package usecase

import (
	"context"
	"fmt"
	"github.com/firstudio-lab/JARITMAS-API/internal/dto"
	"github.com/firstudio-lab/JARITMAS-API/internal/entity"
	"github.com/firstudio-lab/JARITMAS-API/internal/repository"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type CitizensUsecase interface {
	FindCitizenByNIK(ctx context.Context, nik int64) (entity.Citizen, error)
	FindCitizenPage(ctx context.Context, page int) ([]entity.Citizen, error)
	CreateCitizen(ctx context.Context, request dto.CitizenReqCreate) error
	UpdateCitizenByNIK(ctx context.Context, nik int64, request dto.CitizenReqUpdate) error
	DeleteCitizenByNIK(ctx context.Context, nik int64) error
}

type CitizensUsecaseImpl struct {
	repository.CitizensRepository
	*validator.Validate
	DB *gorm.DB
}

func NewCitizensUsecase(citizensRepository repository.CitizensRepository, validate *validator.Validate, DB *gorm.DB) *CitizensUsecaseImpl {
	return &CitizensUsecaseImpl{CitizensRepository: citizensRepository, Validate: validate, DB: DB}
}

func (u CitizensUsecaseImpl) FindCitizenByNIK(ctx context.Context, nik int64) (entity.Citizen, error) {
	citizenByNIK, err := u.CitizensRepository.GetCitizenByNIK(ctx, u.DB, nik)
	if err != nil {
		return entity.Citizen{}, fmt.Errorf("%d:%w", http.StatusNotFound, err)
	}

	return citizenByNIK, nil
}

func (u CitizensUsecaseImpl) FindCitizenPage(ctx context.Context, page int) ([]entity.Citizen, error) {
	perPage, err := u.CitizensRepository.GetAllCitizenPerPage(ctx, u.DB, page)
	if err != nil {
		return nil, fmt.Errorf("%d:%w", http.StatusNotFound, err)
	}

	return perPage, nil

}

func (u CitizensUsecaseImpl) CreateCitizen(ctx context.Context, request dto.CitizenReqCreate) error {
	if err := u.Validate.Struct(&request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errorMessages []string
		for _, validationError := range validationErrors {
			errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' is invalid: %s %s", validationError.Field(), validationError.Tag(), validationError.Param()))
		}

		errValidate := fmt.Sprintf("validation failed: %s", strings.Join(errorMessages, ", "))
		return fmt.Errorf("%s", errValidate)
	}

	// MAPPING DATA
	newCitizen := entity.Citizen{
		NIK:                    request.NIK,
		KK:                     request.KK,
		FullName:               request.FullName,
		Gender:                 request.Gender.ToString(),
		BirthDate:              request.BirthDate, // tanggal ntar
		Age:                    request.Age,
		BirthPlace:             request.BirthPlace,
		Address:                request.Address,
		ProvinceID:             request.ProvinceID,
		DistrictID:             request.DistrictID,
		SubDistrictID:          request.SubDistrictID,
		VillageID:              request.VillageID,
		RT:                     request.RT,
		RW:                     request.RW,
		PostalCode:             request.PostalCode,
		CitizenStatus:          request.CitizenStatus.ToString(),
		BirthCertificate:       request.BirthCertificate.ToString(), //ada atau tidak ada
		BirthCertificateNo:     request.BirthCertificateNo,
		BloodType:              request.BloodType.ToString(),
		Religion:               request.Religion.ToString(),
		MaritalStatus:          request.MaritalStatus.ToString(),
		MaritalCertificate:     request.MaritalCertificate.ToString(),
		MaritalCertificateNo:   request.MaritalCertificateNo,
		MarriageDate:           request.MarriageDate,
		DivorceCertificate:     request.DivorceCertificate.ToString(),
		DivorceCertificateNo:   request.DivorceCertificateNo,
		DivorceCertificateDate: request.DivorceCertificateDate,
		FamilyStatusID:         request.FamilyStatusID,
		MentalDisorders:        request.MentalDisorders.ToString(),
		Disabilities:           request.Disabilities.ToString(),
		EducationStatus:        request.EducationStatus.ToString(),
		JobTypeID:              request.JobTypeID,
		NIKMother:              request.NIKMother,
		Mother:                 request.Mother,
		NIKFather:              request.NIKFather,
		Father:                 request.Father,
		Coordinate:             request.Coordinate,
	}

	if err := u.CitizensRepository.ExistCitizenNIK(ctx, u.DB, request.NIK); err != nil {
		//KARENA TIDAK EXIST KITA CREATE
		if err := u.CitizensRepository.CreateCitizen(ctx, u.DB, newCitizen); err != nil {
			return fmt.Errorf("%d:%w", http.StatusInternalServerError, err)
		}
	}

	return nil
}

func (u CitizensUsecaseImpl) UpdateCitizenByNIK(ctx context.Context, nik int64, request dto.CitizenReqUpdate) error {
	if err := u.Validate.Struct(&request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errorMessages []string
		for _, validationError := range validationErrors {
			errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' is invalid: %s %s", validationError.Field(), validationError.Tag(), validationError.Param()))
		}

		errValidate := fmt.Sprintf("validation failed: %s", strings.Join(errorMessages, ", "))
		return fmt.Errorf("%s", errValidate)
	}

	if err := u.CitizensRepository.ExistCitizenNIK(ctx, u.DB, nik); err != nil {
		return fmt.Errorf("%d:%w", http.StatusNotFound, err)
	}

	//MAPPING
	updatedCitizen := entity.Citizen{
		KK:                     request.KK,
		FullName:               request.FullName,
		Gender:                 request.Gender.ToString(),
		BirthDate:              request.BirthDate, // tanggal ntar
		Age:                    request.Age,
		BirthPlace:             request.BirthPlace,
		Address:                request.Address,
		ProvinceID:             request.ProvinceID,
		DistrictID:             request.DistrictID,
		SubDistrictID:          request.SubDistrictID,
		VillageID:              request.VillageID,
		RT:                     request.RT,
		RW:                     request.RW,
		PostalCode:             request.PostalCode,
		CitizenStatus:          request.CitizenStatus.ToString(),
		BirthCertificate:       request.BirthCertificate.ToString(), //ada atau tidak ada
		BirthCertificateNo:     request.BirthCertificateNo,
		BloodType:              request.BloodType.ToString(),
		Religion:               request.Religion.ToString(),
		MaritalStatus:          request.MaritalStatus.ToString(),
		MaritalCertificate:     request.MaritalCertificate.ToString(),
		MaritalCertificateNo:   request.MaritalCertificateNo,
		MarriageDate:           request.MarriageDate,
		DivorceCertificate:     request.DivorceCertificate.ToString(),
		DivorceCertificateNo:   request.DivorceCertificateNo,
		DivorceCertificateDate: request.DivorceCertificateDate,
		FamilyStatusID:         request.FamilyStatusID,
		MentalDisorders:        request.MentalDisorders.ToString(),
		Disabilities:           request.Disabilities.ToString(),
		EducationStatus:        request.EducationStatus.ToString(),
		JobTypeID:              request.JobTypeID,
		NIKMother:              request.NIKMother,
		Mother:                 request.Mother,
		NIKFather:              request.NIKFather,
		Father:                 request.Father,
		Coordinate:             request.Coordinate,
	}

	if err := u.CitizensRepository.UpdateCitizen(ctx, u.DB, nik, updatedCitizen); err != nil {
		return fmt.Errorf("%d:%w", http.StatusInternalServerError, err)
	}

	return nil
}

func (u CitizensUsecaseImpl) DeleteCitizenByNIK(ctx context.Context, nik int64) error {
	if err := u.CitizensRepository.ExistCitizenNIK(ctx, u.DB, nik); err != nil {
		return fmt.Errorf("%d:%w", http.StatusNotFound, err)
	}

	if err := u.CitizensRepository.DeleteCitizenByNIK(ctx, u.DB, nik); err != nil {
		return fmt.Errorf("%d:%w", http.StatusInternalServerError, err)
	}

	return nil
}
