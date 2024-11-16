package device_usage_test

import (
	"bytes"
	"encoding/json"
	device_usage "energia/controller/device-usage"
	"energia/controller/device-usage/request"
	"energia/entities"
	"energia/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Helper untuk generate token JWT
func generateJWTToken(userID int) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": float64(userID),
	})
}

// Helper untuk setup context Echo
func setupContext(e *echo.Echo, method, path string, body []byte) (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, path, bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c, rec
}

func TestCreateDeviceUsageController(t *testing.T) {
	e := echo.New()
	mockDeviceUsageService := new(mocks.DeviceUsageInterface)
	mockDeviceService := new(mocks.DeviceInterface) // Mock tambahan untuk DeviceService

	controller := device_usage.NewDeviceUsageController(mockDeviceUsageService, mockDeviceService)

	// Data mock untuk test
	startTime := time.Now()
	endTime := startTime.Add(2 * time.Hour)
	mockDevice := entities.Device{ID: 1, Name: "Device 1"}
	mockDeviceUsage := entities.DeviceUsage{
		ID:             1,
		DeviceID:       1,
		StartTime:      startTime,
		EndTime:        endTime,
		Duration:       120.0,
		EnergyConsumed: 200.0,
	}

	// Mock dependensi service
	mockDeviceService.On("FindByID", 1, 1).Return(mockDevice, nil)
	mockDeviceUsageService.On("Create", mock.AnythingOfType("entities.DeviceUsage"), 1).Return(mockDeviceUsage, nil)

	// Membuat request body
	body := request.CreateDeviceUsageRequest{
		DeviceID:  1,
		StartTime: startTime,
		EndTime:   endTime,
	}
	bodyBytes, _ := json.Marshal(body)

	// Setup context dan user token
	c, rec := setupContext(e, http.MethodPost, "/device-usage", bodyBytes)
	c.Set("user", generateJWTToken(1))

	// Eksekusi controller
	err := controller.CreateDeviceUsageController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Memeriksa response
	var resp struct {
		Data map[string]interface{} `json:"data"`
	}
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, float64(1), resp.Data["device_usage"].(map[string]interface{})["id"])
	assert.Equal(t, "Device 1", resp.Data["device_name"])
}

func TestFindAllDeviceUsageController(t *testing.T) {
	e := echo.New()
	mockDeviceUsageService := new(mocks.DeviceUsageInterface)
	mockDeviceService := new(mocks.DeviceInterface) // Mock tambahan untuk DeviceService

	controller := device_usage.NewDeviceUsageController(mockDeviceUsageService, mockDeviceService)

	// Data mock untuk test
	startTime := time.Now()
	endTime := startTime.Add(2 * time.Hour)
	mockDeviceUsages := []entities.DeviceUsage{
		{
			ID:             1,
			DeviceID:       1,
			StartTime:      startTime,
			EndTime:        endTime,
			Duration:       120.0,
			EnergyConsumed: 200.0,
		},
		{
			ID:             2,
			DeviceID:       2,
			StartTime:      startTime,
			EndTime:        endTime,
			Duration:       180.0,
			EnergyConsumed: 300.0,
		},
	}

	mockDevices := []entities.Device{
		{ID: 1, Name: "Device 1"},
		{ID: 2, Name: "Device 2"},
	}

	// Mock dependensi service
	mockDeviceUsageService.On("FindAll", 1).Return(mockDeviceUsages, nil)
	mockDeviceService.On("FindAll", 1).Return(mockDevices, nil)

	// Setup context dan user token
	c, rec := setupContext(e, http.MethodGet, "/device-usage", nil)
	c.Set("user", generateJWTToken(1))

	// Eksekusi controller
	err := controller.FindAllDeviceUsageController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Memeriksa response
	var resp struct {
		Data []map[string]interface{} `json:"data"`
	}
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Len(t, resp.Data, 2)

	// Validasi data pada response
	assert.Equal(t, float64(1), resp.Data[0]["device_usage"].(map[string]interface{})["id"])
	assert.Equal(t, "Device 1", resp.Data[0]["device_name"])
	assert.Equal(t, float64(2), resp.Data[1]["device_usage"].(map[string]interface{})["id"])
	assert.Equal(t, "Device 2", resp.Data[1]["device_name"])
}
