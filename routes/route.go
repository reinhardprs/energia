package routes

import (
	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthRoutes        *AuthRoutes
	DeviceRoutes      *DeviceRoutes
	DeviceUsageRoutes *DeviceUsageRoutes
	UserUsageRoutes   *UserUsageRoutes
	WeatherRoutes     *WeatherRoutes
	SuggestionRoutes  *SuggestionRoutes
	EmailRoutes       *EmailRoutes
}

func (rc *RouteController) InitRoute(e *echo.Echo) {
	rc.AuthRoutes.InitAuthRoutes(e)
	rc.DeviceRoutes.InitDeviceRoutes(e)
	rc.DeviceUsageRoutes.InitDeviceUsageRoutes(e)
	rc.UserUsageRoutes.InitUserUsageRoutes(e)
	rc.WeatherRoutes.InitWeatherRoutes(e)
	rc.SuggestionRoutes.InitSuggestionRoutes(e)
	rc.EmailRoutes.InitEmailRoutes(e)
}
