package routes

import (
	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthRoutes *AuthRoutes
}

func (rc *RouteController) InitRoute(e *echo.Echo) {
	rc.AuthRoutes.InitAuthRoutes(e)
}