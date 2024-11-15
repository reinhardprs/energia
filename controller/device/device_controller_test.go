package device_test

import (
	"bytes"
	"encoding/json"
	"energia/controller/device"
	"energia/controller/device/request"
	"energia/controller/device/response"
	"energia/entities"
	"energia/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func generateJWTToken(userID int) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": float64(userID),
	})
}

func setupContext(e *echo.Echo, method, path string, body []byte) (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, path, bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c, rec
}

func TestCreateDeviceController(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.DeviceInterface)
	controller := device.NewDeviceController(mockService)

	mockDevice := entities.Device{ID: 1, UserID: 1, Name: "Test Device", Power: 100.0}
	mockService.On("Create", 1, mock.AnythingOfType("entities.Device")).Return(mockDevice, nil)

	body := request.CreateDeviceRequest{Name: "Test Device", Power: 100.0}
	bodyBytes, _ := json.Marshal(body)

	c, rec := setupContext(e, http.MethodPost, "/devices", bodyBytes)
	c.Set("user", generateJWTToken(1))

	err := controller.CreateDeviceController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var resp struct {
		Data response.DeviceResponse `json:"data"`
	}
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, 1, resp.Data.ID)
	assert.Equal(t, "Test Device", resp.Data.Name)
}

func TestUpdateDeviceController(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.DeviceInterface)
	controller := device.NewDeviceController(mockService)

	mockDevice := entities.Device{
		ID:     1,
		UserID: 1,
		Name:   "Updated Device",
		Power:  150.0,
	}
	mockService.On("Update", 1, mock.AnythingOfType("entities.Device")).Return(mockDevice, nil)

	body := request.UpdateDeviceRequest{
		Name:  "Updated Device",
		Power: 150.0,
	}
	bodyBytes, _ := json.Marshal(body)

	c, rec := setupContext(e, http.MethodPut, "/devices/1", bodyBytes)
	c.Set("user", generateJWTToken(1))
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := controller.UpdateDeviceController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var resp struct {
		Data response.DeviceResponse `json:"data"`
	}
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Device", resp.Data.Name)
}

func TestDeleteDeviceController(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.DeviceInterface)
	controller := device.NewDeviceController(mockService)

	mockService.On("Delete", 1, 1).Return(nil)

	c, rec := setupContext(e, http.MethodDelete, "/devices/1", nil)
	c.Set("user", generateJWTToken(1))
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := controller.DeleteDeviceController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetDeviceController(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.DeviceInterface)
	controller := device.NewDeviceController(mockService)

	mockDevice := entities.Device{
		ID:     1,
		UserID: 1,
		Name:   "Test Device",
		Power:  100.0,
	}
	mockService.On("FindByID", 1, 1).Return(mockDevice, nil)

	c, rec := setupContext(e, http.MethodGet, "/devices/1", nil)
	c.Set("user", generateJWTToken(1))
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := controller.GetDeviceController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var resp struct {
		Data response.DeviceResponse `json:"data"`
	}
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "Test Device", resp.Data.Name)
}

func TestGetDevicesController(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.DeviceInterface)
	controller := device.NewDeviceController(mockService)

	mockDevices := []entities.Device{
		{ID: 1, UserID: 1, Name: "Device 1", Power: 100.0},
		{ID: 2, UserID: 1, Name: "Device 2", Power: 200.0},
	}
	mockService.On("FindAll", 1).Return(mockDevices, nil)

	c, rec := setupContext(e, http.MethodGet, "/devices", nil)
	c.Set("user", generateJWTToken(1))

	err := controller.GetDevicesController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var resp struct {
		Data []response.DeviceResponse `json:"data"`
	}
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Len(t, resp.Data, 2)
	assert.Equal(t, "Device 1", resp.Data[0].Name)
}
