package device

type DeviceType string
type DeviceStatus string

const (
	DeviceTypeESP32   DeviceType = "esp32"

	DeviceStatusOnline  DeviceStatus = "online"
	DeviceStatusOffline DeviceStatus = "offline"
)