// Code generated by mockery v2.52.4. DO NOT EDIT.

package product_service_test

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	proto "pim-sys/gen/go/products"
)

// Product is an autogenerated mock type for the Product type
type Product struct {
	mock.Mock
}

// AlterProduct provides a mock function with given fields: ctx, content
func (_m *Product) AlterProduct(ctx context.Context, content *proto.ProductInfoWithId) error {
	ret := _m.Called(ctx, content)

	if len(ret) == 0 {
		panic("no return value specified for AlterProduct")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.ProductInfoWithId) error); ok {
		r0 = rf(ctx, content)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteProduct provides a mock function with given fields: ctx, content
func (_m *Product) DeleteProduct(ctx context.Context, content *proto.DeleteProductRequest) error {
	ret := _m.Called(ctx, content)

	if len(ret) == 0 {
		panic("no return value specified for DeleteProduct")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.DeleteProductRequest) error); ok {
		r0 = rf(ctx, content)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListProducts provides a mock function with given fields: ctx
func (_m *Product) ListProducts(ctx context.Context) ([]*proto.ProductInfoWithId, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListProducts")
	}

	var r0 []*proto.ProductInfoWithId
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*proto.ProductInfoWithId, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*proto.ProductInfoWithId); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*proto.ProductInfoWithId)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewProduct provides a mock function with given fields: ctx, content
func (_m *Product) NewProduct(ctx context.Context, content *proto.ProductInfo) error {
	ret := _m.Called(ctx, content)

	if len(ret) == 0 {
		panic("no return value specified for NewProduct")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.ProductInfo) error); ok {
		r0 = rf(ctx, content)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SellProduct provides a mock function with given fields: ctx, content
func (_m *Product) SellProduct(ctx context.Context, content *proto.SellProductRequest) error {
	ret := _m.Called(ctx, content)

	if len(ret) == 0 {
		panic("no return value specified for SellProduct")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.SellProductRequest) error); ok {
		r0 = rf(ctx, content)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewProduct creates a new instance of Product. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProduct(t interface {
	mock.TestingT
	Cleanup(func())
}) *Product {
	mock := &Product{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
