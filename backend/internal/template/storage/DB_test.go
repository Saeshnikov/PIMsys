// Code generated by mockery v2.52.4. DO NOT EDIT.

package template_storage_test

import (
	sql "database/sql"

	mock "github.com/stretchr/testify/mock"
)

// DB is an autogenerated mock type for the DB type
type DB struct {
	mock.Mock
}

// Close provides a mock function with no fields
func (_m *DB) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Prepare provides a mock function with given fields: query
func (_m *DB) Prepare(query string) (*sql.Stmt, error) {
	ret := _m.Called(query)

	if len(ret) == 0 {
		panic("no return value specified for Prepare")
	}

	var r0 *sql.Stmt
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*sql.Stmt, error)); ok {
		return rf(query)
	}
	if rf, ok := ret.Get(0).(func(string) *sql.Stmt); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Stmt)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDB creates a new instance of DB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *DB {
	mock := &DB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
