package cfg

import (
	"fmt"
	"github.com/firstudio-lab/JARITMAS-API/internal/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func GetPool(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DataBase.User,
		config.DataBase.Pass,
		config.DataBase.Host,
		config.DataBase.Port,
		config.DataBase.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	if err = db.AutoMigrate(
		&entity.User{},
	); err != nil {
		return nil, fmt.Errorf("failed auto migrate bcs %e", err)
	}

	sqlPool, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to setup coonnection pool %e", err)
	}
	sqlPool.SetMaxOpenConns(60)
	sqlPool.SetMaxIdleConns(30)
	sqlPool.SetConnMaxLifetime(60 * time.Minute)

	return db, nil
}
