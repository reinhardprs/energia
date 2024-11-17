package response

import "energia/entities"

// DeviceResponse is the response for the device endpoint
// @Description DeviceResponse is the response for the device endpoint
// @Param ID int true "ID of the device"
// @Param UserID int true "ID of the user"
// @Param Name string true "Name of the device"
// @Param Power float32 true "Power of the device"
type DeviceResponse struct {
	ID     int     `json:"id"`
	UserID int     `json:"user_id"`
	Name   string  `json:"name"`
	Power  float32 `json:"power"`
}

func FromEntities(device entities.Device) DeviceResponse {
	return DeviceResponse{
		ID:     device.ID,
		UserID: device.UserID,
		Name:   device.Name,
		Power:  device.Power,
	}
}

func FromEntitiesArray(devices []entities.Device) []DeviceResponse {
	var deviceResponses []DeviceResponse
	for _, device := range devices {
		deviceResponses = append(deviceResponses, FromEntities(device))
	}
	return deviceResponses
}
