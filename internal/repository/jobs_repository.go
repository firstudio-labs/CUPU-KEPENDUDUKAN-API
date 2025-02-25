package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/firstudio-lab/JARITMAS-API/internal/entity"
	"github.com/firstudio-lab/JARITMAS-API/pkg/logger"
	"gorm.io/gorm"
)

type JobsRepository interface {
	FindAllJobs(ctx context.Context) ([]entity.Job, error)
	CreateJobs(ctx context.Context, newJob entity.Job) error
	UpdateJobById(ctx context.Context, id int, UpdateJob entity.Job) error
	DeleteJobById(ctx context.Context, id int) error
	ExistJobCode(ctx context.Context, code string) error
}
type JobsRepositoryImpl struct {
	*gorm.DB
}

func NewJobsRepository(DB *gorm.DB) *JobsRepositoryImpl {
	return &JobsRepositoryImpl{DB: DB}
}

func (r JobsRepositoryImpl) FindAllJobs(ctx context.Context) ([]entity.Job, error) {
	var jobs []entity.Job
	if err := r.DB.WithContext(ctx).Find(&jobs).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return jobs, fmt.Errorf("failed to get job data is not found")
		}
		logger.Log.Debugf("QUERY Error %v", err)
		return jobs, fmt.Errorf("internal error please try again later")
	}

	return jobs, nil
}

func (r JobsRepositoryImpl) CreateJobs(ctx context.Context, newJob entity.Job) error {
	// Validasi apakah job dengan code yang sama sudah ada
	var existingJob entity.Job
	if err := r.DB.WithContext(ctx).Where("code = ?", newJob.Code).First(&existingJob).Error; err == nil {
		// Jika ditemukan job dengan code yang sama
		return fmt.Errorf("job dengan code %s sudah ada", newJob.Code)
	}
	if err := r.DB.WithContext(ctx).Create(&newJob).Error; err != nil {
		logger.Log.Debugf("QUERY Error %v", err)
		return fmt.Errorf("internal error, coba lagi nanti")
	}

	return nil
}

func (r JobsRepositoryImpl) UpdateJobById(ctx context.Context, id int, UpdateJob entity.Job) error {
	// Validasi apakah job dengan ID tersebut ada
	var job entity.Job
	if err := r.DB.WithContext(ctx).Where("id =?", id).First(&job).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Job tidak ditemukan
			return fmt.Errorf("job dengan ID %d tidak ditemukan", id)
		}
		// Terjadi error lain
		logger.Log.Debugf("QUERY Error %v", err)
		return fmt.Errorf("internal error, coba lagi nanti")
	}

	// Validasi apakah job dengan code yang baru sama dengan yang sudah ada
	var existingJob entity.Job
	if err := r.DB.WithContext(ctx).Where("code = ? AND id != ?", UpdateJob.Code, id).First(&existingJob).Error; err == nil {
		// Jika ditemukan job dengan code yang sama, tetapi ID berbeda
		return fmt.Errorf("job dengan code %s sudah ada", UpdateJob.Code)
	}

	// Melakukan update job jika ID ditemukan dan tidak ada duplikat code
	if err := r.DB.WithContext(ctx).Where("id =?", id).Updates(UpdateJob).Error; err != nil {
		logger.Log.Debugf("QUERY Error %v", err)
		return fmt.Errorf("internal error, coba lagi nanti")
	}

	return nil
}

func (r JobsRepositoryImpl) DeleteJobById(ctx context.Context, id int) error {
	// Validasi apakah job dengan ID tersebut ada
	var job entity.Job
	if err := r.DB.WithContext(ctx).Where("id =?", id).First(&job).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Job tidak ditemukan
			return fmt.Errorf("job dengan ID %d tidak ditemukan", id)
		}
		// Terjadi error lain
		logger.Log.Debugf("QUERY Error %v", err)
		return fmt.Errorf("internal error, coba lagi nanti")
	}

	// Menghapus job jika ID ditemukan
	if err := r.DB.WithContext(ctx).Where("id =?", id).Delete(&entity.Job{}).Error; err != nil {
		logger.Log.Debugf("QUERY Error %v", err)
		return fmt.Errorf("internal error, coba lagi nanti")
	}

	return nil
}

func (r JobsRepositoryImpl) ExistJobCode(ctx context.Context, code string) error {
	var exists bool
	if err := r.DB.WithContext(ctx).Table("jobs").Select("1").Where("id = ?", code).
		Limit(1).Scan(&exists).Error; err != nil {
		logger.Log.Errorf("QUERY Error %v", err)
		return fmt.Errorf("failed to check exist code existence: %w", err)
	}

	if !exists {
		//if data not found return error
		return fmt.Errorf("code %d already usage", code)
	}

	return nil
}
