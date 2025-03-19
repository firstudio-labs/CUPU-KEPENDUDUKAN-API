package dto

import (
	"encoding/json"
	"github.com/firstudio-lab/JARITMAS-API/internal/entity"
	"time"
)

type RelocationRequest struct {
	KKRequest int64 `json:"kk_request" validate:"required"`
	///
	ReasonMoving             entity.ReasonsMoving            `json:"reason_moving" validate:"oneof=1 2 3 4 5 6 7 "`
	ClassificationRelocation entity.ClassificationRelocation `json:"classification_relocation" validate:"oneof=1 2 3 4 5"`
	MovingDate               string                          `json:"moving_date" validate:"required"`
	RelocationType           entity.RelocationType           `json:"relocation_type" validate:"oneof=1 2 3 4"`
	StatusKKMove             entity.StatusKKStay             `json:"status_kk_move" validate:"oneof=1 2 3 4"`
	NIKFamilyMove            []*int64                        `json:"nik_family_move" validate:"required"`
	StatusKKStay             entity.StatusKKMove             `json:"status_kk_stay" validate:"oneof=1 2 3 4 5"`
	NewProvinceID            int                             `json:"new_province_id" validate:"required"`
	NewDistrictID            int                             `json:"new_district_id" validate:"required"`
	NewSubDistrictID         int                             `json:"new_sub_district_id" json:"new_sub_district_id"`
	NewVillageID             int                             `json:"new_village_id" validate:"required"`
	NewRT                    string                          `json:"new_rt" validate:"required"`
	NewRW                    string                          `json:"new_rw" validate:"required"`
	///
	NewKK             *int64   `json:"new_kk"`
	NewHeadOfFamily   *int64   `json:"new_head_of_family"`
	NIKFamilyStay     []*int64 `json:"nik_family_stay"`
	ProvinceIDStay    *int     `json:"province_id_stay"`
	DistrictIDStay    *int     `json:"district_id_stay"`
	SubDistrictIDStay *int     `json:"sub_district_id_stay"`
	VillageIDStay     *int     `json:"village_id_stay"`
}

func RelocationRequestToEntity(request RelocationRequest) *entity.Relocation {
	move, _ := json.Marshal(request.NIKFamilyMove)
	stay, _ := json.Marshal(request.NIKFamilyStay)

	moveString := string(move)
	stayString := string(stay)

	return &entity.Relocation{
		KKRequest:                request.KKRequest,
		ReasonMoving:             request.ReasonMoving.ToString(),
		ClassificationRelocation: request.ClassificationRelocation.ToString(),
		MovingDate:               request.MovingDate,
		RelocationType:           request.RelocationType.ToString(),
		StatusKKMove:             request.StatusKKMove.ToString(),
		NIKFamilyMove:            moveString,
		StatusKKStay:             request.StatusKKStay.ToString(),
		NewProvinceID:            request.NewProvinceID,
		NewDistrictID:            request.NewDistrictID,
		NewSubDistrictID:         request.NewSubDistrictID,
		NewVillageID:             request.NewVillageID,
		NewRT:                    request.NewRT,
		NewRW:                    request.NewRW,
		NewKK:                    request.NewKK,
		NewHeadOfFamily:          request.NewHeadOfFamily,
		NIKFamilyStay:            &stayString,
		ProvinceIDStay:           request.ProvinceIDStay,
		DistrictIDStay:           request.DistrictIDStay,
		SubDistrictIDStay:        request.SubDistrictIDStay,
		VillageIDStay:            request.VillageIDStay,
		CreatedAt:                time.Now().UnixNano(),
		VerificationStatus:       false, //DEFAULT KITA SET FLASE DULU
		UpdatedAt:                nil,
		DeletedAt:                nil,
	}

}
