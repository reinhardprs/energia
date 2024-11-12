package device

import (
	"energia/entities"

	"gorm.io/gorm"
)

func NewDeviceRepo(db *gorm.DB) *DeviceRepo {
	return &DeviceRepo{
		db: db,
	}
}

type DeviceRepo struct {
	db *gorm.DB
}

func (deviceRepo *DeviceRepo) Create(userID int, device entities.Device) (entities.Device, error) {
	deviceDb := FromEntities(device)
	deviceDb.UserID = userID

	result := deviceRepo.db.Create(&deviceDb)

	if result.Error != nil {
		return entities.Device{}, result.Error
	}

	return deviceDb.ToEntities(), nil
}

func (deviceRepo *DeviceRepo) FindAll(userID int) ([]entities.Device, error) {
	var devicesDb []Device
	result := deviceRepo.db.Find(&devicesDb, "user_id = ?", userID)

	if result.Error != nil {
		return nil, result.Error
	}

	devices := make([]entities.Device, len(devicesDb))
	for i, deviceDb := range devicesDb {
		devices[i] = deviceDb.ToEntities()
	}

	return devices, nil
}

func (deviceRepo *DeviceRepo) FindByID(userID int, deviceID int) (entities.Device, error) {
	var deviceDb Device
	result := deviceRepo.db.First(&deviceDb, "user_id = ? AND id = ?", userID, deviceID)

	if result.Error != nil {
		return entities.Device{}, result.Error
	}

	return deviceDb.ToEntities(), nil
}

func (deviceRepo *DeviceRepo) Update(userID int, device entities.Device) (entities.Device, error) {
	deviceDb := Device{}
	result := deviceRepo.db.Where("id = ? AND user_id = ?", device.ID, userID).First(&deviceDb)
	if result.Error != nil {
			return entities.Device{}, result.Error
	}

	deviceDb.Name = device.Name
	deviceDb.Power = device.Power

	if err := deviceRepo.db.Save(&deviceDb).Error; err != nil {
			return entities.Device{}, err
	}

	return deviceDb.ToEntities(), nil
}



func (deviceRepo *DeviceRepo) Delete(deviceID int, userID int) error {
	result := deviceRepo.db.Delete(&Device{}, "user_id = ? AND id = ?", deviceID, userID)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
