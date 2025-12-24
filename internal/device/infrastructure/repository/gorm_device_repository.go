package repository

import (
	"errors"

	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/domain/entity"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/infrastructure/mapper"
	"github.com/perpavbek/ziyatcode-device-gateway/internal/device/infrastructure/model"
	"gorm.io/gorm"
)

type GormDeviceRepository struct {
	db *gorm.DB
	mapper *mapper.DeviceModelMapper
}

func NewGormDeviceRepository(db *gorm.DB, mapper *mapper.DeviceModelMapper) *GormDeviceRepository{
	return &GormDeviceRepository{
		db:db,
		mapper: mapper,
	}
}

func (repo *GormDeviceRepository) Create(device *entity.Device) (*entity.Device, error){
	model := repo.mapper.ToModel(device)
	err := repo.db.Create(model).Error

	if err != nil{
		return nil, err
	}

	return repo.mapper.ToEntity(model), nil
}

func (repo *GormDeviceRepository) FindByUuid(uuid string) (*entity.Device, error) {
	var m model.DeviceModel
	
	err := repo.db.First(&m, "uuid = ?", uuid).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return repo.mapper.ToEntity(&m), nil
}

func (repo *GormDeviceRepository) FindAll() ([]*entity.Device, error) {
	var models []model.DeviceModel
	err := repo.db.Find(&models).Error;
	if err != nil {
		return nil, err
	}

	devices := make([]*entity.Device, len(models))
	for i, m := range models {
		devices[i] = repo.mapper.ToEntity(&m)
	}

	return devices, nil
}

func (repo *GormDeviceRepository) Update(device entity.Device) (*entity.Device, error) {
	m := repo.mapper.ToModel(&device)

	err := repo.db.Save(m).Error;

	if err != nil {
		return nil, err
	}

	return repo.mapper.ToEntity(m), nil
}

func (repo *GormDeviceRepository) Delete(uuid string) error {
	err := repo.db.Delete(&model.DeviceModel{}, "uuid = ?", uuid).Error; 

	if err != nil {
		return err
	}
	
	return nil
}