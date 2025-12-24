package port

import "github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/entity"

type DeviceRepositoryPort interface {
	Create(device *entity.Device) (*entity.Device, error)
	FindByUuid(uuid string) (*entity.Device, error)
	FindAll() ([]*entity.Device, error)
	Update(device entity.Device) (*entity.Device, error)
	Delete(uuid string) (error)
}