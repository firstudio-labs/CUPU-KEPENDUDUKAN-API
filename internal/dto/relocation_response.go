package dto

import (
	"github.com/firstudio-lab/JARITMAS-API/internal/entity"
)

type RelocationResponse struct {
	ID        uint  `json:"id"`
	KKRequest int64 `json:"kk_request" validate:"required"`
	///
	ReasonMoving             string `json:"reason_moving"`
	ClassificationRelocation string `json:"classification_relocation"`
	MovingDate               string `json:"moving_date"`
	RelocationType           string `json:"relocation_type"`
	StatusKKMove             string `json:"status_kk_move"`
	NIKFamilyMove            string `json:"nik_family_move"`
	StatusKKStay             string `json:"status_kk_stay"`
	NewProvinceID            int    `json:"new_province_id"`
	NewDistrictID            int    `json:"new_district_id"`
	NewSubDistrictID         int    `json:"new_sub_district_id"`
	NewVillageID             int    `json:"new_village_id"`
	NewRT                    string `json:"new_rt"`
	NewRW                    string `json:"new_rw"`
	///
	NewKK              *int64 `json:"new_kk"`
	NewHeadOfFamily    *int64 `json:"new_head_of_family"`
	NIKFamilyStay      string `json:"nik_family_stay"`
	ProvinceIDStay     *int   `json:"province_id_stay"`
	DistrictIDStay     *int   `json:"district_id_stay"`
	SubDistrictIDStay  *int   `json:"sub_district_id_stay"`
	VillageIDStay      *int   `json:"village_id_stay"`
	VerificationStatus bool   `json:"verification_status"`
}

func RelocationsEntityToDTO(request *[]entity.Relocation) []RelocationResponse {
	var relocations []RelocationResponse
	for _, i := range *request {
		relocations = append(relocations, RelocationResponse{
			ID:                       i.ID,
			KKRequest:                i.KKRequest,
			ReasonMoving:             i.ReasonMoving,
			ClassificationRelocation: i.ClassificationRelocation,
			MovingDate:               i.MovingDate,
			RelocationType:           i.RelocationType,
			StatusKKMove:             i.StatusKKMove,
			NIKFamilyMove:            i.NIKFamilyMove,
			StatusKKStay:             i.StatusKKStay,
			NewProvinceID:            i.NewProvinceID,
			NewDistrictID:            i.NewDistrictID,
			NewSubDistrictID:         i.NewSubDistrictID,
			NewVillageID:             i.NewVillageID,
			NewRT:                    i.NewRT,
			NewRW:                    i.NewRW,
			NewKK:                    i.NewKK,
			NewHeadOfFamily:          i.NewHeadOfFamily,
			NIKFamilyStay:            *i.NIKFamilyStay,
			ProvinceIDStay:           i.ProvinceIDStay,
			DistrictIDStay:           i.DistrictIDStay,
			SubDistrictIDStay:        i.SubDistrictIDStay,
			VillageIDStay:            i.VillageIDStay,
			VerificationStatus:       i.VerificationStatus,
		})
	}

	return relocations
}
