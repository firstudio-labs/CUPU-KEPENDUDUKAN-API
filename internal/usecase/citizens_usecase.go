package usecase

import (
	"context"
	"fmt"
	"github.com/firstudio-lab/JARITMAS-API/internal/dto"
	"github.com/firstudio-lab/JARITMAS-API/internal/entity"
	"github.com/firstudio-lab/JARITMAS-API/internal/repository"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"math"
	"net/http"
	"strings"
)

type CitizensUsecase interface {
	FindCitizenByNIK(ctx context.Context, nik int64) (dto.CitizensDTO, error)
	FindCitizenPage(ctx context.Context, page int) (dto.CitizenResponse, error)
	CreateCitizen(ctx context.Context, request dto.CitizenReqCreate) error
	UpdateCitizenByNIK(ctx context.Context, nik int64, request dto.CitizenReqUpdate) error
	DeleteCitizenByNIK(ctx context.Context, nik int64) error
	FindMemberByKK(ctx context.Context, kk int64) ([]dto.CitizensDTO, error)
	FindAllCitizens(ctx context.Context) ([]dto.CitizensDTO, error)
	FindNameSimilar(ctx context.Context, namePattern string) ([]dto.SimilarNameResponse, error)
}

type CitizensUsecaseImpl struct {
	repository.CitizensRepository
	*validator.Validate
	DB *gorm.DB
}

func NewCitizensUsecase(citizensRepository repository.CitizensRepository, validate *validator.Validate, DB *gorm.DB) *CitizensUsecaseImpl {
	return &CitizensUsecaseImpl{CitizensRepository: citizensRepository, Validate: validate, DB: DB}
}

func (u CitizensUsecaseImpl) FindCitizenByNIK(ctx context.Context, nik int64) (dto.CitizensDTO, error) {
	request, err := u.CitizensRepository.GetCitizenByNIK(ctx, u.DB, nik)
	if err != nil {
		return dto.CitizensDTO{}, fmt.Errorf("%d:%w", http.StatusNotFound, err)
	}

	Citizen := dto.CitizensDTOtoEntity(request)

	return Citizen, nil
}

func (u CitizensUsecaseImpl) FindCitizenPage(ctx context.Context, page int) (dto.CitizenResponse, error) {
	var totalItems int64
	if err := u.DB.Model(&entity.Citizen{}).Count(&totalItems).Error; err != nil {
		return dto.CitizenResponse{}, fmt.Errorf("%d:%w", http.StatusInternalServerError, fmt.Errorf("data citizen not found: %w", err))
	}

	perPage, err := u.CitizensRepository.GetAllCitizenPerPage(ctx, u.DB, page)
	if err != nil {
		return dto.CitizenResponse{}, fmt.Errorf("%d:%w", http.StatusNotFound, err)
	}

	// Tentukan jumlah item per halaman (sudah fix 10) jika mau edit cek reposiutory ya
	itemsPerPage := 10
	totalPage := int(math.Ceil(float64(totalItems) / float64(itemsPerPage)))

	// Tentukan halaman berikutnya dan sebelumnya
	var nextPage, prevPage int
	if page < totalPage {
		nextPage = page + 1
	} else {
		nextPage = totalPage
	}

	if page > 1 {
		prevPage = page - 1
	} else {
		prevPage = 1
	}

	pagination := dto.Pagination{
		CurrentPage:  page,
		TotalPage:    totalPage,
		TotalItems:   int(totalItems),
		ItemsPerPage: 10,
		NextPage:     nextPage,
		PrevPage:     prevPage,
	}

	citizens := dto.CitizensDTOtoEntities(perPage)

	response := dto.CitizenResponse{
		Pagination: pagination,
		Citizens:   citizens,
	}

	return response, nil

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
		FamilyStatus:           request.FamilyStatus.ToString(),
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

	if err := u.CitizensRepository.CreateCitizen(ctx, u.DB, newCitizen); err != nil {
		return fmt.Errorf("%d:%w", http.StatusBadRequest, err)
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
		FamilyStatus:           request.FamilyStatus.ToString(),
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

	if err := u.CitizensRepository.DeleteCitizenByNIK(ctx, u.DB, nik); err != nil {
		return fmt.Errorf("%d:%w", http.StatusInternalServerError, err)
	}

	return nil
}

func (u CitizensUsecaseImpl) FindMemberByKK(ctx context.Context, kk int64) ([]dto.CitizensDTO, error) {
	familyMember, err := u.CitizensRepository.FindMemberByKK(ctx, u.DB, kk)
	if err != nil {
		return nil, fmt.Errorf("%d:%w", http.StatusInternalServerError, err)
	}

	var citizens []dto.CitizensDTO
	for _, request := range familyMember {
		newCitizen := dto.CitizensDTO{
			ID:                     request.ID,
			NIK:                    request.NIK,
			KK:                     request.KK,
			FullName:               request.FullName,
			Gender:                 request.Gender,
			BirthDate:              request.BirthDate,
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
			CitizenStatus:          request.CitizenStatus,
			BirthCertificate:       request.BirthCertificate,
			BirthCertificateNo:     request.BirthCertificateNo,
			BloodType:              request.BloodType,
			Religion:               request.Religion,
			MaritalStatus:          request.MaritalStatus,
			MaritalCertificate:     request.MaritalCertificate,
			MaritalCertificateNo:   request.MaritalCertificateNo,
			MarriageDate:           request.MarriageDate,
			DivorceCertificate:     request.DivorceCertificate,
			DivorceCertificateNo:   request.DivorceCertificateNo,
			DivorceCertificateDate: request.DivorceCertificateDate,
			FamilyStatus:           request.FamilyStatus,
			MentalDisorders:        request.MentalDisorders,
			Disabilities:           request.Disabilities,
			EducationStatus:        request.EducationStatus,
			JobTypeID:              request.JobTypeID,
			NIKMother:              request.NIKMother,
			Mother:                 request.Mother,
			NIKFather:              request.NIKFather,
			Father:                 request.Father,
			Coordinate:             request.Coordinate,
		}
		citizens = append(citizens, newCitizen)
	}

	return citizens, err
}

func (u CitizensUsecaseImpl) FindAllCitizens(ctx context.Context) ([]dto.CitizensDTO, error) {
	allCitizens, err := u.CitizensRepository.FindAllCitizens(ctx, u.DB)
	if err != nil {
		return nil, fmt.Errorf("%d:%w", http.StatusInternalServerError, err)
	}

	citizens := dto.CitizensDTOtoEntities(allCitizens)

	return citizens, err

}

func (u CitizensUsecaseImpl) FindNameSimilar(ctx context.Context, namePattern string) ([]dto.SimilarNameResponse, error) {
	similar, err := u.CitizensRepository.FindNameSimilar(ctx, u.DB, namePattern)
	if err != nil {
		return nil, fmt.Errorf("%d:%w", http.StatusBadRequest, err)
	}

	var results []dto.SimilarNameResponse
	for _, v := range similar {
		n := dto.SimilarNameResponse{FullName: v.FullName}
		results = append(results, n)
	}

	return results, nil
}
