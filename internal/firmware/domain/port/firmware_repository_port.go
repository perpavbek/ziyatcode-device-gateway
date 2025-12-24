package port

import "github.com/perpavbek/ziyatcode-device-gateway/internal/firmware/domain/entity"

type FirmwareRepositoryPort interface {
	Create(firmware entity.Firmware) (*entity.Firmware, error)
	FindById(id uint64) (*entity.Firmware, error)
	FindAll() ([]*entity.Firmware, error)
	Update(device entity.Firmware) (*entity.Firmware, error)
	Delete(id uint64) (error)
}