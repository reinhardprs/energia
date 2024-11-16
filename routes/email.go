package routes

import (
	"energia/controller/email"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type EmailRoutes struct {
	EmailController *email.EmailController
}

func (er *EmailRoutes) InitEmailRoutes(e *echo.Echo) {
	eJWT := e.Group("")
	eJWT.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))

	eSuggestion := eJWT.Group("/report")
	eSuggestion.GET("", er.EmailController.SendDeviceUsageReportHandler)
}
