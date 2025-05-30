// Code generated by mockery v2.51.1. DO NOT EDIT.

package authmocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// TokenExtractor is an autogenerated mock type for the TokenExtractor type
type TokenExtractor struct {
	mock.Mock
}

// GetAPIKey provides a mock function with given fields: headers
func (_m *TokenExtractor) GetAPIKey(headers http.Header) (string, error) {
	ret := _m.Called(headers)

	if len(ret) == 0 {
		panic("no return value specified for GetAPIKey")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(http.Header) (string, error)); ok {
		return rf(headers)
	}
	if rf, ok := ret.Get(0).(func(http.Header) string); ok {
		r0 = rf(headers)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(http.Header) error); ok {
		r1 = rf(headers)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBearerToken provides a mock function with given fields: headers
func (_m *TokenExtractor) GetBearerToken(headers http.Header) (string, error) {
	ret := _m.Called(headers)

	if len(ret) == 0 {
		panic("no return value specified for GetBearerToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(http.Header) (string, error)); ok {
		return rf(headers)
	}
	if rf, ok := ret.Get(0).(func(http.Header) string); ok {
		r0 = rf(headers)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(http.Header) error); ok {
		r1 = rf(headers)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTokenExtractor creates a new instance of TokenExtractor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTokenExtractor(t interface {
	mock.TestingT
	Cleanup(func())
}) *TokenExtractor {
	mock := &TokenExtractor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
