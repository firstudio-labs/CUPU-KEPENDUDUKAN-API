package cfg

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
	"time"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
	dbErr      error
)

func GetPool(config *Config) (*gorm.DB, error) {
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.DataBase.User,
			config.DataBase.Pass,
			config.DataBase.Host,
			config.DataBase.Port,
			config.DataBase.Name,
		)

		//OPEN CONNECTION
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			dbErr = fmt.Errorf("failed to connect to the database: %v", err)
		}

		// AUTO MIGRATE
		//if err = db.AutoMigrate(
		//	&entity.IndonesiaProvince{},
		//	&entity.IndonesiaDistrict{},
		//	&entity.IndonesiaSubDistrict{},
		//	&entity.IndonesiaVillage{},
		//	&entity.Job{},
		//
		//	//MAIN
		//	&entity.Citizen{},
		//	&entity.Relocation{},
		//	&entity.Approved{},
		//); err != nil {
		//	err = fmt.Errorf("failed auto migrate bcs %e", err)
		//}

		// SET CONNECTION POOL
		sqlPool, err := db.DB()
		if err != nil {
			dbErr = fmt.Errorf("failed to get instance conn %v", err)
		}
		sqlPool.SetMaxOpenConns(60)
		sqlPool.SetMaxIdleConns(30)
		sqlPool.SetConnMaxLifetime(60 * time.Minute)

		dbInstance = db
	})

	///// CALL SEEDING
	//// if we run include
	//if os.Getenv("SEED_DATA") == "true" {
	//	_ = SeedingUserAdmin(db)
	//	_ = SeedingJobs(db)
	//	_ = SeedingSHDK(db)
	//
	//	// INI NANTI DI UBAH PKAI DATA EXEL
	//	_ = Province(db)
	//	_ = District(db)
	//	_ = SubDistrict(db)
	//	_ = Village(db)
	//
	//	log2.Log.Debug("SUCCESS TO SEED DATA")
	//}

	return dbInstance, nil
}
