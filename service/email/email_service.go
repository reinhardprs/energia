package email

import (
	"energia/entities"
	device_usage "energia/repository/device-usage"
	"energia/service/device" // Import DeviceService
	"fmt"
	"net/smtp"
	"os"
	"time"
)

type EmailService struct {
	DeviceUsageRepo device_usage.DeviceUsageRepoInterface
	DeviceService   device.DeviceInterface // Tambahkan DeviceService
}

// NewEmailService membuat instance dari EmailService
func NewEmailService(
	deviceUsageRepo device_usage.DeviceUsageRepoInterface,
	deviceService device.DeviceInterface, // Tambahkan parameter DeviceService
) *EmailService {
	return &EmailService{
		DeviceUsageRepo: deviceUsageRepo,
		DeviceService:   deviceService,
	}
}

// GenerateDeviceUsageReport menghasilkan laporan penggunaan perangkat untuk user
func (e *EmailService) GenerateDeviceUsageReport(userID int) (string, error) {
	// Ambil data penggunaan perangkat untuk hari ini
	today := time.Now()
	deviceUsages, err := e.DeviceUsageRepo.GetDeviceUsageByDate(userID, today)
	if err != nil {
		return "", fmt.Errorf("gagal mengambil data penggunaan perangkat: %v", err)
	}

	// Ambil daftar perangkat untuk user
	devices, err := e.DeviceService.FindAll(userID)
	if err != nil {
		return "", fmt.Errorf("gagal mengambil data perangkat: %v", err)
	}

	// Buat map perangkat untuk mempermudah pencarian nama perangkat
	deviceMap := make(map[int]string)
	for _, device := range devices {
		deviceMap[device.ID] = device.Name
	}

	// Format data menjadi laporan
	report := "Laporan Penggunaan Perangkat Hari Ini\n\n"
	for _, usage := range deviceUsages {
		var durationFormatted string
		if usage.Duration >= 60 {
			hours := int(usage.Duration) / 60
			minutes := int(usage.Duration) % 60
			durationFormatted = fmt.Sprintf("%d hours %d minutes", hours, minutes)
		} else {
			durationFormatted = fmt.Sprintf("%.2f minutes", usage.Duration)
		}

		// Gunakan nama perangkat jika tersedia, jika tidak gunakan ID
		deviceName := deviceMap[usage.DeviceID]
		if deviceName == "" {
			deviceName = fmt.Sprintf("Device %d", usage.DeviceID)
		}

		report += fmt.Sprintf(
			"Device: %s\nStart Time: %s\nEnd Time: %s\nDuration: %s\nEnergy Consumed: %.2f kWh\n\n",
			deviceName,
			usage.StartTime.Format("2006-01-02 15:04:05"),
			usage.EndTime.Format("2006-01-02 15:04:05"),
			durationFormatted,
			usage.EnergyConsumed,
		)
	}

	return report, nil
}

// SendEmail mengirimkan email ke alamat yang ditentukan
func (e *EmailService) SendEmail(email entities.Email) error {
	// Setup autentikasi email
	auth := smtp.PlainAuth(
		"",
		os.Getenv("MAIL_USER"),
		os.Getenv("MAIL_PASSWORD"),
		os.Getenv("MAIL_HOST"),
	)

	// Kirim email menggunakan struktur entities.Email
	msg := []byte("Subject: " + email.Subject + "\r\n\r\n" + email.Body)
	err := smtp.SendMail(
		os.Getenv("MAIL_HOST")+":"+os.Getenv("MAIL_PORT"),
		auth,
		os.Getenv("MAIL_USER"),
		[]string{email.To}, // Menggunakan email.To untuk penerima
		msg,
	)

	return err
}
