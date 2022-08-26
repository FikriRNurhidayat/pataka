// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	mock "github.com/stretchr/testify/mock"
)

// AudienceRepository is an autogenerated mock type for the AudienceRepository type
type AudienceRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *AudienceRepository) Delete(_a0 context.Context, _a1 *domain.Audience) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Audience) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBy provides a mock function with given fields: ctx, args
func (_m *AudienceRepository) DeleteBy(ctx context.Context, args *domain.AudienceFilterArgs) error {
	ret := _m.Called(ctx, args)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.AudienceFilterArgs) error); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, fn, ui
func (_m *AudienceRepository) Get(ctx context.Context, fn string, ui string) (*domain.Audience, error) {
	ret := _m.Called(ctx, fn, ui)

	var r0 *domain.Audience
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *domain.Audience); ok {
		r0 = rf(ctx, fn, ui)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Audience)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, fn, ui)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBy provides a mock function with given fields: ctx, args
func (_m *AudienceRepository) GetBy(ctx context.Context, args *domain.AudienceFilterArgs) (*domain.Audience, error) {
	ret := _m.Called(ctx, args)

	var r0 *domain.Audience
	if rf, ok := ret.Get(0).(func(context.Context, *domain.AudienceFilterArgs) *domain.Audience); ok {
		r0 = rf(ctx, args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Audience)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.AudienceFilterArgs) error); ok {
		r1 = rf(ctx, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: _a0, _a1
func (_m *AudienceRepository) List(_a0 context.Context, _a1 *domain.AudienceListArgs) ([]domain.Audience, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []domain.Audience
	if rf, ok := ret.Get(0).(func(context.Context, *domain.AudienceListArgs) []domain.Audience); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Audience)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.AudienceListArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: _a0, _a1
func (_m *AudienceRepository) Save(_a0 context.Context, _a1 *domain.Audience) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Audience) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Size provides a mock function with given fields: _a0, _a1
func (_m *AudienceRepository) Size(_a0 context.Context, _a1 *domain.AudienceFilterArgs) (uint32, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(context.Context, *domain.AudienceFilterArgs) uint32); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.AudienceFilterArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAudienceRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewAudienceRepository creates a new instance of AudienceRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAudienceRepository(t mockConstructorTestingTNewAudienceRepository) *AudienceRepository {
	mock := &AudienceRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
