package auth

import (
	"energia/controller/auth/request"
	"energia/controller/auth/response"
	"energia/controller/base"
	"energia/service/auth"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
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
	if err := c.Bind(&userLogin); err != nil {
		return base.ErrorResponse(c, err)
	}

	user, err := userController.authService.Login(userLogin.ToEntities())
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, response.FromEntities(user))
}

func (userController AuthController) RegisterController(c echo.Context) error {
	userRegister := request.RegisterRequest{}
	if err := c.Bind(&userRegister); err != nil {
		return base.ErrorResponse(c, err)
	}

	user, err := userController.authService.Register(userRegister.ToEntities())
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, response.FromEntities(user))
}

func (userController AuthController) UserController(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID := claims["userID"]

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User authenticated",
		"userID":  userID,
	})
}
