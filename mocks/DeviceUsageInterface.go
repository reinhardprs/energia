// Code generated by mockery v2.47.0. DO NOT EDIT.

package mocks

import (
	entities "energia/entities"

	mock "github.com/stretchr/testify/mock"
)

// DeviceUsageInterface is an autogenerated mock type for the DeviceUsageInterface type
type DeviceUsageInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: deviceUsage, userID
func (_m *DeviceUsageInterface) Create(deviceUsage entities.DeviceUsage, userID int) (entities.DeviceUsage, error) {
	ret := _m.Called(deviceUsage, userID)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 entities.DeviceUsage
	var r1 error
	if rf, ok := ret.Get(0).(func(entities.DeviceUsage, int) (entities.DeviceUsage, error)); ok {
		return rf(deviceUsage, userID)
	}
	if rf, ok := ret.Get(0).(func(entities.DeviceUsage, int) entities.DeviceUsage); ok {
		r0 = rf(deviceUsage, userID)
	} else {
		r0 = ret.Get(0).(entities.DeviceUsage)
	}

	if rf, ok := ret.Get(1).(func(entities.DeviceUsage, int) error); ok {
		r1 = rf(deviceUsage, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields: userID
func (_m *DeviceUsageInterface) FindAll(userID int) ([]entities.DeviceUsage, error) {
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

// NewDeviceUsageInterface creates a new instance of DeviceUsageInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDeviceUsageInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *DeviceUsageInterface {
	mock := &DeviceUsageInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
