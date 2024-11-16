package device_usage

import (
	"energia/constant"
	"energia/controller/base"
	"energia/controller/device-usage/request"
	"energia/controller/device-usage/response"
	ds "energia/service/device"
	du_service "energia/service/device-usage"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type DeviceUsageController struct {
	deviceUsageService du_service.DeviceUsageInterface
	deviceService      ds.DeviceInterface // Tambahkan ini
}

func NewDeviceUsageController(dus du_service.DeviceUsageInterface, ds ds.DeviceInterface) *DeviceUsageController {
	return &DeviceUsageController{
		deviceUsageService: dus,
		deviceService:      ds,
	}
}

func (deviceUsageController DeviceUsageController) CreateDeviceUsageController(c echo.Context) error {
	var req request.CreateDeviceUsageRequest
	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	userToken := c.Get("user").(*jwt.Token)
	claims, ok := userToken.Claims.(jwt.MapClaims)
	if !ok {
		return base.ErrorResponse(c, constant.INVALID_TOKEN_CLAIMS)
	}

	userID, ok := claims["userID"].(float64)
	if !ok {
		return base.ErrorResponse(c, constant.USER_ID_NOT_FOUND_IN_TOKEN)
	}

	device, err := deviceUsageController.deviceService.FindByID(int(userID), req.DeviceID)
	if err != nil {
		return base.ErrorResponse(c, constant.DEVICE_NOT_FOUND)
	}

	deviceUsage := req.ToEntities()
	deviceUsage.UserID = int(userID)

	createdUsage, err := deviceUsageController.deviceUsageService.Create(deviceUsage, int(userID))
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, map[string]interface{}{
		"device_usage": response.FromEntities(createdUsage),
		"device_name":  device.Name,
	})
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

	devices, err := deviceUsageController.deviceService.FindAll(int(userID))
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	deviceMap := make(map[int]string)
	for _, device := range devices {
		deviceMap[device.ID] = device.Name
	}

	deviceUsages, err := deviceUsageController.deviceUsageService.FindAll(int(userID))
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	var usageResponses []map[string]interface{}
	for _, usage := range deviceUsages {
		deviceName, exists := deviceMap[usage.DeviceID]
		if !exists {
			continue
		}

		usageResponses = append(usageResponses, map[string]interface{}{
			"device_usage": response.FromEntities(usage),
			"device_name":  deviceName,
		})
	}

	if len(usageResponses) == 0 {
		return base.ErrorResponse(c, constant.DEVICE_USAGE_NOT_FOUND)
	}

	return base.SuccessResponse(c, usageResponses)
}
