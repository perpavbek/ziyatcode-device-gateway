package mapper

import (
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/entity"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/infrastructure/model"
)

type DeviceModelMapper struct{}

func NewDeviceModelMapper() *DeviceModelMapper {
	return &DeviceModelMapper{}
}

func (m *DeviceModelMapper) ToEntity(model *model.DeviceModel) *entity.Device {
	return entity.NewDevice(
		model.UUID,
		model.Name,
		model.DType,
		model.Address)
}

func (m *DeviceModelMapper) ToModel(device *entity.Device) *model.DeviceModel {
	return &model.DeviceModel{
		UUID:    device.UUID(),
		Name:    device.Name(),
		DType:   device.Type(),
		Address: device.Address(),
	}
}
