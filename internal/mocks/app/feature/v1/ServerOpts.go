// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	feature "github.com/fikrirnurhidayat/ffgo/internal/app/feature/v1"
	mock "github.com/stretchr/testify/mock"
)

// ServerOpts is an autogenerated mock type for the ServerOpts type
type ServerOpts struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *ServerOpts) Execute(_a0 *feature.Server) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewServerOpts interface {
	mock.TestingT
	Cleanup(func())
}

// NewServerOpts creates a new instance of ServerOpts. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewServerOpts(t mockConstructorTestingTNewServerOpts) *ServerOpts {
	mock := &ServerOpts{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
