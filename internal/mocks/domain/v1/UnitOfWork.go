// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	mock "github.com/stretchr/testify/mock"
)

// UnitOfWork is an autogenerated mock type for the UnitOfWork type
type UnitOfWork struct {
	mock.Mock
}

// Do provides a mock function with given fields: _a0, _a1
func (_m *UnitOfWork) Do(_a0 context.Context, _a1 domain.Block) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Block) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUnitOfWork interface {
	mock.TestingT
	Cleanup(func())
}

// NewUnitOfWork creates a new instance of UnitOfWork. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUnitOfWork(t mockConstructorTestingTNewUnitOfWork) *UnitOfWork {
	mock := &UnitOfWork{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
