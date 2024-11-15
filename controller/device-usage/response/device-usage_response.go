package response

import "energia/entities"

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
