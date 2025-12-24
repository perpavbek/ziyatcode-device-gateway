package usecase

import (
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/application/dto"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/application/mapper"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/entity"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/exception"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/port"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/valueobject"
)

type ConnectDeviceUseCase struct {
	repo port.DeviceRepositoryPort
	connector port.DeviceConnectionPort
	state port.DeviceStatePort
	discovery port.DeviceDiscoveryPort
	deviceMapper *mapper.DeviceMapper
}

func NewConnectDeviceUseCase(
	repo port.DeviceRepositoryPort,
	connector port.DeviceConnectionPort,
	state port.DeviceStatePort,
	discovery port.DeviceDiscoveryPort,
	deviceMapper *mapper.DeviceMapper) *ConnectDeviceUseCase {
	
	return &ConnectDeviceUseCase{
		repo: repo,
		connector: connector,
		state: state,
		discovery: discovery,
		deviceMapper: deviceMapper,
	}
}

func (uc *ConnectDeviceUseCase) Execute(dto dto.DeviceConnectRequestDto) (*entity.Device, error){
	scanned, err := uc.discovery.GetScanned();

	if err != nil{
		return nil, err
	}

	var dev *valueobject.ScannedDevice;

	for _, d := range *scanned.Devices{
		if d.UUID == dto.UUID{
			dev = &d
		}
	}

	if dev == nil{
		return nil, &exception.DeviceNotFoundError{UUID: dto.UUID}
	}

	existing, err := uc.repo.FindByUuid(dto.UUID)

	if err != nil{
		return nil, err
	}

	if existing != nil{
		updated, err := uc.repo.Update(*uc.deviceMapper.FromScannedDeviceToEntity(*dev))

		if err != nil{
			return nil, err
		}

		return updated, nil;
	}

	ok, err := uc.connector.Connect(*dev)

	if err != nil{
		return nil, err
	}

	if(!ok){
		return nil, err
	}

	device := uc.deviceMapper.FromScannedDeviceToEntity(*dev)
	saved, err := uc.repo.Create(device)
	
	if err != nil{
		return nil, err
	}

	uc.state.UpdateOnline(saved.UUID())

	return saved, nil
}