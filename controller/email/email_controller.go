package email

import (
	"energia/controller/email/response"
	"energia/entities"
	"energia/service/email"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type EmailController struct {
	EmailService email.EmailServiceInterface
}

func NewEmailController(emailService email.EmailServiceInterface) *EmailController {
	return &EmailController{
		EmailService: emailService,
	}
}

// SendDeviceUsageReportHandler is the controller for the send device usage report endpoint
// @Summary Send device usage report
// @Description Send device usage report
// @Tags Email
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.EmailResponse
// @Failure 400 {object} base.BaseResponse
// @Router /report [get]
func (e *EmailController) SendDeviceUsageReportHandler(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)

	email := claims["email"].(string)

	userID := int(claims["userID"].(float64))

	report, err := e.EmailService.GenerateDeviceUsageReport(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	emailContent := entities.Email{
		To:      email,
		Subject: "Laporan Penggunaan Perangkat Hari Ini",
		Body:    report,
	}

	err = e.EmailService.SendEmail(emailContent)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := response.EmailResponse{
		Message: "Laporan email telah dikirim ke " + email,
	}
	return c.JSON(http.StatusOK, response)
}
