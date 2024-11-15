package user_usage_test

import (
	"bytes"
	"encoding/json"
	user_usage "energia/controller/user-usage"
	"energia/controller/user-usage/response"
	"energia/entities"
	"energia/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestCreateUserUsageController(t *testing.T) {
	e := echo.New()
	mockUserUsageService := new(mocks.UserUsageInterface)
	controller := user_usage.NewUserUsageController(mockUserUsageService)

	reqBody := `{"date": "2024-11-01"}`
	req := httptest.NewRequest(http.MethodPost, "/user-usage", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	claims := jwt.MapClaims{"userID": float64(1)}
	token := &jwt.Token{Claims: claims}
	c.Set("user", token)

	mockUserUsageService.On("Create", 1, mock.Anything).Return(entities.UserUsage{
		UserID:      1,
		TotalEnergy: 100.0,
		TotalCost:   135200.0,
	}, nil)

	if assert.NoError(t, controller.CreateUserUsageController(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestFindUserUsageController(t *testing.T) {
	e := echo.New()
	mockService := new(mocks.UserUsageInterface)
	controller := user_usage.NewUserUsageController(mockService)

	mockDate := time.Now().Format("2006-01-02")
	mockParsedDate, _ := time.Parse("2006-01-02", mockDate)

	mockUserUsages := []entities.UserUsage{
		{
			ID:          1,
			UserID:      1,
			Date:        mockParsedDate,
			TotalEnergy: 100.0,
			TotalCost:   135200.0,
		},
		{
			ID:          2,
			UserID:      1,
			Date:        mockParsedDate,
			TotalEnergy: 150.0,
			TotalCost:   202800.0,
		},
	}

	mockService.On("GetUserUsage", 1).Return(mockUserUsages, nil)

	c, rec := setupContext(e, http.MethodGet, "/user-usage", nil)
	c.Set("user", generateJWTToken(1))

	err := controller.FindUserUsageController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var resp []response.UserUsageResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Len(t, resp, 2)

	assert.Equal(t, 1, resp[0].UserID)
	assert.InDelta(t, 100.0, resp[0].TotalEnergy, 0.01)
	assert.InDelta(t, 135200.0, resp[0].TotalCost, 0.01)

	assert.Equal(t, 1, resp[1].UserID)
	assert.InDelta(t, 150.0, resp[1].TotalEnergy, 0.01)
	assert.InDelta(t, 202800.0, resp[1].TotalCost, 0.01)
}
