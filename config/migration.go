package config

import (
	"energia/repository/auth"
	"energia/repository/device"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&auth.User{}, &device.Device{})
}
