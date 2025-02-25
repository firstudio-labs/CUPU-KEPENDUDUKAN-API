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
	ExistCitizenNIK(ctx context.Context, tx *gorm.DB, nik int64) error
}

type CitizensRepositoryImpl struct {
}

func NewCitizensRepository() *CitizensRepositoryImpl {
	return &CitizensRepositoryImpl{}
}

func (r CitizensRepositoryImpl) CreateCitizen(ctx context.Context, tx *gorm.DB, citizen entity.Citizen) error {
	if err := tx.WithContext(ctx).Create(&citizen).Error; err != nil {
		logger.Log.Debugf("QUERY Error %v", err)
		return fmt.Errorf("internal error please try again later")
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
	if err := tx.WithContext(ctx).Where("nik = ?", nik).Updates(citizenUpdate).Error; err != nil {
		logger.Log.Errorf("QUERY Error %v", err)
		return err
	}

	return nil
}

func (r CitizensRepositoryImpl) DeleteCitizenByNIK(ctx context.Context, tx *gorm.DB, nik int64) error {
	if err := tx.WithContext(ctx).Where("nik = ?", nik).Delete(&entity.Citizen{}).Error; err != nil {
		logger.Log.Errorf("QUERY Error %v", err)
		return err
	}

	return nil
}

func (r CitizensRepositoryImpl) ExistCitizenNIK(ctx context.Context, tx *gorm.DB, nik int64) error {
	// find onlu oen field biar cepet
	var exists bool
	if err := tx.WithContext(ctx).Table("citizens").Select("1").Where("nik = ?", nik).
		Limit(1).Scan(&exists).Error; err != nil {
		logger.Log.Errorf("QUERY Error %v", err)
		return fmt.Errorf("failed to check NIK existence: %w", err)
	}

	if !exists {
		//if data not found return error
		return fmt.Errorf("citizen with NIK %d not found", nik)
	}

	return nil
}
