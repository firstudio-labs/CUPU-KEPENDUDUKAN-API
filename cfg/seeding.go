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
	repo := repository.NewUserRepositoryImpl(conn)
	getMethod := usecase.NewAuthUsecaseImpl(repo, validator.New())

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

func SeedingSHDK() error {
	panic(entity.FamilyStatus{})
}
func SeedingJobs() error {
	panic(entity.Job{})
}

func Province() error {
	panic("E")
}
func District() error {
	panic("E")
}
func SubDistrict() error {
	panic("E")
}
func Village() error {
	panic("E")
}
