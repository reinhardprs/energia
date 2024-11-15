package device_test

import (
	"energia/entities"
	"energia/mocks"
	"energia/service/device"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeviceService_Create_Success(t *testing.T) {
	mockDeviceRepo := new(mocks.DeviceRepoInterface)
	deviceService := device.NewDeviceService(mockDeviceRepo)

	mockDevice := entities.Device{
		UserID: 1,
		Name:   "Smart Lamp",
		Power:  10.5,
	}

	mockDeviceRepo.On("Create", 1, mock.MatchedBy(func(d entities.Device) bool {
		return d.Name == "Smart Lamp" && d.Power == 10.5
	})).Return(mockDevice, nil)

	createdDevice, err := deviceService.Create(1, mockDevice)

	assert.NoError(t, err)
	assert.Equal(t, "Smart Lamp", createdDevice.Name)
	assert.Equal(t, float32(10.5), createdDevice.Power)
	mockDeviceRepo.AssertExpectations(t)
}

func TestDeviceService_FindAll_Success(t *testing.T) {
	mockDeviceRepo := new(mocks.DeviceRepoInterface)
	deviceService := device.NewDeviceService(mockDeviceRepo)

	mockDevices := []entities.Device{
		{UserID: 1, Name: "Smart Lamp", Power: 10.5},
		{UserID: 1, Name: "Smart AC", Power: 100.0},
	}

	mockDeviceRepo.On("FindAll", 1).Return(mockDevices, nil)

	devices, err := deviceService.FindAll(1)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(devices))
	assert.Equal(t, "Smart Lamp", devices[0].Name)
	assert.Equal(t, "Smart AC", devices[1].Name)
	mockDeviceRepo.AssertExpectations(t)
}

func TestDeviceService_FindByID_Success(t *testing.T) {
	mockDeviceRepo := new(mocks.DeviceRepoInterface)
	deviceService := device.NewDeviceService(mockDeviceRepo)

	mockDevice := entities.Device{
		UserID: 1,
		Name:   "Smart Lamp",
		Power:  10.5,
	}

	mockDeviceRepo.On("FindByID", 1, 123).Return(mockDevice, nil)

	device, err := deviceService.FindByID(1, 123)

	assert.NoError(t, err)
	assert.Equal(t, "Smart Lamp", device.Name)
	assert.Equal(t, float32(10.5), device.Power)
	mockDeviceRepo.AssertExpectations(t)
}

func TestDeviceService_Update_Success(t *testing.T) {
	mockDeviceRepo := new(mocks.DeviceRepoInterface)
	deviceService := device.NewDeviceService(mockDeviceRepo)

	mockDevice := entities.Device{
		UserID: 1,
		Name:   "Smart Lamp",
		Power:  12.0,
	}

	mockDeviceRepo.On("Update", 1, mock.MatchedBy(func(d entities.Device) bool {
		return d.Name == "Smart Lamp" && d.Power == 12.0
	})).Return(mockDevice, nil)

	updatedDevice, err := deviceService.Update(1, mockDevice)

	assert.NoError(t, err)
	assert.Equal(t, "Smart Lamp", updatedDevice.Name)
	assert.Equal(t, float32(12.0), updatedDevice.Power)
	mockDeviceRepo.AssertExpectations(t)
}

func TestDeviceService_Delete_Success(t *testing.T) {
	mockDeviceRepo := new(mocks.DeviceRepoInterface)
	deviceService := device.NewDeviceService(mockDeviceRepo)

	mockDeviceRepo.On("Delete", 1, 123).Return(nil)

	err := deviceService.Delete(1, 123)

	assert.NoError(t, err)
	mockDeviceRepo.AssertExpectations(t)
}
