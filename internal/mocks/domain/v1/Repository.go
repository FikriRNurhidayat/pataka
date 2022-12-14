// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// AudienceRepository provides a mock function with given fields:
func (_m *Repository) AudienceRepository() domain.AudienceRepository {
	ret := _m.Called()

	var r0 domain.AudienceRepository
	if rf, ok := ret.Get(0).(func() domain.AudienceRepository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.AudienceRepository)
		}
	}

	return r0
}

// FeatureRepository provides a mock function with given fields:
func (_m *Repository) FeatureRepository() domain.FeatureRepository {
	ret := _m.Called()

	var r0 domain.FeatureRepository
	if rf, ok := ret.Get(0).(func() domain.FeatureRepository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.FeatureRepository)
		}
	}

	return r0
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
