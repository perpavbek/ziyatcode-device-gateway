package usecase

import (
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/entity"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/port"
)

type GetAllConnectedDevicesUseCase struct {
	repo port.DeviceRepositoryPort
}

func NewGetAllConnectedDevicesUseCase(repo port.DeviceRepositoryPort) *GetAllConnectedDevicesUseCase{
	return &GetAllConnectedDevicesUseCase{
		repo: repo,
	}
}

func (uc *GetAllConnectedDevicesUseCase) Execute() ([]*entity.Device, error){
	devices, err := uc.repo.FindAll()
	if err != nil{
		return nil, err
	}
	return devices, nil
}