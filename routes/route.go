package routes

import (
	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthRoutes        *AuthRoutes
	DeviceRoutes      *DeviceRoutes
	DeviceUsageRoutes *DeviceUsageRoutes
	UserUsageRoutes   *UserUsageRoutes
}

func (rc *RouteController) InitRoute(e *echo.Echo) {
	rc.AuthRoutes.InitAuthRoutes(e)
	rc.DeviceRoutes.InitDeviceRoutes(e)
	rc.DeviceUsageRoutes.InitDeviceUsageRoutes(e)
	rc.UserUsageRoutes.InitUserUsageRoutes(e)
}
