package device

type Device struct {
	uuid    string
	name    string
	dtype   DeviceType
	address string
}

func New(
	uuid string,
	name string,
	dtype DeviceType,
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

func (d *Device) Type() DeviceType {
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