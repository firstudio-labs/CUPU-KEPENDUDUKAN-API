package dto

import (
	"github.com/firstudio-lab/JARITMAS-API/internal/entity"
)

type CitizenReqUpdate struct {
	KK                     int64                        `json:"kk"`
	FullName               string                       `json:"full_name"`
	Gender                 entity.GenderOptions         `validate:"oneof=1 2" json:"gender"`
	BirthDate              string                       `json:"birth_date"`
	Age                    int                          `json:"age"`
	BirthPlace             string                       `json:"birth_place"`
	Address                string                       `json:"address"`
	ProvinceID             int                          `json:"province_id"`
	DistrictID             int                          `json:"district_id"`
	SubDistrictID          int                          `json:"sub_district_id"`
	VillageID              int                          `json:"village_id"`
	RT                     string                       `json:"rt"`
	RW                     string                       `json:"rw"`
	PostalCode             int                          `json:"postal_code"`
	CitizenStatus          entity.CitizenStatusOption   `validate:"oneof=1 2" json:"citizen_status"`    //(WNI, WNA)
	BirthCertificate       entity.AvailableStatus       `validate:"oneof=1 2" json:"birth_certificate"` // ADA/TIDAK-ADA
	BirthCertificateNo     string                       `json:"birth_certificate_no"`
	BloodType              entity.BloodType             `validate:"oneof=1 2 3 4 5 6 7 8 9 10 11 12 13" json:"blood_type"`
	Religion               entity.ReligionOption        `validate:"oneof=1 2 3 4 5 6 7" json:"religion"`
	MaritalStatus          entity.MaritalStatusOption   `validate:"oneof=1 2 3 4 5 6" json:"marital_status"`
	MaritalCertificate     entity.AvailableStatus       `validate:"oneof=1 2" json:"marital_certificate"` // ADA/TIDAK-ADA
	MaritalCertificateNo   string                       `json:"marital_certificate_no"`
	MarriageDate           string                       `json:"marriage_date"`
	DivorceCertificate     entity.AvailableStatus       `validate:"oneof=1 2" json:"divorce_certificate"` // ADA/TIDAK-ADA
	DivorceCertificateNo   string                       `json:"divorce_certificate_no"`
	DivorceCertificateDate string                       `json:"divorce_certificate_date"`
	FamilyStatus           entity.FamilyStatus          `json:"family_status"`
	MentalDisorders        entity.AvailableStatus       `validate:"oneof=1 2" json:"mental_disorders"` // ADA/TIDAK-ADA
	Disabilities           entity.DisablitesStatus      `validate:"oneof=1 2 3 4 5 6" json:"disabilities"`
	EducationStatus        entity.EducationStatusOption `validate:"oneof=1 2 3 4 5 6 7 8 9 10" json:"education_status"`
	JobTypeID              int                          `json:"job_type_id"`
	NIKMother              string                       `json:"nik_mother"`
	Mother                 string                       `json:"mother"`
	NIKFather              string                       `json:"nik_father"`
	Father                 string                       `json:"father"`
	Coordinate             string                       `json:"coordinate"`
}
