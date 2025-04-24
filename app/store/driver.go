package store

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"kryptonim-interview/app/errs"
	"kryptonim-interview/config"
)

func InitDriver(cfg *config.App) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Warsaw",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		errs.FatalOnError(err, "[1312261226] cannot init store driver")
	}

	return db
}
