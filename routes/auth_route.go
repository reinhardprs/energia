package routes

import (
	"os"
	"energia/controller/auth"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type AuthRoutes struct {
	AuthController *auth.AuthController
}

func (ar *AuthRoutes) InitAuthRoutes(e *echo.Echo) {
	e.POST("/login", ar.AuthController.LoginController)
	e.POST("/register", ar.AuthController.RegisterController)

	eJWT := e.Group("")
	eJWT.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))

	eUser := eJWT.Group("/users")
	eUser.GET("", ar.AuthController.LoginController)
}