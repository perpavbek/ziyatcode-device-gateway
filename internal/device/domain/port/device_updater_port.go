package port

import "github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/entity"

type DeviceUpdaterPort interface {
	Update(device *entity.Device, binary []byte) error
}