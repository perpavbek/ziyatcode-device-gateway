package entity

import deviceEnum "github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/enum"

type Firmware struct {
	id         *uint64
	name       string
	version    string
	checksum   string
	deviceType deviceEnum.DeviceType
}

func NewFirmware(
	id *uint64,
	name string,
	version string,
	checksum string,
	deviceType deviceEnum.DeviceType) *Firmware {
	return &Firmware{
		id:         id,
		name:       name,
		version:    version,
		checksum:   checksum,
		deviceType: deviceType,
	}
}

func (f *Firmware) Id() uint64 {
	return *f.id
}

func (f *Firmware) Name() string {
	return f.name
}

func (f *Firmware) Version() string {
	return f.version
}

func (f *Firmware) Checksum() string {
	return f.checksum
}

func (f *Firmware) DeviceType() deviceEnum.DeviceType {
	return f.deviceType
}

func (f *Firmware) Rename(name string) {
	f.name = name
}
