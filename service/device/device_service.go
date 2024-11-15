package device

import (
	"energia/constant"
	"energia/entities"
	"energia/repository/device"
)

func NewDeviceService(dr device.DeviceRepoInterface) *DeviceService {
	return &DeviceService{
		deviceRepoInterface: dr,
	}
}

type DeviceService struct {
	deviceRepoInterface device.DeviceRepoInterface
}

func (deviceService DeviceService) Create(userID int, device entities.Device) (entities.Device, error) {
	if device.Name == "" {
		return entities.Device{}, constant.NAME_IS_EMPTY
	} else if device.Power == 0 {
		return entities.Device{}, constant.POWER_IS_EMPTY
	}

	device.UserID = userID

	device, err := deviceService.deviceRepoInterface.Create(userID, device)
	if err != nil {
		return entities.Device{}, err
	}

	return device, nil
}

func (deviceService DeviceService) FindAll(userID int) ([]entities.Device, error) {
	devices, err := deviceService.deviceRepoInterface.FindAll(userID)
	if err != nil {
		return nil, err
	}

	return devices, nil
}

func (deviceService DeviceService) FindByID(userID int, deviceID int) (entities.Device, error) {
	device, err := deviceService.deviceRepoInterface.FindByID(userID, deviceID)
	if err != nil {
		return entities.Device{}, err
	}

	return device, nil
}

func (deviceService DeviceService) Update(userID int, device entities.Device) (entities.Device, error) {
	if device.Name == "" {
		return entities.Device{}, constant.NAME_IS_EMPTY
	} else if device.Power == 0 {
		return entities.Device{}, constant.POWER_IS_EMPTY
	}

	device.UserID = userID

	device, err := deviceService.deviceRepoInterface.Update(userID, device)
	if err != nil {
		return entities.Device{}, err
	}

	return device, nil
}

func (deviceService DeviceService) Delete(userID int, deviceID int) error {
	err := deviceService.deviceRepoInterface.Delete(userID, deviceID)
	if err != nil {
		return err
	}

	return nil
}
