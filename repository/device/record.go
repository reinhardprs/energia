package device

import (
	"energia/entities"
	"energia/repository/auth"
	"time"
)

type Device struct {
	ID        int       `gorm:"primaryKey"`
	UserID    int       `gorm:"not null"`                       
	User      auth.User `gorm:"foreignKey:UserID;references:ID"`
	Name      string    `gorm:"not null"`
	Power     float32   `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func FromEntities(device entities.Device) Device {
	return Device{
		ID:        device.ID,
		UserID:    device.UserID,
		Name:      device.Name,
		Power:     device.Power,
		CreatedAt: device.CreatedAt,
		UpdatedAt: device.UpdatedAt,
	}
}

func (device Device) ToEntities() entities.Device {
	return entities.Device{
		ID:        device.ID,
		UserID:    device.UserID,
		Name:      device.Name,
		Power:     device.Power,
		CreatedAt: device.CreatedAt,
		UpdatedAt: device.UpdatedAt,
	}
}
