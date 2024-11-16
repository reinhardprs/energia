package email

import (
	"energia/controller/email/response" // Path untuk response
	"energia/entities"
	"energia/service/email" // Path untuk email service
	"net/http"

	"github.com/golang-jwt/jwt/v5" // JWT untuk mengambil user dari context
	"github.com/labstack/echo/v4"  // Echo untuk HTTP routing
)

// EmailController menangani pengiriman email
type EmailController struct {
	EmailService email.EmailServiceInterface
}

// NewEmailController membuat instance dari EmailController
func NewEmailController(emailService email.EmailServiceInterface) *EmailController {
	return &EmailController{
		EmailService: emailService,
	}
}

// SendDeviceUsageReportHandler menangani request untuk mengirim laporan penggunaan perangkat
func (e *EmailController) SendDeviceUsageReportHandler(c echo.Context) error {
	// Ambil token user dari konteks
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)

	// Ambil email dari token
	email := claims["email"].(string)

	// Ambil userID dari token
	userID := int(claims["userID"].(float64))

	// Panggil service untuk menghasilkan laporan penggunaan perangkat
	report, err := e.EmailService.GenerateDeviceUsageReport(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Siapkan objek email menggunakan entities.Email
	emailContent := entities.Email{
		To:      email,                                   // Alamat email penerima
		Subject: "Laporan Penggunaan Perangkat Hari Ini", // Subjek email
		Body:    report,                                  // Isi laporan sebagai body email
	}

	// Kirim email menggunakan service
	err = e.EmailService.SendEmail(emailContent)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Response sukses
	response := response.EmailResponse{
		Message: "Laporan email telah dikirim ke " + email,
	}
	return c.JSON(http.StatusOK, response)
}
