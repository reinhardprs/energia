package weather

import "energia/entities"

type WeatherRepoInterface interface {
	Create(weather entities.Weather) (entities.Weather, error)
	FindAll(city string) ([]entities.Weather, error)
}
