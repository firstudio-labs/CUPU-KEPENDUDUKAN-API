package cfg

import (
	"context"
	"github.com/firstudio-lab/JARITMAS-API/internal/dto"
	"github.com/firstudio-lab/JARITMAS-API/internal/entity"
	"github.com/firstudio-lab/JARITMAS-API/internal/repository"
	"github.com/firstudio-lab/JARITMAS-API/internal/usecase"
	"github.com/firstudio-lab/JARITMAS-API/pkg/logger"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func SeedingUserAdmin(conn *gorm.DB) error {
	repo := repository.NewUserRepository(conn)
	getMethod := usecase.NewAuthUsecase(repo, validator.New())

	if err := getMethod.Register(context.Background(), dto.RegisterRequest{
		Username:        "koriebruh",
		Password:        "koriebruh",
		ConfirmPassword: "koriebruh",
	}); err != nil {
		logger.Log.Errorf("FAILED TO SEED DATA")
		return err
	}

	return nil
}

func SeedingSHDK(conn *gorm.DB) error {
	familyStatuses := []entity.FamilyStatus{
		{ID: 1, Name: "Anak"},
		{ID: 2, Name: "Ibu"},
		{ID: 3, Name: "Ayah"},
		{ID: 4, Name: "Kakek"},
		{ID: 5, Name: "Nenek"},
	}
	if tx := conn.Create(&familyStatuses).Error; tx != nil {
		return tx
	}

	return nil
}

func SeedingJobs(conn *gorm.DB) error {
	jobs := []entity.Job{
		{ID: 1, Name: "Dokter"},
		{ID: 2, Name: "Guru"},
		{ID: 3, Name: "Wiraswasta"},
		{ID: 4, Name: "Polisi"},
		{ID: 5, Name: "Petani"},
	}

	if tx := conn.Create(&jobs).Error; tx != nil {
		return tx
	}

	return nil
}

func Province(conn *gorm.DB) error {
	provinces := []entity.Province{
		{ID: 1, ProvinceCode: "01", Name: "Aceh", Coordinate: "4.5, 96.9"},
		{ID: 2, ProvinceCode: "02", Name: "Bali", Coordinate: "-8.4095, 115.1889"},
		{ID: 3, ProvinceCode: "03", Name: "Jawa Barat", Coordinate: "-6.993, 107.631"},
		{ID: 4, ProvinceCode: "04", Name: "Yogyakarta", Coordinate: "-7.7956, 110.3695"},
		{ID: 5, ProvinceCode: "05", Name: "Sumatera Utara", Coordinate: "3.595, 98.675"},
	}

	if tx := conn.Create(&provinces).Error; tx != nil {
		return tx
	}

	return nil
}
func District(conn *gorm.DB) error {
	districts := []entity.District{
		{ID: 1, DistrictCode: "101", Name: "Medan", Coordinate: "3.5952, 98.6722"},
		{ID: 2, DistrictCode: "102", Name: "Denpasar", Coordinate: "-8.6589, 115.2167"},
		{ID: 3, DistrictCode: "103", Name: "Bandung", Coordinate: "-6.9175, 107.6191"},
		{ID: 4, DistrictCode: "104", Name: "Sleman", Coordinate: "-7.639, 110.392"},
		{ID: 5, DistrictCode: "105", Name: "Pematangsiantar", Coordinate: "2.9758, 99.0705"},
	}

	if tx := conn.Create(&districts).Error; tx != nil {
		return tx
	}

	return nil

}
func SubDistrict(conn *gorm.DB) error {
	subDistricts := []entity.SubDistrict{
		{ID: 1, SubDistrictCode: "201", Name: "Medan Kota", Coordinate: "3.5897, 98.6727"},
		{ID: 2, SubDistrictCode: "202", Name: "Denpasar Barat", Coordinate: "-8.6795, 115.2163"},
		{ID: 3, SubDistrictCode: "203", Name: "Bandung Kulon", Coordinate: "-6.9283, 107.6305"},
		{ID: 4, SubDistrictCode: "204", Name: "Sleman Tengah", Coordinate: "-7.6753, 110.4193"},
		{ID: 5, SubDistrictCode: "205", Name: "Pematangsiantar Barat", Coordinate: "2.9736, 99.0735"},
	}

	if tx := conn.Create(&subDistricts).Error; tx != nil {
		return tx
	}

	return nil
}
func Village(conn *gorm.DB) error {
	villages := []entity.Village{
		{ID: 1, VillageCode: "301", Name: "Desa Medan", Coordinate: "3.5932, 98.6759"},
		{ID: 2, VillageCode: "302", Name: "Desa Denpasar", Coordinate: "-8.6632, 115.2181"},
		{ID: 3, VillageCode: "303", Name: "Desa Bandung", Coordinate: "-6.9287, 107.6161"},
		{ID: 4, VillageCode: "304", Name: "Desa Sleman", Coordinate: "-7.6293, 110.3959"},
		{ID: 5, VillageCode: "305", Name: "Desa Pematangsiantar", Coordinate: "2.9731, 99.0719"},
	}

	if tx := conn.Create(&villages).Error; tx != nil {
		return tx
	}

	return nil
}
