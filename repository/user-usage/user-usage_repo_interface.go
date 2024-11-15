package user_usage

import "energia/entities"

type UserUsageRepoInterface interface {
	Create(userUsage entities.UserUsage) (entities.UserUsage, error)
	FindAll(userID int) ([]entities.UserUsage, error)
}
