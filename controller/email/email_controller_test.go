package email_test

import (
	"energia/controller/email"
	"energia/entities"
	"energia/mocks"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSendDeviceUsageReportHandler(t *testing.T) {
	e := echo.New()
	mockEmailService := new(mocks.EmailServiceInterface)

	emailController := email.NewEmailController(mockEmailService)

	// Mock JWT token
	userClaims := jwt.MapClaims{
		"userID": float64(1), // Gunakan float64 untuk userID
		"email":  "testuser@example.com",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	// Mock Echo context
	req := httptest.NewRequest(http.MethodPost, "/send-report", nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer mockToken")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.Set("user", token)

	// Mock EmailService.GenerateDeviceUsageReport
	mockEmailService.On("GenerateDeviceUsageReport", 1).Return("Mock Report", nil)

	// Mock EmailService.SendEmail
	mockEmailService.On("SendEmail", mock.AnythingOfType("entities.Email")).Return(nil)

	// Execute handler
	err := emailController.SendDeviceUsageReportHandler(ctx)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"message": "Laporan email telah dikirim ke testuser@example.com"}`, rec.Body.String())

	// Validasi mock behavior
	mockEmailService.AssertCalled(t, "GenerateDeviceUsageReport", 1)
	mockEmailService.AssertCalled(t, "SendEmail", mock.MatchedBy(func(email entities.Email) bool {
		return email.To == "testuser@example.com" &&
			email.Subject == "Laporan Penggunaan Perangkat Hari Ini" &&
			email.Body == "Mock Report"
	}))
}

func TestSendDeviceUsageReportHandler_GenerateReportError(t *testing.T) {
	e := echo.New()
	mockEmailService := new(mocks.EmailServiceInterface)

	emailController := email.NewEmailController(mockEmailService)

	// Mock JWT token
	userClaims := jwt.MapClaims{
		"userID": float64(1), // Gunakan float64 untuk userID
		"email":  "testuser@example.com",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	// Mock Echo context
	req := httptest.NewRequest(http.MethodPost, "/send-report", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.Set("user", token)

	// Mock GenerateDeviceUsageReport dengan error
	mockEmailService.On("GenerateDeviceUsageReport", 1).Return("", errors.New("mock error"))

	// Execute handler
	err := emailController.SendDeviceUsageReportHandler(ctx)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.JSONEq(t, `{"error": "mock error"}`, rec.Body.String())
}

func TestSendDeviceUsageReportHandler_SendEmailError(t *testing.T) {
	e := echo.New()
	mockEmailService := new(mocks.EmailServiceInterface)

	emailController := email.NewEmailController(mockEmailService)

	// Mock JWT token
	userClaims := jwt.MapClaims{
		"userID": float64(1), // Gunakan float64 untuk userID
		"email":  "testuser@example.com",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	// Mock Echo context
	req := httptest.NewRequest(http.MethodPost, "/send-report", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.Set("user", token)

	// Mock GenerateDeviceUsageReport
	mockEmailService.On("GenerateDeviceUsageReport", 1).Return("Mock Report", nil)

	// Mock SendEmail dengan error
	mockEmailService.On("SendEmail", mock.AnythingOfType("entities.Email")).Return(errors.New("mock error"))

	// Execute handler
	err := emailController.SendDeviceUsageReportHandler(ctx)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.JSONEq(t, `{"error": "mock error"}`, rec.Body.String())
}
