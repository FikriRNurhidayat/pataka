// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// AudienceProtoable is an autogenerated mock type for the AudienceProtoable type
type AudienceProtoable struct {
	mock.Mock
}

type mockConstructorTestingTNewAudienceProtoable interface {
	mock.TestingT
	Cleanup(func())
}

// NewAudienceProtoable creates a new instance of AudienceProtoable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAudienceProtoable(t mockConstructorTestingTNewAudienceProtoable) *AudienceProtoable {
	mock := &AudienceProtoable{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}