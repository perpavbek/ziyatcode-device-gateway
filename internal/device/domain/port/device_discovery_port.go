package port

import (
	"time"

	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/valueobject"
)

type DeviceDiscoveryPort interface {
	Scan(timeout *time.Duration) (*valueobject.ScannedDeviceBatch, error)
	GetScanned() (*valueobject.ScannedDeviceBatch, error)
}