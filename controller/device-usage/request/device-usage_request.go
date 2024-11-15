package request

import (
	"energia/entities"
	"time"
)

type CreateDeviceUsageRequest struct {
	DeviceID  int       `json:"device_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

func (createDeviceUsageRequest CreateDeviceUsageRequest) ToEntities() entities.DeviceUsage {
	return entities.DeviceUsage{
		DeviceID:  createDeviceUsageRequest.DeviceID,
		StartTime: createDeviceUsageRequest.StartTime,
		EndTime:   createDeviceUsageRequest.EndTime,
	}
}
