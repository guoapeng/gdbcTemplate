// Code generated by mockery v2.25.0. DO NOT EDIT.

package mocks

import (
	sql "database/sql"

	mock "github.com/stretchr/testify/mock"
)

// OpenDb is an autogenerated mock type for the OpenDb type
type OpenDb struct {
	mock.Mock
}

type OpenDb_Expecter struct {
	mock *mock.Mock
}

func (_m *OpenDb) EXPECT() *OpenDb_Expecter {
	return &OpenDb_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: driverName, dataSourceName
func (_m *OpenDb) Execute(driverName string, dataSourceName string) (*sql.DB, error) {
	ret := _m.Called(driverName, dataSourceName)

	var r0 *sql.DB
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*sql.DB, error)); ok {
		return rf(driverName, dataSourceName)
	}
	if rf, ok := ret.Get(0).(func(string, string) *sql.DB); ok {
		r0 = rf(driverName, dataSourceName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.DB)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(driverName, dataSourceName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OpenDb_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type OpenDb_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - driverName string
//   - dataSourceName string
func (_e *OpenDb_Expecter) Execute(driverName interface{}, dataSourceName interface{}) *OpenDb_Execute_Call {
	return &OpenDb_Execute_Call{Call: _e.mock.On("Execute", driverName, dataSourceName)}
}

func (_c *OpenDb_Execute_Call) Run(run func(driverName string, dataSourceName string)) *OpenDb_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *OpenDb_Execute_Call) Return(_a0 *sql.DB, _a1 error) *OpenDb_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OpenDb_Execute_Call) RunAndReturn(run func(string, string) (*sql.DB, error)) *OpenDb_Execute_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewOpenDb interface {
	mock.TestingT
	Cleanup(func())
}

// NewOpenDb creates a new instance of OpenDb. It also registers a testing interface 
// on the mock and a cleanup function to assert the mocks expectations.
func NewOpenDb(t mockConstructorTestingTNewOpenDb) *OpenDb {
	mock := &OpenDb{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
