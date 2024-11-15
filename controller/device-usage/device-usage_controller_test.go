package device_usage_test

import (
	"bytes"
	"encoding/json"
	device_usage "energia/controller/device-usage"
	"energia/controller/device-usage/request"
	"energia/controller/device-usage/response"
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

func TestCreateDeviceUsageController(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.DeviceUsageInterface)
	controller := device_usage.NewDeviceUsageController(mockService)

	startTime := time.Now()
	endTime := startTime.Add(2 * time.Hour)
	mockDeviceUsage := entities.DeviceUsage{
		ID:             1,
		DeviceID:       1,
		StartTime:      startTime,
		EndTime:        endTime,
		Duration:       120.0,
		EnergyConsumed: 200.0,
	}

	mockService.On("Create", mock.AnythingOfType("entities.DeviceUsage"), 1).Return(mockDeviceUsage, nil)

	body := request.CreateDeviceUsageRequest{
		DeviceID:  1,
		StartTime: startTime,
		EndTime:   endTime,
	}
	bodyBytes, _ := json.Marshal(body)

	c, rec := setupContext(e, http.MethodPost, "/device-usage", bodyBytes)
	c.Set("user", generateJWTToken(1))

	err := controller.CreateDeviceUsageController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Memeriksa response
	var resp struct {
		Data response.DeviceUsageResponse `json:"data"`
	}
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, 1, resp.Data.ID)
	assert.Equal(t, 1, resp.Data.DeviceID)
	assert.Equal(t, float32(120.0), resp.Data.Duration)
	assert.Equal(t, float32(200.0), resp.Data.EnergyConsumed)
}

func TestFindAllDeviceUsageController(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.DeviceUsageInterface)
	controller := device_usage.NewDeviceUsageController(mockService)

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

	mockService.On("FindAll", 1).Return(mockDeviceUsages, nil)

	c, rec := setupContext(e, http.MethodGet, "/device-usage", nil)
	c.Set("user", generateJWTToken(1))

	err := controller.FindAllDeviceUsageController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var resp struct {
		Data []response.DeviceUsageResponse `json:"data"`
	}
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Len(t, resp.Data, 2)
	assert.Equal(t, 1, resp.Data[0].ID)
	assert.Equal(t, 2, resp.Data[1].ID)
}
