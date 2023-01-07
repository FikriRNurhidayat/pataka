// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	mock "github.com/stretchr/testify/mock"
)

// Authenticatable is an autogenerated mock type for the Authenticatable type
type Authenticatable struct {
	mock.Mock
}

// Call provides a mock function with given fields: ctx, scopes
func (_m *Authenticatable) Call(ctx context.Context, scopes ...string) (*domain.Claims, error) {
	_va := make([]interface{}, len(scopes))
	for _i := range scopes {
		_va[_i] = scopes[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *domain.Claims
	if rf, ok := ret.Get(0).(func(context.Context, ...string) *domain.Claims); ok {
		r0 = rf(ctx, scopes...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Claims)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...string) error); ok {
		r1 = rf(ctx, scopes...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetClaims provides a mock function with given fields: tokenString
func (_m *Authenticatable) GetClaims(tokenString string) (*domain.Claims, error) {
	ret := _m.Called(tokenString)

	var r0 *domain.Claims
	if rf, ok := ret.Get(0).(func(string) *domain.Claims); ok {
		r0 = rf(tokenString)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Claims)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tokenString)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetToken provides a mock function with given fields: ctx
func (_m *Authenticatable) GetToken(ctx context.Context) (string, error) {
	ret := _m.Called(ctx)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsScopeAllowed provides a mock function with given fields: claims, scopes
func (_m *Authenticatable) IsScopeAllowed(claims *domain.Claims, scopes ...string) bool {
	_va := make([]interface{}, len(scopes))
	for _i := range scopes {
		_va[_i] = scopes[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, claims)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*domain.Claims, ...string) bool); ok {
		r0 = rf(claims, scopes...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewAuthenticatable interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthenticatable creates a new instance of Authenticatable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthenticatable(t mockConstructorTestingTNewAuthenticatable) *Authenticatable {
	mock := &Authenticatable{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}