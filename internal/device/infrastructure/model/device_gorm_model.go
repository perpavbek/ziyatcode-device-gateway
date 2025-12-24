package model

import (
	"time"

	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/enum"
)

type DeviceModel struct {
	UUID string 			 `gorm:"primarykey;type:text"`
	Name string				 `gorm:"type:text;not null"`
	DType enum.DeviceType 	 `gorm:"type:text;not null;check:d_type IN ('esp32')"`
	Address string 			 `gorm:"type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}