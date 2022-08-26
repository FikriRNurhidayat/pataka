// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	mock "github.com/stretchr/testify/mock"
)

// AudienceGetable is an autogenerated mock type for the AudienceGetable type
type AudienceGetable struct {
	mock.Mock
}

// Call provides a mock function with given fields: _a0, _a1
func (_m *AudienceGetable) Call(_a0 context.Context, _a1 *domain.GetAudienceParams) (*domain.GetAudienceResult, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *domain.GetAudienceResult
	if rf, ok := ret.Get(0).(func(context.Context, *domain.GetAudienceParams) *domain.GetAudienceResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.GetAudienceResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.GetAudienceParams) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAudienceGetable interface {
	mock.TestingT
	Cleanup(func())
}

// NewAudienceGetable creates a new instance of AudienceGetable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAudienceGetable(t mockConstructorTestingTNewAudienceGetable) *AudienceGetable {
	mock := &AudienceGetable{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
