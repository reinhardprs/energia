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

// LoginController handles user login
// @Summary User Login
// @Description Log in a user using email and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body request.LoginRequest true "Login Request Body"
// @Success 200 {object} response.AuthResponse
// @Failure 400 {object} base.BaseResponse "Invalid Request"
// @Router /login [post]
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

// RegisterController handles user registration
// @Summary User Registration
// @Description Register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param register body request.RegisterRequest true "Register Request Body"
// @Success 200 {object} response.AuthResponse
// @Failure 400 {object} base.BaseResponse "Invalid Request"
// @Router /register [post]
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
