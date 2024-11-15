package device_usage

import (
	"energia/constant"
	"energia/entities"
	"energia/repository/device"
	device_usage "energia/repository/device-usage"
)

func NewDeviceUsageService(dr device.DeviceRepoInterface, dur device_usage.DeviceUsageRepoInterface) *DeviceUsageService {
	return &DeviceUsageService{
		deviceRepoInterface:      dr,
		deviceUsageRepoInterface: dur,
	}
}

type DeviceUsageService struct {
	deviceRepoInterface      device.DeviceRepoInterface
	deviceUsageRepoInterface device_usage.DeviceUsageRepoInterface
}

func (deviceUsageService DeviceUsageService) Create(deviceUsage entities.DeviceUsage, userID int) (entities.DeviceUsage, error) {
	device, err := deviceUsageService.deviceRepoInterface.FindByID(userID, deviceUsage.DeviceID)
	if err != nil {
		return entities.DeviceUsage{}, err
	}
	if device.ID == 0 {
		return entities.DeviceUsage{}, constant.DEVICE_NOT_FOUND
	}

	deviceUsage.DeviceID = device.ID
	duration := deviceUsage.EndTime.Sub(deviceUsage.StartTime).Minutes()
	deviceUsage.Duration = float32(duration)

	energyConsumed := device.Power * deviceUsage.Duration / 1000
	deviceUsage.EnergyConsumed = energyConsumed
	deviceUsage, err = deviceUsageService.deviceUsageRepoInterface.Create(deviceUsage)
	if err != nil {
		return entities.DeviceUsage{}, err
	}

	return deviceUsage, nil
}

func (deviceUsageService DeviceUsageService) FindAll(userID int) ([]entities.DeviceUsage, error) {
	deviceUsages, err := deviceUsageService.deviceUsageRepoInterface.FindAll(userID)
	if err != nil {
		return nil, err
	}

	return deviceUsages, nil
}
