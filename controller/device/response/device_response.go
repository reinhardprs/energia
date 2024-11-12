package response

import "energia/entities"

type DeviceResponse struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Power float32 `json:"power"`
}

func FromEntities(device entities.Device) DeviceResponse {
	return DeviceResponse{
		ID:    device.ID,
		Name:  device.Name,
		Power: device.Power,
	}
}

func FromEntitiesArray(devices []entities.Device) []DeviceResponse {
	var deviceResponses []DeviceResponse
	for _, device := range devices {
		deviceResponses = append(deviceResponses, FromEntities(device))
	}
	return deviceResponses
}
