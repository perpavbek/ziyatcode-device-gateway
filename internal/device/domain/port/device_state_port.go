package port

import "github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/valueobject"

type DeviceStatePort interface {
	UpdateOnline(uuid string) error
	Delete(uuid string) error
	GetStatus(uuid string) (valueobject.DeviceState, error)
}