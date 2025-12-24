package dto

type DeviceResponseDto struct {
	UUID string `json:"uuid"`
	Name string	`json:"name"`
	Type string `json:"type"`
	Address string `json:"address"`
}