package dto

type FirmwareResponseDto struct{
    Id uint64 `json:"id"`
    Name string `json:"name"`
    Version string `json:"version"`
    Checksum string `json:"checksum"`
    DeviceType string `json:"deviceType"`
}