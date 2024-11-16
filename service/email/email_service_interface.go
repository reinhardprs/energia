package email

import "energia/entities"

// EmailServiceInterface mendefinisikan interface untuk service email
type EmailServiceInterface interface {
	GenerateDeviceUsageReport(userID int) (string, error)
	SendEmail(email entities.Email) error // Menggunakan entities.Email
}
