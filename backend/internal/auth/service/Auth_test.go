// Code generated by mockery v2.52.4. DO NOT EDIT.

package auth_service_test

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Auth is an autogenerated mock type for the Auth type
type Auth struct {
	mock.Mock
}

// IsAdmin provides a mock function with given fields: ctx, userID
func (_m *Auth) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for IsAdmin")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (bool, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) bool); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, email, password
func (_m *Auth) Login(ctx context.Context, email string, password string) (string, error) {
	ret := _m.Called(ctx, email, password)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (string, error)); ok {
		return rf(ctx, email, password)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, email, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterNewUser provides a mock function with given fields: ctx, email, password, name, phone
func (_m *Auth) RegisterNewUser(ctx context.Context, email string, password string, name string, phone string) (int64, error) {
	ret := _m.Called(ctx, email, password, name, phone)

	if len(ret) == 0 {
		panic("no return value specified for RegisterNewUser")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) (int64, error)); ok {
		return rf(ctx, email, password, name, phone)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) int64); ok {
		r0 = rf(ctx, email, password, name, phone)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string) error); ok {
		r1 = rf(ctx, email, password, name, phone)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuth creates a new instance of Auth. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuth(t interface {
	mock.TestingT
	Cleanup(func())
}) *Auth {
	mock := &Auth{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
