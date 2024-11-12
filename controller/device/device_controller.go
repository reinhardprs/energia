package device

import (
	"energia/controller/base"
	"energia/controller/device/request"
	"energia/controller/device/response"
	"energia/service/device"
	"errors"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func NewDeviceController(ds device.DeviceInterface) *DeviceController {
	return &DeviceController{
		deviceService: ds,
	}
}

type DeviceController struct {
	deviceService device.DeviceInterface
}

func (deviceController DeviceController) CreateDeviceController(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID := int(claims["userID"].(float64))

	deviceCreate := request.CreateDeviceRequest{}
	if err := c.Bind(&deviceCreate); err != nil {
		return base.ErrorResponse(c, err)
	}

	device, err := deviceController.deviceService.Create(userID, deviceCreate.ToEntities())
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, response.FromEntities(device))
}

func (deviceController DeviceController) UpdateDeviceController(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID := int(claims["userID"].(float64))

	deviceIDParam := c.Param("id")
	deviceID, err := strconv.Atoi(deviceIDParam)
	if err != nil {
		return base.ErrorResponse(c, errors.New("Invalid device ID"))
	}

	deviceUpdate := request.UpdateDeviceRequest{}
	if err := c.Bind(&deviceUpdate); err != nil {
		return base.ErrorResponse(c, err)
	}

	device := deviceUpdate.ToEntities()
	device.ID = deviceID
	device.UserID = userID

	updatedDevice, err := deviceController.deviceService.Update(userID, device)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, response.FromEntities(updatedDevice))
}

func (deviceController DeviceController) DeleteDeviceController(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID := int(claims["userID"].(float64))

	deviceIDStr := c.Param("id")
	deviceID, err := strconv.Atoi(deviceIDStr)
	if err != nil {
		return base.ErrorResponse(c, errors.New("invalid device ID"))
	}

	deviceController.deviceService.Delete(userID, deviceID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, nil)
}

func (deviceController DeviceController) GetDeviceController(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID := int(claims["userID"].(float64))

	deviceIDStr := c.Param("id")
	deviceID, err := strconv.Atoi(deviceIDStr)
	if err != nil {
		return base.ErrorResponse(c, errors.New("invalid device ID"))
	}

	device, err := deviceController.deviceService.FindByID(userID, deviceID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, response.FromEntities(device))
}

func (deviceController DeviceController) GetDevicesController(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID := int(claims["userID"].(float64))

	devices, err := deviceController.deviceService.FindAll(userID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, response.FromEntitiesArray(devices))
}
