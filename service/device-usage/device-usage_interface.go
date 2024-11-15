package device_usage

import "energia/entities"

type DeviceUsageInterface interface {
	Create(deviceUsage entities.DeviceUsage, userID int) (entities.DeviceUsage, error)
	FindAll(userID int) ([]entities.DeviceUsage, error)
}
