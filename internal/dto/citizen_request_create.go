package dto

import (
	"github.com/firstudio-lab/JARITMAS-API/internal/entity"
	"time"
)

type CitizenReqCreate struct {
	NIK                    int64 `validate:"required,min=16,max=16"`
	KK                     int64 `validate:"required,min=16,max=16"`
	FullName               string
	Gender                 entity.GenderOptions `validate:"required,oneof=1 2"`
	BirthDate              time.Time
	Age                    int
	BirthPlace             string
	Address                string
	ProvinceID             int
	DistrictID             int
	SubDistrictID          int
	VillageID              int
	RT                     string
	RW                     string
	PostalCode             int
	CitizenStatus          entity.CitizenStatusOption `validate:"required,oneof=1 2"` //(WNI, WNA)
	BirthCertificate       entity.AvailableStatus     `validate:"required,oneof=1 2"` // ADA/TIDAK-ADA
	BirthCertificateNo     int64
	BloodType              entity.BloodType           `validate:"required,oneof=1 2 3 4 5 6 7 8 9 10 11 12 13"`
	Religion               entity.ReligionOption      `validate:"required,oneof=1 2 3 4 5 6 7"`
	MaritalStatus          entity.MaritalStatusOption `validate:"required,oneof=1 2 3 4 5 6"`
	MaritalCertificate     entity.AvailableStatus     `validate:"required,oneof=1 2"` // ADA/TIDAK-ADA
	MaritalCertificateNo   int64
	MarriageDate           time.Time
	DivorceCertificate     entity.AvailableStatus `validate:"required,oneof=1 2"` // ADA/TIDAK-ADA
	DivorceCertificateNo   int64
	DivorceCertificateDate time.Time
	FamilyStatusID         int
	MentalDisorders        entity.AvailableStatus       `validate:"required,oneof=1 2"` // ADA/TIDAK-ADA
	Disabilities           entity.DisablitesStatus      `validate:"required,oneof=1 2 3 4 5 6"`
	EducationStatus        entity.EducationStatusOption `validate:"required,oneof=1 2 3 4 5 6 7 8 9 10"`
	JobTypeID              int
	NIKMother              string
	Mother                 string
	NIKFather              string
	Father                 string
	Coordinate             string
}
