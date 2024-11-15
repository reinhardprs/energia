package auth

import (
	"energia/controller/auth/request"
	"energia/controller/auth/response"
	"energia/controller/base"
	"energia/service/auth"

	"github.com/labstack/echo/v4"
)

func NewAuthController(as auth.AuthInterface) *AuthController {
	return &AuthController{
		authService: as,
	}
}

type AuthController struct {
	authService auth.AuthInterface
}

func (userController AuthController) LoginController(c echo.Context) error {
	userLogin := request.LoginRequest{}
	c.Bind(&userLogin)
	user, err := userController.authService.Login(userLogin.ToEntities())
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, response.FromEntities(user))
}

func (userController AuthController) RegisterController(c echo.Context) error {
	userRegister := request.RegisterRequest{}
	c.Bind(&userRegister)
	user, err := userController.authService.Register(userRegister.ToEntities())
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, response.FromEntities(user))
}
