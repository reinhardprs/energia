package routes

import (
	user_usage "energia/controller/user-usage"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type UserUsageRoutes struct {
	UserUsageController *user_usage.UserUsageController
}

func (ur *UserUsageRoutes) InitUserUsageRoutes(e *echo.Echo) {
	eJWT := e.Group("")
	eJWT.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))

	eUserUsage := eJWT.Group("/user-usage")
	eUserUsage.POST("", ur.UserUsageController.CreateUserUsageController)
	eUserUsage.GET("", ur.UserUsageController.FindUserUsageController) // Hapus :user_id
}
