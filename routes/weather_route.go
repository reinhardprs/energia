package routes

import (
	"energia/controller/weather"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type WeatherRoutes struct {
	WeatherController *weather.WeatherController
}

func (wr *WeatherRoutes) InitWeatherRoutes(e *echo.Echo) {
	eJWT := e.Group("")
	eJWT.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))

	eWeather := eJWT.Group("/weather")
	eWeather.GET("", wr.WeatherController.GetWeatherByCityAndDate)
}
