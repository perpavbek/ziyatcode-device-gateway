package repository

import (
	"errors"

	"github.com/perpavbek/ziyatcode-device-gateway/internal/firmware/domain/entity"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/firmware/infrastructure/mapper"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/firmware/infrastructure/model"
	"gorm.io/gorm"
)

type GormFirmwareRepository struct {
	db *gorm.DB
	mapper *mapper.FirmwareModelMapper
}

func NewGormDeviceRepository(db *gorm.DB, mapper *mapper.FirmwareModelMapper) *GormFirmwareRepository{
	return &GormFirmwareRepository{
		db:db,
		mapper: mapper,
	}
}

func (repo *GormFirmwareRepository) Create(firmware *entity.Firmware) (*entity.Firmware, error){
	model := repo.mapper.ToModel(firmware)
	err := repo.db.Create(model).Error

	if err != nil{
		return nil, err
	}

	return repo.mapper.ToEntity(model), nil
}

func (repo *GormFirmwareRepository) FindByUuid(id uint64) (*entity.Firmware, error) {
	var m model.FirmwareModel
	
	err := repo.db.First(&m, "id = ?", id).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return repo.mapper.ToEntity(&m), nil
}

func (repo *GormFirmwareRepository) FindAll() ([]*entity.Firmware, error) {
	var models []model.FirmwareModel
	err := repo.db.Find(&models).Error;
	if err != nil {
		return nil, err
	}

	devices := make([]*entity.Firmware, len(models))
	for i, m := range models {
		devices[i] = repo.mapper.ToEntity(&m)
	}

	return devices, nil
}

func (repo *GormFirmwareRepository) Update(firmware entity.Firmware) (*entity.Firmware, error) {
	m := repo.mapper.ToModel(&firmware)

	err := repo.db.Save(m).Error;

	if err != nil {
		return nil, err
	}

	return repo.mapper.ToEntity(m), nil
}

func (repo *GormFirmwareRepository) Delete(id uint64) error {
	err := repo.db.Delete(&model.FirmwareModel{}, "id = ?", id).Error; 

	if err != nil {
		return err
	}
	
	return nil
}