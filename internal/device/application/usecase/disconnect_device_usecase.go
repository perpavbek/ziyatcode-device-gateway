package usecase

import (
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/exception"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/port"
)

type DisconnectDeviceUseCase struct {
	repo port.DeviceRepositoryPort
	connector port.DeviceConnectionPort
	state port.DeviceStatePort
}

func NewDisconnectDeviceUseCase(
	repo port.DeviceRepositoryPort,
	connector port.DeviceConnectionPort,
	state port.DeviceStatePort) *DisconnectDeviceUseCase{
	
	return &DisconnectDeviceUseCase{
		repo: repo,
		connector: connector,
		state: state,
	}
}

func (uc DisconnectDeviceUseCase) Execute(UUID string) error{
	device, err := uc.repo.FindByUuid(UUID)
	if err != nil{
		return err
	}

	if device == nil{
		return &exception.DeviceNotFoundError{UUID: UUID}
	}

	uc.connector.Disconnect(device)
	uc.state.Delete(UUID)
	uc.repo.Delete(UUID)
	return nil
}