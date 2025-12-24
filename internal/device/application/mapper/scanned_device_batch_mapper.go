package mapper

import (
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/application/dto"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/valueobject"
)

type ScannedDeviceBatchMapper struct{
	scannedDeviceMapper *ScannedDeviceMapper
}

func NewScannedDeviceBatchMapper(scannedDeviceMapper *ScannedDeviceMapper) *ScannedDeviceBatchMapper {
	return &ScannedDeviceBatchMapper{
		scannedDeviceMapper: scannedDeviceMapper,
	}
}

func (m *ScannedDeviceBatchMapper) FromEntityToResponseDto(batch *valueobject.ScannedDeviceBatch) dto.ScannedDeviceBatchResponseDto {
	return dto.ScannedDeviceBatchResponseDto{
		ScanTime: batch.ScanTime.String(),
		Devices: m.scannedDeviceMapper.FromEntitiesToResponseDtos(batch.Devices),
	}
}