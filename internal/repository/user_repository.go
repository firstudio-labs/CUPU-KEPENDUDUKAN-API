package repository

import (
	"context"
	"fmt"
	"github.com/firstudio-lab/JARITMAS-API/internal/entity"
	"github.com/firstudio-lab/JARITMAS-API/pkg/logger"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserByUsername(ctx context.Context, username string) (entity.User, error)
	CreateUser(ctx context.Context, newUser entity.User) error
}

type UserRepositoryImpl struct {
	*gorm.DB
}

func NewUserRepositoryImpl(DB *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: DB}
}

func (r UserRepositoryImpl) FindUserByUsername(ctx context.Context, username string) (entity.User, error) {

	var user entity.User
	if err := r.DB.WithContext(ctx).Select("username", "password").Where("username = ?", username).First(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r UserRepositoryImpl) CreateUser(ctx context.Context, newUser entity.User) error {
	var exist entity.User
	if err := r.DB.WithContext(ctx).Where("username = ?", newUser.Username).First(&exist).Error; err != nil {
		logger.Log.Debugf("QUERY Error %v", err)
	}

	if exist.ID != 0 {
		return fmt.Errorf("username %s has been registered", exist.Username)
	}

	if err := r.DB.WithContext(ctx).Create(&newUser).Error; err != nil {
		logger.Log.Debugf("QUERY Error %v", err)
	}

	return nil
}
