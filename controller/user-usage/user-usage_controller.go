package user_usage

import (
	"energia/constant"
	"energia/controller/base"
	"energia/controller/user-usage/request"
	"energia/controller/user-usage/response"
	user_usage "energia/service/user-usage"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func NewUserUsageController(us user_usage.UserUsageInterface) *UserUsageController {
	return &UserUsageController{
		userUsageService: us,
	}
}

type UserUsageController struct {
	userUsageService user_usage.UserUsageInterface
}

// CreateUserUsageController is the controller for the create user usage endpoint
// @Summary Create a new user usage
// @Description Create a new user usage
// @Tags User Usage
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param user_usage body request.CreateUserUsageRequest true "User usage to create"
// @Success 200 {object} response.UserUsageResponse
// @Failure 400 {object} base.BaseResponse
// @Router /user-usage [post]
func (userUsageController *UserUsageController) CreateUserUsageController(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims, ok := userToken.Claims.(jwt.MapClaims)
	if !ok {
		return base.ErrorResponse(c, constant.INVALID_TOKEN_CLAIMS)
	}

	userID, ok := claims["userID"].(float64)
	if !ok {
		return base.ErrorResponse(c, constant.USER_ID_NOT_FOUND_IN_TOKEN)
	}

	var req request.CreateUserUsageRequest
	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return base.ErrorResponse(c, constant.INVALID_DATE_FORMAT)
	}

	createdUserUsage, err := userUsageController.userUsageService.Create(int(userID), date)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated, response.FromEntities(createdUserUsage))
}

// FindUserUsageController is the controller for the find user usage endpoint
// @Summary Find user usage
// @Description Find user usage
// @Tags User Usage
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.UserUsageResponse
// @Failure 400 {object} base.BaseResponse
// @Router /user-usage [get]
func (userUsageController *UserUsageController) FindUserUsageController(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims, ok := userToken.Claims.(jwt.MapClaims)
	if !ok {
		return base.ErrorResponse(c, constant.INVALID_TOKEN_CLAIMS)
	}

	userID, ok := claims["userID"].(float64)
	if !ok {
		return base.ErrorResponse(c, constant.USER_ID_NOT_FOUND_IN_TOKEN)
	}

	userUsages, err := userUsageController.userUsageService.GetUserUsage(int(userID))
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	userUsageResponses := response.FromEntitiesArray(userUsages)

	return c.JSON(http.StatusOK, userUsageResponses)
}
