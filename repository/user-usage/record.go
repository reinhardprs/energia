package user_usage

import (
	"energia/entities"
	"energia/repository/auth"
	"time"
)

type UserUsage struct {
	ID          int       `gorm:"primaryKey"`
	UserID      int       `gorm:"not null"`
	User        auth.User `gorm:"foreignKey:UserID;references:ID"`
	Date        time.Time `gorm:"not null"`
	TotalEnergy float32   `gorm:"not null"`
	TotalCost   float32   `gorm:"not null"`
}

func FromEntities(userUsage entities.UserUsage) UserUsage {
	return UserUsage{
		ID:          userUsage.ID,
		UserID:      userUsage.UserID,
		Date:        userUsage.Date,
		TotalEnergy: userUsage.TotalEnergy,
		TotalCost:   userUsage.TotalCost,
	}
}

func (userUsage UserUsage) ToEntities() entities.UserUsage {
	return entities.UserUsage{
		ID:          userUsage.ID,
		UserID:      userUsage.UserID,
		Date:        userUsage.Date,
		TotalEnergy: userUsage.TotalEnergy,
		TotalCost:   userUsage.TotalCost,
	}
}
