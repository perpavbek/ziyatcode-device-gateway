package model

import deviceEnum "github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/enum"

type FirmwareModel struct {
	Id         uint64			 	 `gorm:"primarykey"`
	Name       string				 `gorm:"type:text;not null"`
	Version    string				 `gorm:"type:text;not null"`
	Checksum   string 				 `gorm:"type:text;not null"`
	DeviceType deviceEnum.DeviceType `gorm:"type:text;not null;check:device_type IN ('esp32')"`
}