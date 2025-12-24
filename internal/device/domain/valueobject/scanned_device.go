package valueobject

import "time"

type ScannedDevice struct {
	UUID     string
	Name     string
	Type     string
	Address  string
	ScanTime *time.Time
}