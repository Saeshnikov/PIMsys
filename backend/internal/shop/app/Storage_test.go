// Code generated by mockery v2.52.4. DO NOT EDIT.

package shop_app_test

import (
	context "context"
	proto "pim-sys/gen/go/shop"

	mock "github.com/stretchr/testify/mock"
)

// Storage is an autogenerated mock type for the Storage type
type Storage struct {
	mock.Mock
}

// AlterShop provides a mock function with given fields: ctx, shopId, name, description, url
func (_m *Storage) AlterShop(ctx context.Context, shopId int32, name string, description string, url string) error {
	ret := _m.Called(ctx, shopId, name, description, url)

	if len(ret) == 0 {
		panic("no return value specified for AlterShop")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32, string, string, string) error); ok {
		r0 = rf(ctx, shopId, name, description, url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateShop provides a mock function with given fields: ctx, userId, name, description, url
func (_m *Storage) CreateShop(ctx context.Context, userId int, name string, description string, url string) error {
	ret := _m.Called(ctx, userId, name, description, url)

	if len(ret) == 0 {
		panic("no return value specified for CreateShop")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string, string, string) error); ok {
		r0 = rf(ctx, userId, name, description, url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteShop provides a mock function with given fields: ctx, shopId
func (_m *Storage) DeleteShop(ctx context.Context, shopId int32) error {
	ret := _m.Called(ctx, shopId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteShop")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) error); ok {
		r0 = rf(ctx, shopId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListShops provides a mock function with given fields: ctx, userId
func (_m *Storage) ListShops(ctx context.Context, userId int32) ([]*proto.ShopInfo, error) {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for ListShops")
	}

	var r0 []*proto.ShopInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) ([]*proto.ShopInfo, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32) []*proto.ShopInfo); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*proto.ShopInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewStorage creates a new instance of Storage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *Storage {
	mock := &Storage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
