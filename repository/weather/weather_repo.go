package weather

import (
	"energia/entities"

	"gorm.io/gorm"
)

type WeatherRepo struct {
	db *gorm.DB
}

func NewWeatherRepo(db *gorm.DB) *WeatherRepo {
	return &WeatherRepo{db: db}
}

func (wr *WeatherRepo) Create(weather entities.Weather) (entities.Weather, error) {
	weatherRecord := FromEntities(weather)

	result := wr.db.Create(&weatherRecord)
	if result.Error != nil {
		return entities.Weather{}, result.Error
	}

	return weatherRecord.ToEntities(), nil
}

func (wr *WeatherRepo) FindAll(city string) ([]entities.Weather, error) {
	var weather []Weather

	result := wr.db.Where("city = ?", city).Find(&weather)
	if result.Error != nil {
		return nil, result.Error
	}

	weathers := make([]entities.Weather, len(weather))
	for i, weatherRecord := range weather {
		weathers[i] = weatherRecord.ToEntities()
	}

	return weathers, nil
}
