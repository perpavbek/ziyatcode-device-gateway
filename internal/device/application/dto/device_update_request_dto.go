package dto

type DeviceUpdateRequestDto struct{
	DeviceUUID string `json:"deviceUuid"`
	FirmwareId string `json:"firmwareId"`
}