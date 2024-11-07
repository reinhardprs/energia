package config

import (
	"energia/repository/auth"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&auth.User{})
}
