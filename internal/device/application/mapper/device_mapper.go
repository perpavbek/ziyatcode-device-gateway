package mapper

import (
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/application/dto"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/entity"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/valueobject"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/utils"
)

type DeviceMapper struct{}

func NewDeviceMapper() *DeviceMapper {
	return &DeviceMapper{}
}

func (m *DeviceMapper) FromEntityToResponseDto(d *entity.Device) dto.DeviceResponseDto {
	return dto.DeviceResponseDto{
		UUID: d.UUID(),
		Name: d.Name(),
		Type: string(d.Type()),
	}
}

func (m *DeviceMapper) FromEntitiesToResponseDtos(devices []*entity.Device) []dto.DeviceResponseDto {
	res := make([]dto.DeviceResponseDto, len(devices))
	for i, d := range devices {
		res[i] = m.FromEntityToResponseDto(d)
	}
	return res
}

func (m *DeviceMapper) FromScannedDeviceToEntity(scanned valueobject.ScannedDevice) *entity.Device {
	dt, err := utils.ParseDeviceType(scanned.Type)
	if err != nil {

	}
	return entity.NewDevice(
		scanned.UUID,
		scanned.Name,
		dt,
		scanned.Address,
	)
}
