package device_usage

import (
	"energia/entities"
	"energia/repository/device"
	"time"
)

type DeviceUsage struct {
	ID             int           `gorm:"primaryKey"`
	DeviceID       int           `gorm:"not null"`
	Device         device.Device `gorm:"foreignKey:DeviceID;references:ID"`
	StartTime      time.Time     `gorm:"not null"`
	EndTime        time.Time     `gorm:"not null"`
	Duration       float32       `gorm:"not null"`
	EnergyConsumed float32       `gorm:"not null"`
}

func FromEntities(deviceUsage entities.DeviceUsage) DeviceUsage {
	return DeviceUsage{
		ID:             deviceUsage.ID,
		DeviceID:       deviceUsage.DeviceID,
		StartTime:      deviceUsage.StartTime,
		EndTime:        deviceUsage.EndTime,
		Duration:       deviceUsage.Duration,
		EnergyConsumed: deviceUsage.EnergyConsumed,
	}
}

func (deviceUsage DeviceUsage) ToEntities() entities.DeviceUsage {
	return entities.DeviceUsage{
		ID:             deviceUsage.ID,
		DeviceID:       deviceUsage.DeviceID,
		StartTime:      deviceUsage.StartTime,
		EndTime:        deviceUsage.EndTime,
		Duration:       deviceUsage.Duration,
		EnergyConsumed: deviceUsage.EnergyConsumed,
	}
}
