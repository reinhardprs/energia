package config

import (
	"energia/repository/auth"
	"energia/repository/device"
	device_usage "energia/repository/device-usage"
	user_usage "energia/repository/user-usage"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(
		&auth.User{},
		&device.Device{},
		&device_usage.DeviceUsage{},
		&user_usage.UserUsage{},
	)
}
