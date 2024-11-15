package weather

import (
	"energia/entities"
	"time"
)

type WeatherRepoInterface interface {
	GetWeather(city string) ([]entities.Weather, error)
	FetchAndStoreWeather(city string) (entities.Weather, error)
	GetWeatherByCityAndDate(city string, date time.Time) (entities.Weather, error)
}
