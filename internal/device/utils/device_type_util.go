package utils

import (
	"fmt"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/enum"
)

func ParseDeviceType(strType string) (enum.DeviceType, error) {
	switch strType {
	case string(enum.DeviceTypeESP32):
		return enum.DeviceTypeESP32, nil
	default:
		return "", fmt.Errorf("unknown device type: %s", strType)
	}
}
