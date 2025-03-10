package handler

import (
	"encoding/json"
	"fmt"
	"github.com/firstudio-lab/JARITMAS-API/internal/entity"
	"github.com/firstudio-lab/JARITMAS-API/pkg/helper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

type CountryHandler interface {
	GetProvince(ctx *gin.Context)
	GetDistrictByProvinceCode(ctx *gin.Context)
	GetSubDistrictByDistrictCode(ctx *gin.Context)
	GetVillageBySUbDistrictCode(ctx *gin.Context)
}

type CountryHandlerImpl struct {
	*gorm.DB
}

func NewCountryHandler(DB *gorm.DB) *CountryHandlerImpl {
	return &CountryHandlerImpl{DB: DB}
}

func (h CountryHandlerImpl) GetProvince(ctx *gin.Context) {
	var results []entity.IndonesiaProvince
	if err := h.DB.WithContext(ctx.Request.Context()).Select("id", "code", "name", "meta").
		Find(&results).Error; err != nil {
		err := fmt.Errorf("%d:%s", http.StatusInternalServerError, "try again later")
		helper.ErrResponses(ctx, err)
		return
	}

	if len(results) == 0 {
		err := fmt.Errorf("%d:%s", http.StatusNotFound, fmt.Sprintf("data in province not ready yet"))
		helper.ErrResponses(ctx, err)
		return
	}

	// Process the results and map Meta to Lat/Long
	var provinces []struct {
		ID   int    `json:"id"`
		Code string `json:"code"`
		Name string `json:"name"`
		Meta struct {
			Lat  string `json:"lat"`
			Long string `json:"long"`
		} `json:"meta"`
	}

	for _, province := range results {
		// Parse the Meta JSON field
		var meta map[string]string
		if err := json.Unmarshal([]byte(*province.Meta), &meta); err != nil {
			// Handle error if Meta field can't be parsed
			log.Printf("Failed to parse Meta field for province %s: %v", province.Name, err)
			meta = make(map[string]string)
		}

		// Add province to the mapped result
		provinces = append(provinces, struct {
			ID   int    `json:"id"`
			Code string `json:"code"`
			Name string `json:"name"`
			Meta struct {
				Lat  string `json:"lat"`
				Long string `json:"long"`
			} `json:"meta"`
		}{
			ID:   int(province.ID),
			Code: province.Code,
			Name: province.Name,
			Meta: struct {
				Lat  string `json:"lat"`
				Long string `json:"long"`
			}{
				Lat:  meta["lat"],
				Long: meta["long"],
			},
		})
	}

	if len(provinces) == 0 {
		err := fmt.Errorf("%d:%s", http.NotFound, "try again later")
		helper.ErrResponses(ctx, err)
		return
	}

	// Return the provinces with the correct structure
	ctx.JSON(http.StatusOK, helper.UseData{
		Status:  "OK",
		Message: "successfully get provinces",
		Data:    provinces,
	})
}
func (h CountryHandlerImpl) GetDistrictByProvinceCode(ctx *gin.Context) {
	// Extract the province code from the URL parameter and convert to int
	id := ctx.Param("province-code")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		err = fmt.Errorf("%d:%s", http.StatusBadRequest, "province code is not suitable")
		helper.ErrResponses(ctx, err)
		return
	}

	// Query districts by province code
	var results []entity.IndonesiaDistrict
	if err := h.DB.WithContext(ctx.Request.Context()).Select("id", "code", "province_code", "name", "meta").
		Where("province_code =?", atoi).Find(&results).Error; err != nil {
		log.Printf("KENA DINSINI")
		err = fmt.Errorf("%d:%s", http.StatusInternalServerError, "try again later")
		helper.ErrResponses(ctx, err)
		return
	}

	if len(results) == 0 {
		err = fmt.Errorf("%d:%s", http.StatusNotFound, fmt.Sprintf("data with porvince-code %v", atoi))
		helper.ErrResponses(ctx, err)
		return
	}

	// Process the results and map Meta to Lat/Long
	var districts []struct {
		ID           int    `json:"id"`
		Code         string `json:"code"`
		ProvinceCode string `json:"province_code"`
		Name         string `json:"name"`
		Meta         struct {
			Lat  string `json:"lat"`
			Long string `json:"long"`
		} `json:"meta"`
	}

	// Loop through each district and process Meta field
	for _, dist := range results {
		// Parse the Meta JSON field
		var meta map[string]string
		if err := json.Unmarshal([]byte(*dist.Meta), &meta); err != nil {
			// Handle error if Meta field can't be parsed
			log.Printf("Failed to parse Meta field for district %s: %v", dist.Name, err)
			meta = make(map[string]string) // Set meta as empty if parsing fails
		}

		// Add district to the mapped result
		districts = append(districts, struct {
			ID           int    `json:"id"`
			Code         string `json:"code"`
			ProvinceCode string `json:"province_code"`
			Name         string `json:"name"`
			Meta         struct {
				Lat  string `json:"lat"`
				Long string `json:"long"`
			} `json:"meta"`
		}{
			ID:           int(dist.ID),
			Code:         dist.Code,
			ProvinceCode: dist.ProvinceCode,
			Name:         dist.Name,
			Meta: struct {
				Lat  string `json:"lat"`
				Long string `json:"long"`
			}{
				Lat:  meta["lat"],  // Use latitude from Meta
				Long: meta["long"], // Use longitude from Meta
			},
		})
	}

	// Send response back to client
	ctx.JSON(http.StatusOK, helper.UseData{
		Status:  "OK",
		Message: fmt.Sprintf("Retrieved districts for province code %d", atoi),
		Data:    districts,
	})
}
func (h CountryHandlerImpl) GetSubDistrictByDistrictCode(ctx *gin.Context) {
	// Extract the district code from the URL parameter and convert to int
	id := ctx.Param("district-code")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		err = fmt.Errorf("%d:%s", http.StatusBadRequest, "district code is not suitable")
		helper.ErrResponses(ctx, err)
		return
	}

	// Query subdistricts by district code
	var results []entity.IndonesiaSubDistrict
	if err := h.DB.WithContext(ctx.Request.Context()).Select("id", "code", "district_code", "name", "meta").
		Where("district_code =?", atoi).Find(&results).Error; err != nil {
		err := fmt.Errorf("%d:%s", http.StatusInternalServerError, "try again later")
		helper.ErrResponses(ctx, err)
		return
	}

	if len(results) == 0 {
		err = fmt.Errorf("%d:%s", http.StatusNotFound, fmt.Sprintf("data with district-code %v", atoi))
		helper.ErrResponses(ctx, err)
		return
	}

	// Process the results and map Meta to Lat/Long
	var subdistricts []struct {
		ID           int    `json:"id"`
		Code         string `json:"code"`
		DistrictCode string `json:"district_code"`
		Name         string `json:"name"`
		Meta         struct {
			Lat  string `json:"lat"`
			Long string `json:"long"`
		} `json:"meta"`
	}

	// Loop through each subdistrict and process Meta field
	for _, subdist := range results {
		// Parse the Meta JSON field
		var meta map[string]string
		if err := json.Unmarshal([]byte(*subdist.Meta), &meta); err != nil {
			// Handle error if Meta field can't be parsed
			log.Printf("Failed to parse Meta field for subdistrict %s: %v", subdist.Name, err)
			meta = make(map[string]string) // Set meta as empty if parsing fails
		}

		// Add subdistrict to the mapped result
		subdistricts = append(subdistricts, struct {
			ID           int    `json:"id"`
			Code         string `json:"code"`
			DistrictCode string `json:"district_code"`
			Name         string `json:"name"`
			Meta         struct {
				Lat  string `json:"lat"`
				Long string `json:"long"`
			} `json:"meta"`
		}{
			ID:           int(subdist.ID),
			Code:         subdist.Code,
			DistrictCode: subdist.DistrictCode,
			Name:         subdist.Name,
			Meta: struct {
				Lat  string `json:"lat"`
				Long string `json:"long"`
			}{
				Lat:  meta["lat"],  // Use latitude from Meta
				Long: meta["long"], // Use longitude from Meta
			},
		})
	}

	// Return subdistrict data in response
	ctx.JSON(http.StatusOK, helper.UseData{
		Status:  "OK",
		Message: fmt.Sprintf("Retrieved subdistricts for district code %d", atoi),
		Data:    subdistricts,
	})
}

