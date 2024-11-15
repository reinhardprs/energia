package routes

import (
	"energia/controller/suggestion"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type SuggestionRoutes struct {
	SuggestionController *controller.SuggestionController
}

func (sr *SuggestionRoutes) InitSuggestionRoutes(e *echo.Echo) {
	eJWT := e.Group("")
	eJWT.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))

	eSuggestion := eJWT.Group("/suggestion")
	eSuggestion.GET("", sr.SuggestionController.GetSuggestions)
}
