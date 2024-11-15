// package request
package request

import (
	"energia/entities"
	"time"
)

type CreateWeatherRequest struct {
	City string `json:"city"`
}

func (weatherRequest CreateWeatherRequest) ToEntities() entities.Weather {
	return entities.Weather{
		City: weatherRequest.City,
		Date: time.Now(),
	}
}
