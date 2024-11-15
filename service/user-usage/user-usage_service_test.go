package user_usage_test

import (
	"energia/entities"
	"energia/mocks"
	user_usage "energia/service/user-usage"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserUsageService_GetUserUsage_Success(t *testing.T) {
	mockUserUsageRepo := new(mocks.UserUsageRepoInterface)
	mockDeviceUsageRepo := new(mocks.DeviceUsageRepoInterface)
	userUsageService := user_usage.NewUserUsageService(mockUserUsageRepo, mockDeviceUsageRepo)

	userID := 1
	expectedUserUsages := []entities.UserUsage{
		{ID: 1, UserID: userID, Date: time.Now(), TotalEnergy: 100, TotalCost: 135200},
		{ID: 2, UserID: userID, Date: time.Now(), TotalEnergy: 150, TotalCost: 202800},
	}

	// Set up mock expectations
	mockUserUsageRepo.On("FindAll", userID).Return(expectedUserUsages, nil)

	// Call the service function
	userUsages, err := userUsageService.GetUserUsage(userID)

	// Assert results
	assert.NoError(t, err)
	assert.Equal(t, expectedUserUsages, userUsages)
	mockUserUsageRepo.AssertExpectations(t)
}

func TestUserUsageService_Create_Success(t *testing.T) {
	mockUserUsageRepo := new(mocks.UserUsageRepoInterface)
	mockDeviceUsageRepo := new(mocks.DeviceUsageRepoInterface)
	userUsageService := user_usage.NewUserUsageService(mockUserUsageRepo, mockDeviceUsageRepo)

	userID := 1
	date := time.Now()
	deviceUsages := []entities.DeviceUsage{
		{ID: 1, DeviceID: 1, StartTime: date.Add(-2 * time.Hour), EndTime: date, Duration: 120, EnergyConsumed: 60},
		{ID: 2, DeviceID: 2, StartTime: date.Add(-1 * time.Hour), EndTime: date, Duration: 60, EnergyConsumed: 30},
	}
	expectedTotalEnergy := float32(90)              // Sum of energy consumed
	expectedTotalCost := expectedTotalEnergy * 1352 // Assume 1352 is the cost per kWh

	// Create expected result
	expectedUserUsage := entities.UserUsage{
		UserID:      userID,
		Date:        date,
		TotalEnergy: expectedTotalEnergy,
		TotalCost:   expectedTotalCost,
	}

	// Set up mock expectations
	mockDeviceUsageRepo.On("GetDeviceUsageByDate", userID, date).Return(deviceUsages, nil)
	mockUserUsageRepo.On("Create", mock.MatchedBy(func(u entities.UserUsage) bool {
		return u.UserID == userID && u.Date.Equal(date) &&
			u.TotalEnergy == expectedTotalEnergy && u.TotalCost == expectedTotalCost
	})).Return(expectedUserUsage, nil)

	// Call the service function
	createdUserUsage, err := userUsageService.Create(userID, date)

	// Assert results
	assert.NoError(t, err)
	assert.Equal(t, expectedUserUsage, createdUserUsage)
	mockDeviceUsageRepo.AssertExpectations(t)
	mockUserUsageRepo.AssertExpectations(t)
}
