package weather

import (
	"energia/controller/base"
	"energia/controller/weather/request"
	"energia/controller/weather/response"
	"energia/service/weather"

	"github.com/labstack/echo/v4"
)

type WeatherController struct {
	weatherService *weather.WeatherService
}

func NewWeatherController(ws *weather.WeatherService) *WeatherController {
	return &WeatherController{
		weatherService: ws,
	}
}

func (wc *WeatherController) GetWeatherByCityAndDate(c echo.Context) error {
	var req request.CreateWeatherRequest
	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	weatherEntity := req.ToEntities()

	weatherData, err := wc.weatherService.GetWeatherByCityAndDate(weatherEntity.City, weatherEntity.Date)
	if err != nil {
		weatherData, err = wc.weatherService.FetchAndStoreWeather(weatherEntity.City)
		if err != nil {
			return base.ErrorResponse(c, err)
		}
	}

	return base.SuccessResponse(c, response.FromEntities(weatherData))
}
