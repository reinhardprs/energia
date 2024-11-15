package device

import "energia/entities"

type DeviceInterface interface {
	Create(userID int, device entities.Device) (entities.Device, error)
	FindAll(userID int) ([]entities.Device, error)
	FindByID(userID int, deviceID int) (entities.Device, error)
	Update(userID int, device entities.Device) (entities.Device, error)
	Delete(userID int, deviceID int) error
}