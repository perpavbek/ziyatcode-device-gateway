package dto

type ScannedDeviceBatchResponseDto struct {
	ScanTime string `json:"scanTime"`
	Devices []ScannedDeviceResponseDto `json:"devices"`
}