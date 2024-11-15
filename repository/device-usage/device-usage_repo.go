package device_usage

import (
	"energia/entities"
	"time"

	"gorm.io/gorm"
)

func NewDeviceUsageRepo(db *gorm.DB) *DeviceUsageRepo {
	return &DeviceUsageRepo{
		db: db,
	}
}

type DeviceUsageRepo struct {
	db *gorm.DB
}

func (deviceUsageRepo *DeviceUsageRepo) Create(deviceUsage entities.DeviceUsage) (entities.DeviceUsage, error) {
	deviceUsageDb := FromEntities(deviceUsage)

	result := deviceUsageRepo.db.Create(&deviceUsageDb)

	if result.Error != nil {
		return entities.DeviceUsage{}, result.Error
	}

	return deviceUsageDb.ToEntities(), nil
}

func (deviceUsageRepo *DeviceUsageRepo) FindAll(userID int) ([]entities.DeviceUsage, error) {
	var deviceUsagesDb []DeviceUsage

	result := deviceUsageRepo.db.Preload("Device", "user_id = ?", userID).
		Find(&deviceUsagesDb)

	if result.Error != nil {
		return nil, result.Error
	}

	deviceUsages := make([]entities.DeviceUsage, len(deviceUsagesDb))
	for i, deviceUsageDb := range deviceUsagesDb {
		deviceUsages[i] = deviceUsageDb.ToEntities()
	}

	return deviceUsages, nil
}

func (deviceUsageRepo *DeviceUsageRepo) GetDeviceUsageByDate(userID int, date time.Time) ([]entities.DeviceUsage, error) {
	var deviceUsagesDb []DeviceUsage

	result := deviceUsageRepo.db.Preload("Device", "user_id = ?", userID).
		Where("DATE(start_time) = ?", date.Format("2006-01-02")).Find(&deviceUsagesDb)

	if result.Error != nil {
		return nil, result.Error
	}

	deviceUsages := make([]entities.DeviceUsage, len(deviceUsagesDb))
	for i, deviceUsageDb := range deviceUsagesDb {
		deviceUsages[i] = deviceUsageDb.ToEntities()
	}

	return deviceUsages, nil
}
