package JARITMAS_API

import (
	"context"
	"fmt"
	"github.com/firstudio-lab/JARITMAS-API/cfg"
	"github.com/firstudio-lab/JARITMAS-API/internal/dto"
	"github.com/firstudio-lab/JARITMAS-API/internal/repository"
	"github.com/firstudio-lab/JARITMAS-API/internal/usecase"
	"github.com/go-playground/validator/v10"
	"testing"
	"time"
)

func TestInsertData(t *testing.T) {

	ctx := context.Background()

	DEBE, _ := cfg.GetPool(cfg.GetConfig())
	validate := validator.New()
	citizensRepository := repository.NewCitizensRepository()
	citizensUsecase := usecase.NewCitizensUsecase(citizensRepository, validate, DEBE)

	reqCreate := dto.CitizenReqCreate{
		NIK:                    1234567890123456,
		KK:                     1234567891234567,
		FullName:               "HEHEHE",
		Gender:                 1,
		BirthDate:              time.Now(),
		Age:                    19,
		BirthPlace:             "DENMARK",
		Address:                "Osaka",
		ProvinceID:             2,
		DistrictID:             1,
		SubDistrictID:          2,
		VillageID:              3,
		RT:                     "02",
		RW:                     "01",
		PostalCode:             3045,
		CitizenStatus:          2,
		BirthCertificate:       1,
		BirthCertificateNo:     1234567891234567,
		BloodType:              2,
		Religion:               1,
		MaritalStatus:          2,
		MaritalCertificate:     1,
		MaritalCertificateNo:   1234567891234567,
		MarriageDate:           time.Now(),
		DivorceCertificate:     2,
		DivorceCertificateNo:   2,
		DivorceCertificateDate: time.Now(),
		FamilyStatusID:         2,
		MentalDisorders:        2,
		Disabilities:           2,
		EducationStatus:        2,
		JobTypeID:              2,
		NIKMother:              "1234567891234567",
		Mother:                 "ibueee",
		NIKFather:              "1234567891234567",
		Father:                 "buapak",
		Coordinate:             "1234567891234567_ 1234567891234567",
	}

	if err := citizensUsecase.CreateCitizen(ctx, reqCreate); err != nil {
		fmt.Println("Error rek")
	}
	fmt.Println("SUccess ")
}