func (h CountryHandlerImpl) GetVillageBySUbDistrictCode(ctx *gin.Context) {
	id := ctx.Param("sub-district-code")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		err = fmt.Errorf("%d:%s", http.StatusBadRequest, "sub district code is not suitable")
		helper.ErrResponses(ctx, err)
		return
	}

	var results []entity.IndonesiaVillage
	if err := h.DB.WithContext(ctx.Request.Context()).Select("id", "code", "sub_district_code", "name", "meta").
		Where("sub_district_code =?", atoi).Find(&results).Error; err != nil {
		err := fmt.Errorf("%d:%s", http.StatusInternalServerError, "try again later")
		helper.ErrResponses(ctx, err)
		return
	}

	if len(results) == 0 {
		err = fmt.Errorf("%d:%s", http.StatusNotFound, fmt.Sprintf("data with sub-district-code %v", atoi))
		helper.ErrResponses(ctx, err)
		return
	}

	// Process the results and map Meta to Lat/Long
	var villages []struct {
		ID              int    `json:"id"`
		Code            string `json:"code"`
		SubDistrictCode string `json:"sub_district_code"`
		Name            string `json:"name"`
		Meta            struct {
			Lat  string `json:"lat"`
			Long string `json:"long"`
		} `json:"meta"`
	}

	for _, village := range results {
		// Parse the Meta JSON field
		var meta map[string]string
		if err := json.Unmarshal([]byte(*village.Meta), &meta); err != nil {
			// Handle error if Meta field can't be parsed
			log.Printf("Failed to parse Meta field for village %s: %v", village.Name, err)
			meta = make(map[string]string) // Set meta as empty if parsing fails
		}

		// Add village to the mapped result
		villages = append(villages, struct {
			ID              int    `json:"id"`
			Code            string `json:"code"`
			SubDistrictCode string `json:"sub_district_code"`
			Name            string `json:"name"`
			Meta            struct {
				Lat  string `json:"lat"`
				Long string `json:"long"`
			} `json:"meta"`
		}{
			ID:              int(village.ID),
			Code:            village.Code,
			SubDistrictCode: village.SubDistrictCode,
			Name:            village.Name,
			Meta: struct {
				Lat  string `json:"lat"`
				Long string `json:"long"`
			}{
				Lat:  meta["lat"],  // Use latitude from Meta
				Long: meta["long"], // Use longitude from Meta
			},
		})
	}
	ctx.JSON(http.StatusOK, helper.UseData{Status: "OK", Message: fmt.Sprintf("retrieve data sub_district where sub_district code %d", atoi), Data: villages})

}
