package JARITMAS_API

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/firstudio-lab/JARITMAS-API/cfg"
	"github.com/firstudio-lab/JARITMAS-API/internal/entity"
	"github.com/firstudio-lab/JARITMAS-API/internal/repository"
	"github.com/firstudio-lab/JARITMAS-API/pkg/logger"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"testing"
)

func TestInsertExcel(t *testing.T) {
	db, _ := cfg.GetPool(cfg.GetConfig())

	file, err := os.Open("apdk_mento_tocsv.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	cvInt64 := func(x string) int64 {
		atoi, _ := strconv.Atoi(x)
		return int64(atoi)
	}

	cvInt := func(x string) int {
		atoi, _ := strconv.Atoi(x)
		return atoi
	}

	for i, record := range records {
		if i == 0 {
			continue
		}

		fmt.Println("record 13", record[13])

		citizen := entity.Citizen{
			NIK:                  cvInt64(record[0]),
			KK:                   cvInt64(record[1]),
			FullName:             record[2],
			Gender:               FormatGender(record[3]),
			BirthDate:            record[4],
			Age:                  cvInt(record[5]),
			BirthPlace:           record[6],
			Address:              record[7],
			ProvinceID:           StrProvinceToInt(db, record[11]),    //
			DistrictID:           StrSubDistrictToInt(db, record[13]), //
			SubDistrictID:        StrDistrictToInt(db, record[15]),    //
			VillageID:            StrVillageToInt(db, record[16]),     //
			RT:                   record[8],
			RW:                   record[9],
			PostalCode:           cvInt(record[10]),
			CitizenStatus:        "WNI",
			BirthCertificate:     IfNotNullAvaliable(record[26]),
			BirthCertificateNo:   record[26],
			BloodType:            BlodCv(record[24]),
			Religion:             ReligionCv(record[22]),
			MaritalStatus:        MaritalCv(record[20]),
			MaritalCertificate:   IfNotNullAvaliable(record[28]),
			MaritalCertificateNo: record[28],
			//MarriageDate:           record[],                       //?
			DivorceCertificate:   IfNotNullAvaliable(record[28]), //
			DivorceCertificateNo: record[32],                     //
			//DivorceCertificateDate: record[],
			FamilyStatus: record[19], //
			//MentalDisorders:        record[],
			//Disabilities:           record[],
			EducationStatus: EduCV(record[21]),
			JobTypeID:       StrJobToInt(db, record[23]), // ISI NULL DULU
			//NIKMother:              record[],
			Mother: record[32],
			//NIKFather:              record[],
			Father: record[31],
			//Coordinate:             record[],
		}

		fmt.Println("OIIIIIIIIIIIIII ", citizen.MentalDisorders)

		if err := db.Create(&citizen).Error; err != nil {
			logger.Log.Printf("ERROR NI DI INDEX KE %d VALUE YG DI INSERT %v", i, citizen)
		}

	}

}

func EduCV(string2 string) string {
	switch string2 {
	case "Tidak/Belum Sekolah":
		return entity.NotInSchool.ToString()
	case "Belum Tamat SD/Sederajat":
		return entity.NotFinishedElementary.ToString()
	case "Tamat SD/Sederajat":
		return entity.CompletedElementary.ToString()
	case "SLTP/Sederajat":
		return entity.JuniorHighSchool.ToString()
	case "SLTA/Sederajat":
		return entity.SeniorHighSchool.ToString()
	case "Diploma I/II":
		return entity.DiplomaIorII.ToString()
	case "Akademi/Diploma III/S. Muda":
		return entity.AcademyOrDiplomaIII.ToString()
	case "Diploma IV/Strata I":
		return entity.DiplomaIVOrBachelorsOrMasters.ToString()
	default:
		return entity.OtherEducation.ToString()
	}
}

func StrJobToInt(db *gorm.DB, string2 string) int {
	var result entity.Job
	_ = db.Where("name = ?", string2).First(&result)
	return result.ID
}

func StrProvinceToInt(db *gorm.DB, string2 string) int {
	var result entity.IndonesiaProvince
	_ = db.Where("name = ?", string2).First(&result)
	return int(result.ID)
}

func StrDistrictToInt(db *gorm.DB, string2 string) int {
	var result entity.IndonesiaDistrict
	_ = db.Where("name = ?", string2).First(&result)
	return int(result.ID)
}

func StrSubDistrictToInt(db *gorm.DB, string2 string) int {
	var result entity.IndonesiaSubDistrict
	_ = db.Where("name = ?", string2).First(&result)
	return int(result.ID)
}

func StrVillageToInt(db *gorm.DB, string2 string) int {
	var result entity.IndonesiaVillage
	_ = db.Where("name = ?", string2).First(&result)
	return int(result.ID)
}

func FormatGender(g string) string {
	if g == "Laki-laki" {
		return entity.Man.ToString()
	}

	return entity.Girl.ToString()
}

func IfNotNullAvaliable(t interface{}) string {
	if t == nil || t == "" {
		return entity.StatusTidakAda.ToString()
	}

	return entity.StatusAda.ToString()
}

func ReligionCv(r string) string {
	switch r {
	case "Islam":
		return entity.Islam.ToString()
	case "Kristen":
		return entity.Kristen.ToString()
	case "Hindu":
		return entity.Hindu.ToString()
	case "Katholik":
		return entity.Katholik.ToString()
	case "Buddha":
		return entity.Buddha.ToString()
	case "Kong Hu Cu":
		return entity.KongHuCu.ToString()
	default:
		return entity.Etc.ToString()
	}
}

func BlodCv(b string) string {
	switch b {
	case "A":
		return entity.A.ToString()
	case "B":
		return entity.B.ToString()
	case "AB":
		return entity.AB.ToString()
	case "O":
		return entity.O.ToString()
	case "A+":
		return entity.APositive.ToString()
	case "A-":
		return entity.ANegative.ToString()
	case "B+":
		return entity.BPositive.ToString()
	case "B-":
		return entity.BNegative.ToString()
	case "AB+":
		return entity.ABPositive.ToString()
	case "AB-":
		return entity.ABNegative.ToString()
	default:
		return entity.BloodTypeUnknown.ToString()
	}
}

func MaritalCv(b string) string {
	switch b {
	case "KAWIN":
		return entity.MaritalStatusRegistered.ToString()
	case "BELUM KAWIN":
		return entity.MaritalStatusUnmarried.ToString()
	case "CERAI HIDUP":
		return entity.MaritalStatusRegisteredDivorce.ToString()
	case "CERAI MATI":
		return entity.MaritalStatusWidowed.ToString()
	default:
		return ""

	}
}

func TestJobs(*testing.T) {
	ctx := context.Background()
	db, _ := cfg.GetPool(cfg.GetConfig()) // Koneksi ke database
	jobsRepository := repository.NewJobsRepository(db)

	file, err := os.Open("apdk_mento_tocsv.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	pkcount := 1
	for i, record := range records {

		if i == 1 {
			continue
		}

		job := entity.Job{
			Code: fmt.Sprintf("PK%d", pkcount),
			Name: record[23],
		}

		var existingJob entity.Job
		err := db.Where("name = ?", job.Name).First(&existingJob).Error
		if err == nil {
			fmt.Printf("Job with name %s already exists, skipping insert...\n", job.Name)
			continue
		}

		err = jobsRepository.CreateJobs(ctx, job)
		if err != nil {
			fmt.Println("Error inserting job:", err)
		} else {
			fmt.Printf("Job %s inserted successfully.\n", job.Name)
			pkcount++
		}
	}
}
