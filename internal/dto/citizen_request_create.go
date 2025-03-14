package dto

import (
	"github.com/firstudio-lab/JARITMAS-API/internal/entity"
)

type CitizenReqCreate struct {
	NIK                    int64                        `validate:"required" json:"nik"`
	KK                     int64                        `validate:"required" json:"kk"`
	FullName               string                       `json:"full_name"`
	Gender                 entity.GenderOptions         `validate:"required,oneof=1 2" json:"gender"`
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
	PostalCode             *int                         `json:"postal_code"`
	CitizenStatus          entity.CitizenStatusOption   `validate:"required,oneof=1 2" json:"citizen_status"`    //(WNI, WNA)
	BirthCertificate       entity.AvailableStatus       `validate:"required,oneof=1 2" json:"birth_certificate"` // ADA/TIDAK-ADA
	BirthCertificateNo     string                       `json:"birth_certificate_no"`
	BloodType              entity.BloodType             `validate:"required,oneof=1 2 3 4 5 6 7 8 9 10 11 12 13" json:"blood_type"`
	Religion               entity.ReligionOption        `validate:"required,oneof=1 2 3 4 5 6 7" json:"religion"`
	MaritalStatus          entity.MaritalStatusOption   `validate:"required,oneof=1 2 3 4 5 6" json:"marital_status"`
	MaritalCertificate     entity.AvailableStatus       `validate:"required,oneof=1 2" json:"marital_certificate"` // ADA/TIDAK-ADA
	MaritalCertificateNo   string                       `json:"marital_certificate_no"`
	MarriageDate           string                       `json:"marriage_date"`
	DivorceCertificate     entity.AvailableStatus       `validate:"required,oneof=1 2" json:"divorce_certificate"` // ADA/TIDAK-ADA
	DivorceCertificateNo   string                       `json:"divorce_certificate_no"`
	DivorceCertificateDate string                       `json:"divorce_certificate_date"`
	FamilyStatus           entity.FamilyStatus          `json:"family_status"`
	MentalDisorders        entity.AvailableStatus       `validate:"required,oneof=1 2" json:"mental_disorders"` // ADA/TIDAK-ADA
	Disabilities           entity.DisablitesStatus      `validate:"required,oneof=1 2 3 4 5 6" json:"disabilities"`
	EducationStatus        entity.EducationStatusOption `validate:"required,oneof=1 2 3 4 5 6 7 8 9 10" json:"education_status"`
	JobTypeID              int                          `json:"job_type_id"`
	NIKMother              string                       `json:"nik_mother"`
	Mother                 string                       `json:"mother"`
	NIKFather              string                       `json:"nik_father"`
	Father                 string                       `json:"father"`
	Coordinate             string                       `json:"coordinate"`
	///NEW
	Telephone         *string `json:"telephone"`
	Email             *string `json:"email"`
	Hamlet            *string `json:"hamlet"` //DUSUN
	ForeignAddress    *string `json:"foreign_address"`
	City              *string `json:"city"`
	State             *string `json:"state"` // PROVINCE OT NEGARA BAGIAN
	Country           *string `json:"country"`
	ForeignPostalCode *string `json:"foreign_postal_code"`
	Status            *string `json:"status"`
}
