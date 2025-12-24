package usecase

import (
	"io"
	"mime/multipart"

	"github.com/perpavbek/ziyatcode-device-gateway/internal/firmware/application/dto"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/firmware/application/mapper"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/firmware/domain/entity"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/firmware/domain/port"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/firmware/utils"
)

type UploadFirmwareUseCase struct {
	repo port.FirmwareRepositoryPort
	storage port.FirmwareStoragePort
	mapper *mapper.FirmwareMapper
}

func NewUploadFirmwareUseCase(
	repo port.FirmwareRepositoryPort,
	storage port.FirmwareStoragePort,
	mapper *mapper.FirmwareMapper) *UploadFirmwareUseCase{
	return &UploadFirmwareUseCase{
		repo: repo,
		storage: storage,
	}
}

func (uc *UploadFirmwareUseCase) Execute(dto dto.FirmwareUploadDto, file multipart.File) (*entity.Firmware, error){
	defer file.Close()
	
	data, err := io.ReadAll(file)
    if err != nil {
        return nil, err
    }

	uploaded, err := uc.mapper.FromUploadDtoToEntity(&dto)
	
	if err != nil{
		return nil, err
	}

	err = uc.storage.Save(utils.ResolveFirmwareStorageKey(uploaded), data)

	if err != nil{
		return nil, err
	}

	firmware, err := uc.repo.Create(*uploaded)

	if err != nil{
		return nil, err
	}

	return firmware, nil
}