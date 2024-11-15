// Code generated by mockery v2.47.0. DO NOT EDIT.

package mocks

import (
	entities "energia/entities"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// DeviceUsageRepoInterface is an autogenerated mock type for the DeviceUsageRepoInterface type
type DeviceUsageRepoInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: deviceUsage
func (_m *DeviceUsageRepoInterface) Create(deviceUsage entities.DeviceUsage) (entities.DeviceUsage, error) {
	ret := _m.Called(deviceUsage)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 entities.DeviceUsage
	var r1 error
	if rf, ok := ret.Get(0).(func(entities.DeviceUsage) (entities.DeviceUsage, error)); ok {
		return rf(deviceUsage)
	}
	if rf, ok := ret.Get(0).(func(entities.DeviceUsage) entities.DeviceUsage); ok {
		r0 = rf(deviceUsage)
	} else {
		r0 = ret.Get(0).(entities.DeviceUsage)
	}

	if rf, ok := ret.Get(1).(func(entities.DeviceUsage) error); ok {
		r1 = rf(deviceUsage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields: userID
func (_m *DeviceUsageRepoInterface) FindAll(userID int) ([]entities.DeviceUsage, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 []entities.DeviceUsage
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]entities.DeviceUsage, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(int) []entities.DeviceUsage); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.DeviceUsage)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceUsageByDate provides a mock function with given fields: userID, date
func (_m *DeviceUsageRepoInterface) GetDeviceUsageByDate(userID int, date time.Time) ([]entities.DeviceUsage, error) {
	ret := _m.Called(userID, date)

	if len(ret) == 0 {
		panic("no return value specified for GetDeviceUsageByDate")
	}

	var r0 []entities.DeviceUsage
	var r1 error
	if rf, ok := ret.Get(0).(func(int, time.Time) ([]entities.DeviceUsage, error)); ok {
		return rf(userID, date)
	}
	if rf, ok := ret.Get(0).(func(int, time.Time) []entities.DeviceUsage); ok {
		r0 = rf(userID, date)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.DeviceUsage)
		}
	}

	if rf, ok := ret.Get(1).(func(int, time.Time) error); ok {
		r1 = rf(userID, date)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDeviceUsageRepoInterface creates a new instance of DeviceUsageRepoInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDeviceUsageRepoInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *DeviceUsageRepoInterface {
	mock := &DeviceUsageRepoInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}