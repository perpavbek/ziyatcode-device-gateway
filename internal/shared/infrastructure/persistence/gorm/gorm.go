package gormdb

import (
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/infrastructure/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.DeviceModel{})
	
	return db, nil
}