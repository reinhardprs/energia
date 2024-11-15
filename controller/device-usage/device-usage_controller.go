package device_usage

import (
	"energia/constant"
	"energia/controller/base"
	"energia/controller/device-usage/request"
	"energia/controller/device-usage/response"
	du_service "energia/service/device-usage"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func NewDeviceUsageController(dus du_service.DeviceUsageInterface) *DeviceUsageController {
	return &DeviceUsageController{
		deviceUsageService: dus,
	}
}

type DeviceUsageController struct {
	deviceUsageService du_service.DeviceUsageInterface
}

func (deviceUsageController DeviceUsageController) CreateDeviceUsageController(c echo.Context) error {
	var req request.CreateDeviceUsageRequest
	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	deviceUsage := req.ToEntities()

	userToken := c.Get("user").(*jwt.Token)
	claims, ok := userToken.Claims.(jwt.MapClaims)
	if !ok {
		return base.ErrorResponse(c, constant.INVALID_TOKEN_CLAIMS)
	}

	userID, ok := claims["userID"].(float64)
	if !ok {
		return base.ErrorResponse(c, constant.USER_ID_NOT_FOUND_IN_TOKEN)
	}

	createdUsage, err := deviceUsageController.deviceUsageService.Create(deviceUsage, int(userID))
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, response.FromEntities(createdUsage))
}

func (deviceUsageController DeviceUsageController) FindAllDeviceUsageController(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims, ok := userToken.Claims.(jwt.MapClaims)
	if !ok {
		return base.ErrorResponse(c, constant.INVALID_TOKEN_CLAIMS)
	}

	userID, ok := claims["userID"].(float64)
	if !ok {
		return base.ErrorResponse(c, constant.USER_ID_NOT_FOUND_IN_TOKEN)
	}

	deviceUsages, err := deviceUsageController.deviceUsageService.FindAll(int(userID))
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, response.FromEntitiesArray(deviceUsages))
}
