package dto

type FirmwareUploadDto struct{
	Name string `json:"name"`
	Version string `json:"version"`
	Checksum string `json:"checksum"`
	DeviceType string `json:"deviceType"`
}