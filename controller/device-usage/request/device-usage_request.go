package request

import (
	"energia/entities"
	"time"
)

// CreateDeviceUsageRequest is the request for the create device-usage endpoint
// @Description CreateDeviceUsageRequest is the request for the create device-usage endpoint
// @Param DeviceID int true "ID of the device"
// @Param StartTime string true "Start time of the device usage"
// @Param EndTime string true "End time of the device usage"
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
