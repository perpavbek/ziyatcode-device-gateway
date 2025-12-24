package usecase

import (
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/entity"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/exception"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/port"
)

type GetDeviceUseCase struct {
	repo port.DeviceRepositoryPort
}

func NewGetDeviceUseCase(repo port.DeviceRepositoryPort) *GetDeviceUseCase{
	return &GetDeviceUseCase{
		repo: repo,
	}
}

func (uc *GetDeviceUseCase) Execute(uuid string) (*entity.Device, error){
	device, err := uc.repo.FindByUuid(uuid)
	
	if err != nil{
		return nil, err
	}

	if device == nil{
		return nil, &exception.DeviceNotFoundError{UUID: uuid}
	}

	return device, nil
}