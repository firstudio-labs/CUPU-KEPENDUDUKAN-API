package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/firstudio-lab/JARITMAS-API/internal/entity"
	"github.com/firstudio-lab/JARITMAS-API/pkg/logger"
	"gorm.io/gorm"
)

type CitizensRepository interface {
	CreateCitizen(ctx context.Context, tx *gorm.DB, citizen entity.Citizen) error
	GetCitizenByNIK(ctx context.Context, tx *gorm.DB, nik int64) (entity.Citizen, error)
	GetAllCitizenPerPage(ctx context.Context, tx *gorm.DB, page int) ([]entity.Citizen, error)
	UpdateCitizen(ctx context.Context, tx *gorm.DB, nik int64, citizenUpdate entity.Citizen) error
	DeleteCitizenByNIK(ctx context.Context, tx *gorm.DB, nik int64) error
}

type CitizensRepositoryImpl struct {
}

func NewCitizensRepository() *CitizensRepositoryImpl {
	return &CitizensRepositoryImpl{}
}

func (r CitizensRepositoryImpl) CreateCitizen(ctx context.Context, tx *gorm.DB, citizen entity.Citizen) error {
	var existingCitizen entity.Citizen
	result := tx.Where("nik = ?", citizen.NIK).First(&existingCitizen)

	// If record found, it means NIK already exists
	if result.RowsAffected > 0 {
		return errors.New("citizen with this NIK already exists")
	}

	// If NIK doesn't exist, create new citizen
	if err := tx.Create(&citizen).Error; err != nil {
		return err
	}

	return nil
}

func (r CitizensRepositoryImpl) GetCitizenByNIK(ctx context.Context, tx *gorm.DB, nik int64) (entity.Citizen, error) {
	// CHECK NIK EXIST KAH?
	var result entity.Citizen
	if err := tx.WithContext(ctx).Where("nik = ?", nik).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, fmt.Errorf("citizens with nik = %d not found", nik)
		}
		logger.Log.Debugf("QUERY Error %v", err)
		return result, fmt.Errorf("internal error please try again later")
	}

	return result, nil
}

func (r CitizensRepositoryImpl) GetAllCitizenPerPage(ctx context.Context, tx *gorm.DB, page int) ([]entity.Citizen, error) {
	var results []entity.Citizen
	perPage := 10
	offset := (page - 1) * perPage

	if err := tx.WithContext(ctx).
		Limit(perPage). // Batasi jumlah item per halaman
		Offset(offset). // Mulai dari offset yang dihitung
		Find(&results). // Ambil data
		Error; err != nil {
		logger.Log.Errorf("QUERY Error %v", err)
		return nil, fmt.Errorf("data not ready yet")
	}

	return results, nil
}

func (r CitizensRepositoryImpl) UpdateCitizen(ctx context.Context, tx *gorm.DB, nik int64, citizenUpdate entity.Citizen) error {
	// First, check if the citizen exists
	var existingCitizen entity.Citizen
	result := tx.WithContext(ctx).Where("nik = ?", nik).First(&existingCitizen)

	// If no record is found, return the custom not found error
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("failed update nik %d notfound", nik)
		}
		return fmt.Errorf("Other error")
	}

	if err := tx.WithContext(ctx).Where("nik = ?", nik).Updates(citizenUpdate).Error; err != nil {
		logger.Log.Errorf("QUERY Error %v", err)
		return err
	}

	return nil
}

func (r CitizensRepositoryImpl) DeleteCitizenByNIK(ctx context.Context, tx *gorm.DB, nik int64) error {
	// Cek apakah citizen dengan NIK tersebut ada
	var citizen entity.Citizen
	if err := tx.WithContext(ctx).Where("nik = ?", nik).First(&citizen).Error; err != nil {
		// Jika error yang terjadi adalah 'record not found', return error custom
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("citizen with NIK %d not found", nik)
		}
		// Jika ada error lain, log dan kembalikan error tersebut
		logger.Log.Errorf("QUERY Error while checking citizen existence: %v", err)
		return err
	}

	// Jika citizen ditemukan, lanjutkan untuk menghapusnya
	if err := tx.WithContext(ctx).Where("nik = ?", nik).Delete(&entity.Citizen{}).Error; err != nil {
		logger.Log.Errorf("QUERY Error while deleting citizen: %v", err)
		return err
	}

	return nil
}
