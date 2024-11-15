package entities

import "time"

type DeviceUsage struct {
	ID             int
	UserID         int
	DeviceID       int
	StartTime      time.Time
	EndTime        time.Time
	Duration       float32
	EnergyConsumed float32
}
