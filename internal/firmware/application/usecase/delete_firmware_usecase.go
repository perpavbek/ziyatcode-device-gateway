package usecase

import (
	"github.com/perpavbek/ziyatcode-device-gateway/internal/firmware/domain/exception"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/firmware/domain/port"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/firmware/utils"
)

type DeleteFirmwareUseCase struct {
	repo port.FirmwareRepositoryPort
	storage port.FirmwareStoragePort
}

func NewDeleteFirmwareUseCase(
	repo port.FirmwareRepositoryPort,
	storage port.FirmwareStoragePort) *DeleteFirmwareUseCase{
	return &DeleteFirmwareUseCase{
		repo: repo,
		storage: storage,
	}
}

func (uc *DeleteFirmwareUseCase) Execute(id uint64) error{
	firmware, err := uc.repo.FindById(id)

	if err != nil{
		return err
	}

	if firmware == nil{
		return &exception.FirmwareNotFoundError{Id: id}
	}

	err = uc.storage.Delete(utils.ResolveFirmwareStorageKey(firmware))
	
	if err != nil{
		return err
	}
	
	err = uc.repo.Delete(id)

	if err != nil{
		return err
	}

	return nil
}