// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// FeatureResultable is an autogenerated mock type for the FeatureResultable type
type FeatureResultable struct {
	mock.Mock
}

type mockConstructorTestingTNewFeatureResultable interface {
	mock.TestingT
	Cleanup(func())
}

// NewFeatureResultable creates a new instance of FeatureResultable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFeatureResultable(t mockConstructorTestingTNewFeatureResultable) *FeatureResultable {
	mock := &FeatureResultable{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
