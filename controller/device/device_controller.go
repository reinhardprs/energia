package device

import (
	"energia/constant"
	"energia/controller/base"
	"energia/controller/device/request"
	"energia/controller/device/response"
	"energia/service/device"
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

// CreateDeviceController is the controller for the create device endpoint
// @Summary Create a new device
// @Description Create a new device
// @Tags Device
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param device body request.CreateDeviceRequest true "Device to create"
// @Success 200 {object} response.DeviceResponse
// @Failure 400 {object} base.BaseResponse
// @Router /device [post]
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

// UpdateDeviceController is the controller for the update device endpoint
// @Summary Update a device
// @Description Update a device
// @Tags Device
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "ID of the device to update"
// @Param device body request.UpdateDeviceRequest true "Device to update"
// @Success 200 {object} response.DeviceResponse
// @Failure 400 {object} base.BaseResponse
// @Router /device/{id} [put]
func (deviceController DeviceController) UpdateDeviceController(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID := int(claims["userID"].(float64))

	deviceIDParam := c.Param("id")
	deviceID, err := strconv.Atoi(deviceIDParam)
	if err != nil {
		return base.ErrorResponse(c, constant.INVALID_DEVICE_ID)
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

// DeleteDeviceController is the controller for the delete device endpoint
// @Summary Delete a device
// @Description Delete a device
// @Tags Device
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "ID of the device to delete"
// @Success 200 {object} base.BaseResponse
// @Failure 400 {object} base.BaseResponse
// @Router /device/{id} [delete]
func (deviceController DeviceController) DeleteDeviceController(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID := int(claims["userID"].(float64))

	deviceIDStr := c.Param("id")
	deviceID, err := strconv.Atoi(deviceIDStr)
	if err != nil {
		return base.ErrorResponse(c, constant.INVALID_DEVICE_ID)
	}

	err = deviceController.deviceService.Delete(userID, deviceID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, nil)
}

// GetDeviceController is the controller for the get device endpoint
// @Summary Get a device
// @Description Get a device
// @Tags Device
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "ID of the device to get"
// @Success 200 {object} response.DeviceResponse
// @Failure 400 {object} base.BaseResponse
// @Router /device/{id} [get]
func (deviceController DeviceController) GetDeviceController(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID := int(claims["userID"].(float64))

	deviceIDStr := c.Param("id")
	deviceID, err := strconv.Atoi(deviceIDStr)
	if err != nil {
		return base.ErrorResponse(c, constant.INVALID_DEVICE_ID)
	}

	device, err := deviceController.deviceService.FindByID(userID, deviceID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, response.FromEntities(device))
}

// GetDevicesController is the controller for the get devices endpoint
// @Summary Get all devices
// @Description Get all devices
// @Tags Device
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.DeviceResponse
// @Failure 400 {object} base.BaseResponse
// @Router /device [get]
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
