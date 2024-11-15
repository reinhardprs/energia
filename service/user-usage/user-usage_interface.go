package user_usage

import (
	"energia/entities"
	"time"
)

type UserUsageInterface interface {
	GetUserUsage(userID int) ([]entities.UserUsage, error)
	Create(userID int, date time.Time) (entities.UserUsage, error)
}
