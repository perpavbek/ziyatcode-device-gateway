package exception

import "fmt"

type DeviceNotFoundError struct {
	UUID string
}

func (e *DeviceNotFoundError) Error() string {
	return fmt.Sprintf("Device with uuid %s not found", e.UUID)
}