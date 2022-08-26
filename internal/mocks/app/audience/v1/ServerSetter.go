// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	audience "github.com/fikrirnurhidayat/ffgo/internal/app/audience/v1"
	mock "github.com/stretchr/testify/mock"
)

// ServerSetter is an autogenerated mock type for the ServerSetter type
type ServerSetter struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *ServerSetter) Execute(_a0 *audience.Server) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewServerSetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewServerSetter creates a new instance of ServerSetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewServerSetter(t mockConstructorTestingTNewServerSetter) *ServerSetter {
	mock := &ServerSetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}