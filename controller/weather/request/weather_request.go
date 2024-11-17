// package request
package request

import (
	"energia/entities"
	"time"
)

// CreateWeatherRequest is the request for the create weather endpoint
// @Description CreateWeatherRequest is the request for the create weather endpoint
// @Param City string true "City of the weather"
type CreateWeatherRequest struct {
	City string `json:"city"`
}

func (weatherRequest CreateWeatherRequest) ToEntities() entities.Weather {
	return entities.Weather{
		City: weatherRequest.City,
		Date: time.Now(),
	}
}
