package request

import (
	"energia/entities"
	"time"
)

// CreateUserUsageRequest is the request for the create user-usage endpoint
// @Description CreateUserUsageRequest is the request for the create user-usage endpoint
// @Param Date string true "Date of the user usage"
type CreateUserUsageRequest struct {
	Date string `json:"date"`
}

func (createUserUsageRequest CreateUserUsageRequest) ToEntities() (entities.UserUsage, error) {
	parsedDate, err := time.Parse("2006-01-02", createUserUsageRequest.Date)
	if err != nil {
		return entities.UserUsage{}, err
	}

	return entities.UserUsage{
		Date: parsedDate,
	}, nil
}
