package routes

import (
    "os"
    "energia/controller/device"

    echojwt "github.com/labstack/echo-jwt/v4"
    "github.com/labstack/echo/v4"
)

type DeviceRoutes struct {
    DeviceController *device.DeviceController
}

func (dr *DeviceRoutes) InitDeviceRoutes(e *echo.Echo) {
    eJWT := e.Group("")
    eJWT.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))

    eDevice := eJWT.Group("/devices")
    eDevice.POST("", dr.DeviceController.CreateDeviceController)
    eDevice.GET("/:id", dr.DeviceController.GetDeviceController)
    eDevice.PUT("/:id", dr.DeviceController.UpdateDeviceController)
    eDevice.DELETE("/:id", dr.DeviceController.DeleteDeviceController)
    eDevice.GET("", dr.DeviceController.GetDevicesController)
}
