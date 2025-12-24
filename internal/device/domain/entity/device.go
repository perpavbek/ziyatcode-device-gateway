package entity

import (
	"time"

	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/enum"
)

type Device struct {
	uuid     string
	name     string
	dtype    enum.DeviceType
	address  string
	status   enum.DeviceStatus
	lastSeen *time.Time
}

func NewDevice(
	uuid string,
	name string,
	dtype enum.DeviceType,
	address string,
) *Device {
	return &Device{
		uuid:    uuid,
		name:    name,
		dtype:   dtype,
		address: address,
	}
}

func (d *Device) UUID() string {
	return d.uuid
}

func (d *Device) Name() string {
	return d.name
}

func (d *Device) Type() enum.DeviceType {
	return d.dtype
}

func (d *Device) Address() string {
	return d.address
}

func (d *Device) UpdateAddress(address string) {
	d.address = address
}

func (d *Device) Rename(name string) {
	d.name = name
}
