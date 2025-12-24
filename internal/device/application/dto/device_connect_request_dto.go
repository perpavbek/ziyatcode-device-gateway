package dto

type DeviceConnectRequestDto struct {
	UUID string `json:"uuid"`
	Password *string `json:"password"`
}