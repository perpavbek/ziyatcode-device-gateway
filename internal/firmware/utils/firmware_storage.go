package utils

import (
	"fmt"

	"github.com/perpavbek/ziyatcode-device-gateway/internal/firmware/domain/entity"
)

func ResolveFirmwareStorageKey(firmware *entity.Firmware) string{
	return fmt.Sprintf(`%s_%s`, firmware.Name(), firmware.Version())
}