package request

import "energia/entities"

// CreateDeviceRequest is the request for the create device endpoint
// @Description CreateDeviceRequest is the request for the create device endpoint
// @Param Name string true "Name of the device"
// @Param Power float32 true "Power of the device"
type CreateDeviceRequest struct {
	Name  string  `json:"name"`
	Power float32 `json:"power"`
}

// UpdateDeviceRequest is the request for the update device endpoint
// @Description UpdateDeviceRequest is the request for the update device endpoint
// @Param Name string true "Name of the device"
// @Param Power float32 true "Power of the device"
type UpdateDeviceRequest struct {
	Name  string  `json:"name"`
	Power float32 `json:"power"`
}

func (createDeviceRequest CreateDeviceRequest) ToEntities() entities.Device {
	return entities.Device{
		Name:  createDeviceRequest.Name,
		Power: createDeviceRequest.Power,
	}
}

func (updateDeviceRequest UpdateDeviceRequest) ToEntities() entities.Device {
	return entities.Device{
		Name:  updateDeviceRequest.Name,
		Power: updateDeviceRequest.Power,
	}
}
