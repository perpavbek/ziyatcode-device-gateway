package mapper

import (
	"github.com/perpavbek/ziyatcode-device-gateway/internal/firmware/domain/entity"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/firmware/infrastructure/model"
)

type FirmwareModelMapper struct{}

func NewFirmwareModelMapper() *FirmwareModelMapper {
	return &FirmwareModelMapper{}
}

func (m *FirmwareModelMapper) ToEntity(model *model.FirmwareModel) *entity.Firmware {
	return entity.NewFirmware(
		&model.Id, 
		model.Name, 
		model.Version, 
		model.Checksum, 
		model.DeviceType)
}

func (m *FirmwareModelMapper) ToModel(firmware *entity.Firmware) *model.FirmwareModel {
	return &model.FirmwareModel{
		Id: firmware.Id(),
		Name: firmware.Name(),
		Version: firmware.Version(),
		Checksum: firmware.Checksum(),
		DeviceType: firmware.DeviceType(),
	}
}
