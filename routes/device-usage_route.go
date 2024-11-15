package routes

import (
	device_usage "energia/controller/device-usage"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type DeviceUsageRoutes struct {
	DeviceUsageController *device_usage.DeviceUsageController
}

func (dur *DeviceUsageRoutes) InitDeviceUsageRoutes(e *echo.Echo) {
	eJWT := e.Group("")
	eJWT.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))

	eDeviceUsage := eJWT.Group("/device-usage")
	eDeviceUsage.POST("", dur.DeviceUsageController.CreateDeviceUsageController)
	eDeviceUsage.GET("", dur.DeviceUsageController.FindAllDeviceUsageController)
}
