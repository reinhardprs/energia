package device_usage_test

import (
	"energia/entities"
	"energia/mocks"
	device_usage "energia/service/device-usage"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeviceUsageService_Create_Success(t *testing.T) {
	mockDeviceRepo := new(mocks.DeviceRepoInterface)
	mockDeviceUsageRepo := new(mocks.DeviceUsageRepoInterface)
	deviceUsageService := device_usage.NewDeviceUsageService(mockDeviceRepo, mockDeviceUsageRepo)

	// Mock device data
	mockDevice := entities.Device{
		ID:     1,
		UserID: 1,
		Name:   "Smart Lamp",
		Power:  100,
	}

	// Mock device usage data with specific start and end times
	startTime := time.Now()
	endTime := startTime.Add(1 * time.Hour) // 1-hour duration for simplicity
	mockDeviceUsage := entities.DeviceUsage{
		DeviceID:       mockDevice.ID,
		StartTime:      startTime,
		EndTime:        endTime,
		Duration:       60, // in minutes
		EnergyConsumed: 6,  // calculated as 100W * 60 mins / 1000
	}

	// Define mocks
	mockDeviceRepo.On("FindByID", 1, 1).Return(mockDevice, nil)
	mockDeviceUsageRepo.On("Create", mock.AnythingOfType("entities.DeviceUsage")).Return(mockDeviceUsage, nil)

	// Call the service method
	createdDeviceUsage, err := deviceUsageService.Create(mockDeviceUsage, 1)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, float32(60), createdDeviceUsage.Duration)
	assert.Equal(t, float32(6), createdDeviceUsage.EnergyConsumed)
	mockDeviceRepo.AssertExpectations(t)
	mockDeviceUsageRepo.AssertExpectations(t)
}

func TestDeviceUsageService_FindAll_Success(t *testing.T) {
	mockDeviceUsageRepo := new(mocks.DeviceUsageRepoInterface)
	deviceUsageService := device_usage.NewDeviceUsageService(nil, mockDeviceUsageRepo)

	mockDeviceUsages := []entities.DeviceUsage{
		{ID: 1, DeviceID: 1, Duration: 60, EnergyConsumed: 6},
		{ID: 2, DeviceID: 2, Duration: 45, EnergyConsumed: 4.5},
	}

	mockDeviceUsageRepo.On("FindAll", 1).Return(mockDeviceUsages, nil)

	foundDeviceUsages, err := deviceUsageService.FindAll(1)

	assert.NoError(t, err)
	assert.Equal(t, mockDeviceUsages, foundDeviceUsages)
	mockDeviceUsageRepo.AssertExpectations(t)
}
