package mapper

import (
	deviceUtils "github.com/perpavbek/ziyatcode-device-gateway/internal/device/utils"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/firmware/application/dto"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/firmware/domain/entity"
)

type FirmwareMapper struct{}

func NewFirmwareMapper() *FirmwareMapper {
	return &FirmwareMapper{}
}

func (m *FirmwareMapper) FromEntityToResponseDto(firmware *entity.Firmware) dto.FirmwareResponseDto{
	return dto.FirmwareResponseDto{
		Id: firmware.Id(),
		Name: firmware.Name(),
		Version: firmware.Version(),
		Checksum: firmware.Checksum(),
		DeviceType: string(firmware.DeviceType()),
	}
}

func (m *FirmwareMapper) FromUploadDtoToEntity(dto *dto.FirmwareUploadDto) (*entity.Firmware, error){
	deviceType, err := deviceUtils.ParseDeviceType(dto.DeviceType)
	if err != nil{
		return nil, err
	}
	return entity.New(nil, dto.Name, dto.Version, dto.Checksum, deviceType), nil
}