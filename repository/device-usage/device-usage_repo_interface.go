package device_usage

import (
	"energia/entities"
	"time"
)

type DeviceUsageRepoInterface interface {
	Create(deviceUsage entities.DeviceUsage) (entities.DeviceUsage, error)
	FindAll(userID int) ([]entities.DeviceUsage, error)
	GetDeviceUsageByDate(userID int, date time.Time) ([]entities.DeviceUsage, error)
}
