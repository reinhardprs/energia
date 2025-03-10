package weather

import (
	"energia/entities"
	"time"
)

type Weather struct {
	ID          int       `gorm:"primaryKey"`
	City        string    `gorm:"not null;type:varchar(255)"`
	Date        time.Time `gorm:"not null"`
	Temperature float32   `gorm:"not null;type:float"`
	Humidity    float32   `gorm:"not null;type:float"`
	Description string    `gorm:"not null;type:text"`
}

func FromEntities(weather entities.Weather) Weather {
	return Weather{
		ID:          weather.ID,
		City:        weather.City,
		Date:        weather.Date,
		Temperature: weather.Temperature,
		Humidity:    weather.Humidity,
		Description: weather.Description,
	}
}

func (weather Weather) ToEntities() entities.Weather {
	return entities.Weather{
		ID:          weather.ID,
		City:        weather.City,
		Date:        weather.Date,
		Temperature: weather.Temperature,
		Humidity:    weather.Humidity,
		Description: weather.Description,
	}
}
