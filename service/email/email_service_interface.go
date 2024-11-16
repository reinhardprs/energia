package email

import "energia/entities"

type EmailServiceInterface interface {
	GenerateDeviceUsageReport(userID int) (string, error)
	SendEmail(email entities.Email) error
}
