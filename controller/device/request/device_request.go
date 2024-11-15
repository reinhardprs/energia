package request

import "energia/entities"

type CreateDeviceRequest struct {
	Name string `json:"name"`
	Power float32 `json:"power"`
}

type UpdateDeviceRequest struct {
	Name string `json:"name"`
	Power float32 `json:"power"`
}

func (createDeviceRequest CreateDeviceRequest) ToEntities() entities.Device {
	return entities.Device{
		Name: createDeviceRequest.Name,
		Power: createDeviceRequest.Power,
	}
}

func (updateDeviceRequest UpdateDeviceRequest) ToEntities() entities.Device {
	return entities.Device{
		Name: updateDeviceRequest.Name,
		Power: updateDeviceRequest.Power,
	}
}

