// Code generated by mockery v2.51.1. DO NOT EDIT.

package mocks

import (
	context "context"

	database "github.com/seanhuebl/unity-wealth/internal/database"
	mock "github.com/stretchr/testify/mock"
)

// DeviceQuerier is an autogenerated mock type for the DeviceQuerier type
type DeviceQuerier struct {
	mock.Mock
}

// CreateDeviceInfo provides a mock function with given fields: ctx, arg
func (_m *DeviceQuerier) CreateDeviceInfo(ctx context.Context, arg database.CreateDeviceInfoParams) (string, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateDeviceInfo")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, database.CreateDeviceInfoParams) (string, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, database.CreateDeviceInfoParams) string); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, database.CreateDeviceInfoParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceInfoByUser provides a mock function with given fields: ctx, arg
func (_m *DeviceQuerier) GetDeviceInfoByUser(ctx context.Context, arg database.GetDeviceInfoByUserParams) (string, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for GetDeviceInfoByUser")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, database.GetDeviceInfoByUserParams) (string, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, database.GetDeviceInfoByUserParams) string); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, database.GetDeviceInfoByUserParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDeviceQuerier creates a new instance of DeviceQuerier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDeviceQuerier(t interface {
	mock.TestingT
	Cleanup(func())
}) *DeviceQuerier {
	mock := &DeviceQuerier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
