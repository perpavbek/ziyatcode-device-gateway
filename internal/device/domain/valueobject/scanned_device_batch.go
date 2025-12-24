package valueobject

import "time"

type ScannedDeviceBatch struct {
	ScanTime *time.Time
	Devices  *[]ScannedDevice
}
