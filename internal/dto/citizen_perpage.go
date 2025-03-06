package dto

import (
	"github.com/firstudio-lab/JARITMAS-API/internal/entity"
)

type CitizenResponse struct {
	Pagination `json:"pagination"`
	Citizens   []CitizensDTO `json:"citizens"`
}

type Pagination struct {
	CurrentPage  int `json:"current_page"`
	TotalPage    int `json:"total_page"`
	TotalItems   int `json:"total_items"`
	ItemsPerPage int `json:"items_per_page"`
	NextPage     int `json:"next_page"`
	PrevPage     int `json:"prev_page"`
}

type CitizensDTO struct {
	ID                     int    `json:"id"`
	NIK                    int64  `json:"nik"`
	KK                     int64  `json:"kk"`
	FullName               string `json:"full_name"`
	Gender                 string `json:"gender"`
	BirthDate              string `json:"birth_date"`
	Age                    int    `json:"age"`
	BirthPlace             string `json:"birth_place"`
	Address                string `json:"address"`
	ProvinceID             int    `json:"province_id"`
	DistrictID             int    `json:"district_id"`
	SubDistrictID          int    `json:"sub_district_id"`
	VillageID              int    `json:"village_id"`
	RT                     string `json:"rt"`
	RW                     string `json:"rw"`
	PostalCode             int    `json:"postal_code"`
	CitizenStatus          string `json:"citizen_status"`
	BirthCertificate       string `json:"birth_certificate"`
	BirthCertificateNo     string `json:"birth_certificate_no"`
	BloodType              string `json:"blood_type"`
	Religion               string `json:"religion"`
	MaritalStatus          string `json:"marital_status"`
	MaritalCertificate     string `json:"marital_certificate"`
	MaritalCertificateNo   string `json:"marital_certificate_no"`
	MarriageDate           string `json:"marriage_date"`
	DivorceCertificate     string `json:"divorce_certificate"`
	DivorceCertificateNo   string `json:"divorce_certificate_no"`
	DivorceCertificateDate string `json:"divorce_certificate_date"`
	FamilyStatus           string `json:"family_status"`
	MentalDisorders        string `json:"mental_disorders"`
	Disabilities           string `json:"disabilities"`
	EducationStatus        string `json:"education_status"`
	JobTypeID              int    `json:"job_type_id"`
	NIKMother              string `json:"nik_mother"`
	Mother                 string `json:"mother"`
	NIKFather              string `json:"nik_father"`
	Father                 string `json:"father"`
	Coordinate             string `json:"coordinate"`
}

func CitizensDTOtoEntity(request entity.Citizen) CitizensDTO {
	return CitizensDTO{
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

}

func CitizensDTOtoEntities(c []entity.Citizen) []CitizensDTO {
	var citizens []CitizensDTO
	for _, request := range c {
		newCitizen := CitizensDTO{
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
	return citizens
}
