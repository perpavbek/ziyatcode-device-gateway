package usecase

import (
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/port"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/valueobject"
)

type GetScannedDevicesUseCase struct {
	discovery port.DeviceDiscoveryPort
}

func NewGetScannedDevicesUseCase(discovery port.DeviceDiscoveryPort) *GetScannedDevicesUseCase{
	return &GetScannedDevicesUseCase{
		discovery: discovery,
	}
}

func (uc *GetScannedDevicesUseCase) Execute() (*valueobject.ScannedDeviceBatch, error){
	batch, err := uc.discovery.GetScanned()

	if err != nil{
		return nil, err
	}

	return batch, nil
}