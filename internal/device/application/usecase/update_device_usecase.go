package usecase

import "github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/port"

type UpdateDeviceUseCase struct {
	deviceRepo port.DeviceRepositoryPort
}