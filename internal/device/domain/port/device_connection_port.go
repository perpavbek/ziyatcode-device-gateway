package port

import (
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/entity"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/valueobject"
)

type DeviceConnectionPort interface {
	Connect(device valueobject.ScannedDevice) (bool, error)
	Disconnect(device *entity.Device) (bool, error)
}