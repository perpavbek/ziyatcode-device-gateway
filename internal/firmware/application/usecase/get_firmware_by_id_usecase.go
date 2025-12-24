package usecase

import (
	"github.com/perpavbek/ziyatcode-device-gateway/internal/firmware/domain/entity"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/firmware/domain/port"
)

type GetFirmwareByIdUseCase struct {
	repo port.FirmwareRepositoryPort
}

func NewGetFirmwareByIdUseCase(repo port.FirmwareRepositoryPort) *GetFirmwareByIdUseCase{
	return &GetFirmwareByIdUseCase{
		repo: repo,
	}
}

func (uc *GetFirmwareByIdUseCase) Execute(id uint64) (*entity.Firmware, error){
	firmware, err := uc.repo.FindById(id)

	if err != nil{
		return nil, err
	}

	return firmware, nil
}