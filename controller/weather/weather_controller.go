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

// GetWeatherByCityAndDate is the controller for the get weather by city and date endpoint
// @Summary Get weather by city and date
// @Description Get weather by city and date
// @Tags Weather
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param city body request.CreateWeatherRequest true "City of the weather"
// @Success 200 {object} response.WeatherResponse
// @Failure 400 {object} base.BaseResponse
// @Router /weather [get]
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
