package email_test

import (
	"energia/entities"
	"energia/mocks"
	"energia/service/email"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGenerateDeviceUsageReport(t *testing.T) {
	mockDeviceUsageRepo := new(mocks.DeviceUsageRepoInterface)
	mockDeviceService := new(mocks.DeviceInterface)

	emailService := email.NewEmailService(mockDeviceUsageRepo, mockDeviceService)

	today := time.Now()
	mockDeviceUsageRepo.On("GetDeviceUsageByDate", mock.AnythingOfType("int"), mock.MatchedBy(func(date time.Time) bool {
		return date.Year() == today.Year() && date.Month() == today.Month() && date.Day() == today.Day()
	})).Return([]entities.DeviceUsage{
		{DeviceID: 1, StartTime: today.Add(-2 * time.Hour), EndTime: today.Add(-time.Hour), Duration: 120, EnergyConsumed: 1.5},
		{DeviceID: 2, StartTime: today.Add(-3 * time.Hour), EndTime: today.Add(-2 * time.Hour), Duration: 60, EnergyConsumed: 0.8},
	}, nil)

	mockDeviceService.On("FindAll", mock.AnythingOfType("int")).Return([]entities.Device{
		{ID: 1, Name: "Smart AC"},
		{ID: 2, Name: "Eco Heater"},
	}, nil)

	report, err := emailService.GenerateDeviceUsageReport(1)

	assert.NoError(t, err)
	assert.Contains(t, report, "Device: Smart AC")
	assert.Contains(t, report, "Device: Eco Heater")
}

func TestGenerateDeviceUsageReport_ErrorDeviceUsage(t *testing.T) {
	mockDeviceUsageRepo := new(mocks.DeviceUsageRepoInterface)
	mockDeviceService := new(mocks.DeviceInterface)

	emailService := email.NewEmailService(mockDeviceUsageRepo, mockDeviceService)

	mockDeviceUsageRepo.On("GetDeviceUsageByDate", mock.AnythingOfType("int"), mock.Anything).Return(nil, errors.New("error retrieving device usage"))

	_, err := emailService.GenerateDeviceUsageReport(1)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "gagal mengambil data penggunaan perangkat")
}

func TestGenerateDeviceUsageReport_ErrorDeviceService(t *testing.T) {
	mockDeviceUsageRepo := new(mocks.DeviceUsageRepoInterface)
	mockDeviceService := new(mocks.DeviceInterface)

	emailService := email.NewEmailService(mockDeviceUsageRepo, mockDeviceService)

	mockDeviceUsageRepo.On("GetDeviceUsageByDate", mock.AnythingOfType("int"), mock.Anything).Return([]entities.DeviceUsage{
		{DeviceID: 1, StartTime: time.Now(), EndTime: time.Now(), Duration: 30, EnergyConsumed: 0.5},
	}, nil)

	mockDeviceService.On("FindAll", mock.AnythingOfType("int")).Return(nil, errors.New("error retrieving devices"))

	_, err := emailService.GenerateDeviceUsageReport(1)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "gagal mengambil data perangkat")
}
