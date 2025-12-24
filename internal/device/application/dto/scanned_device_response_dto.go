package dto

type ScannedDeviceResponseDto struct{
	UUID string `json:"uuid"`
	ScanTime string `json:"scanTime"`
	Name string `json:"name"`
	Type string `json:"type"`
	Address string `json:"address"`
}