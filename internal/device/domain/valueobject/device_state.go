package valueobject

import (
	"time"

	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/enum"
)

type DeviceState struct {
	Status   enum.DeviceStatus
	LastSeen *time.Time
}

func (s *DeviceState) UpdateStatus(status enum.DeviceStatus, seen *time.Time) {
	s.Status = status
	s.LastSeen = seen
}