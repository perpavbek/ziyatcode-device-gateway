package usecase

import (
	"time"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/port"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/valueobject"
)

type ScanDevicesUseCase struct {
	discovery port.DeviceDiscoveryPort
}

func NewScanDevicesUseCase(discovery port.DeviceDiscoveryPort) *ScanDevicesUseCase{
	return &ScanDevicesUseCase{
		discovery: discovery,
	}
}

func (uc *ScanDevicesUseCase) Execute(timeout *time.Duration) (*valueobject.ScannedDeviceBatch, error){
	batch, err := uc.discovery.Scan(timeout)
	
	if err != nil{
		return nil, err
	}

	return batch, nil
}