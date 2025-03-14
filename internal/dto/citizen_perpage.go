package dto

import (
	"github.com/firstudio-lab/JARITMAS-API/internal/entity"
	"math"
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

func NewPagination(totalItems int64, currentPage int, itemPerPage int) *Pagination {
	totalPage := int(math.Ceil(float64(totalItems) / float64(itemPerPage)))
	// Tentukan halaman berikutnya dan sebelumnya
	var nextPage, prevPage int
	if currentPage < totalPage {
		nextPage = currentPage + 1
	} else {
		nextPage = totalPage
	}

	if currentPage > 1 {
		prevPage = currentPage - 1
	} else {
		prevPage = 1
	}

	return &Pagination{
		CurrentPage:  currentPage,
		TotalPage:    totalPage,
		TotalItems:   int(totalItems),
		ItemsPerPage: itemPerPage,
		NextPage:     nextPage,
		PrevPage:     prevPage,
	}
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
	PostalCode             *int   `json:"postal_code"`
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
	//NEW
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
		//NEW
		Telephone:         request.Telephone,
		Email:             request.Email,
		Hamlet:            request.Hamlet,
		ForeignAddress:    request.ForeignAddress,
		City:              request.City,
		State:             request.State,
		Country:           request.Country,
		ForeignPostalCode: request.ForeignPostalCode,
		Status:            request.Status,
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
			//NEW
			Telephone:         request.Telephone,
			Email:             request.Email,
			Hamlet:            request.Hamlet,
			ForeignAddress:    request.ForeignAddress,
			City:              request.City,
			State:             request.State,
			Country:           request.Country,
			ForeignPostalCode: request.ForeignPostalCode,
			Status:            request.Status,
		}
		citizens = append(citizens, newCitizen)
	}
	return citizens
}
