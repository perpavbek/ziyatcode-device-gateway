package mapper

import (
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/application/dto"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/valueobject"
)

type ScannedDeviceMapper struct{}

func NewScannedDeviceMapper() *ScannedDeviceMapper {
	return &ScannedDeviceMapper{}
}

func (m *ScannedDeviceMapper) FromEntityToResponseDto(d *valueobject.ScannedDevice) dto.ScannedDeviceResponseDto {
	return dto.ScannedDeviceResponseDto{
		UUID: d.UUID,
		Name: d.Name,
		Type: string(d.Type),
	}
}

func (m *ScannedDeviceMapper) FromEntitiesToResponseDtos(devices *[]valueobject.ScannedDevice) []dto.ScannedDeviceResponseDto {
	res := make([]dto.ScannedDeviceResponseDto, len(*devices))
	for i, d := range *devices {
		res[i] = m.FromEntityToResponseDto(&d)
	}
	return res
}