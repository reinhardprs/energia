package request

import (
	"energia/entities"
	"time"
)

type CreateUserUsageRequest struct {
	Date string `json:"date"`
}

func (createUserUsageRequest CreateUserUsageRequest) ToEntities() (entities.UserUsage, error) {
	// Parsing Date dari string ke time.Time dengan format "2006-01-02"
	parsedDate, err := time.Parse("2006-01-02", createUserUsageRequest.Date)
	if err != nil {
		return entities.UserUsage{}, err
	}

	return entities.UserUsage{
		Date: parsedDate,
	}, nil
}

