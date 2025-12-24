package exception

import "fmt"

type FirmwareNotFoundError struct {
	Id uint64
}

func (e *FirmwareNotFoundError) Error() string {
	return fmt.Sprintf("Firmware with id %d not found", e.Id)
}