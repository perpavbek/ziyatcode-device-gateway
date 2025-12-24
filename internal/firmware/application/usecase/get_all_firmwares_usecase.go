package usecase

import (
	"github.com/perpavbek/ziyatcode-device-gateway/internal/firmware/domain/entity"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/firmware/domain/port"
)

type GetAllFirmwaresUseCase struct {
	repo port.FirmwareRepositoryPort
}

func NewGetAllFirmwaresUseCase(repo port.FirmwareRepositoryPort) *GetAllFirmwaresUseCase{
	return &GetAllFirmwaresUseCase{
		repo: repo,
	}
}

func (uc *GetAllFirmwaresUseCase) Execute() ([]*entity.Firmware, error){
	devices, err := uc.repo.FindAll()

	if err != nil{
		return nil, err
	}

	return devices, nil
}