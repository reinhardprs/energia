// package response
package response

import "energia/entities"

// WeatherResponse is the response for the weather endpoint
// @Description WeatherResponse is the response for the weather endpoint
// @Param City string true "City of the weather"
// @Param Temperature float32 true "Temperature of the weather"
// @Param Humidity float32 true "Humidity of the weather"
// @Param Condition string true "Condition of the weather"
// @Param Date string true "Date of the weather"
type WeatherResponse struct {
	City        string  `json:"city"`
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
	Condition   string  `json:"condition"`
	Date        string  `json:"date"`
}

func FromEntities(weather entities.Weather) WeatherResponse {
	return WeatherResponse{
		City:        weather.City,
		Temperature: weather.Temperature,
		Humidity:    weather.Humidity,
		Condition:   weather.Description,
		Date:        weather.Date.Format("2006-01-02"),
	}
}

func FromWeatherEntitiesArray(weathers []entities.Weather) []WeatherResponse {
	var weatherResponses []WeatherResponse
	for _, weather := range weathers {
		weatherResponses = append(weatherResponses, FromEntities(weather))
	}
	return weatherResponses
}
