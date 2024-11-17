package response

import "energia/entities"

// DeviceUsageResponse is the response for the device-usage endpoint
// @Description DeviceUsageResponse is the response for the device-usage endpoint
// @Param ID int true "ID of the device usage"
// @Param DeviceID int true "ID of the device"
// @Param StartTime string true "Start time of the device usage"
// @Param EndTime string true "End time of the device usage"
// @Param Duration float32 true "Duration of the device usage"
// @Param EnergyConsumed float32 true "Energy consumed of the device usage"
type DeviceUsageResponse struct {
	ID             int     `json:"id"`
	DeviceID       int     `json:"device_id"`
	StartTime      string  `json:"start_time"`
	EndTime        string  `json:"end_time"`
	Duration       float32 `json:"duration"`
	EnergyConsumed float32 `json:"energy_consumed"`
}

func FromEntities(deviceUsage entities.DeviceUsage) DeviceUsageResponse {
	return DeviceUsageResponse{
		ID:             deviceUsage.ID,
		DeviceID:       deviceUsage.DeviceID,
		StartTime:      deviceUsage.StartTime.String(),
		EndTime:        deviceUsage.EndTime.String(),
		Duration:       deviceUsage.Duration,
		EnergyConsumed: deviceUsage.EnergyConsumed,
	}
}

func FromEntitiesArray(deviceUsages []entities.DeviceUsage) []DeviceUsageResponse {
	var deviceUsageResponses []DeviceUsageResponse
	for _, deviceUsage := range deviceUsages {
		deviceUsageResponses = append(deviceUsageResponses, FromEntities(deviceUsage))
	}
	return deviceUsageResponses
}
