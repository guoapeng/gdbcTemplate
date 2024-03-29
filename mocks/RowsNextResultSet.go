// Code generated by mockery v2.25.0. DO NOT EDIT.

package mocks

import (
	driver "database/sql/driver"

	mock "github.com/stretchr/testify/mock"
)

// RowsNextResultSet is an autogenerated mock type for the RowsNextResultSet type
type RowsNextResultSet struct {
	mock.Mock
}

type RowsNextResultSet_Expecter struct {
	mock *mock.Mock
}

func (_m *RowsNextResultSet) EXPECT() *RowsNextResultSet_Expecter {
	return &RowsNextResultSet_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *RowsNextResultSet) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RowsNextResultSet_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type RowsNextResultSet_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *RowsNextResultSet_Expecter) Close() *RowsNextResultSet_Close_Call {
	return &RowsNextResultSet_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *RowsNextResultSet_Close_Call) Run(run func()) *RowsNextResultSet_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RowsNextResultSet_Close_Call) Return(_a0 error) *RowsNextResultSet_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RowsNextResultSet_Close_Call) RunAndReturn(run func() error) *RowsNextResultSet_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Columns provides a mock function with given fields:
func (_m *RowsNextResultSet) Columns() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// RowsNextResultSet_Columns_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Columns'
type RowsNextResultSet_Columns_Call struct {
	*mock.Call
}

// Columns is a helper method to define mock.On call
func (_e *RowsNextResultSet_Expecter) Columns() *RowsNextResultSet_Columns_Call {
	return &RowsNextResultSet_Columns_Call{Call: _e.mock.On("Columns")}
}

func (_c *RowsNextResultSet_Columns_Call) Run(run func()) *RowsNextResultSet_Columns_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RowsNextResultSet_Columns_Call) Return(_a0 []string) *RowsNextResultSet_Columns_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RowsNextResultSet_Columns_Call) RunAndReturn(run func() []string) *RowsNextResultSet_Columns_Call {
	_c.Call.Return(run)
	return _c
}

// HasNextResultSet provides a mock function with given fields:
func (_m *RowsNextResultSet) HasNextResultSet() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RowsNextResultSet_HasNextResultSet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HasNextResultSet'
type RowsNextResultSet_HasNextResultSet_Call struct {
	*mock.Call
}

// HasNextResultSet is a helper method to define mock.On call
func (_e *RowsNextResultSet_Expecter) HasNextResultSet() *RowsNextResultSet_HasNextResultSet_Call {
	return &RowsNextResultSet_HasNextResultSet_Call{Call: _e.mock.On("HasNextResultSet")}
}

func (_c *RowsNextResultSet_HasNextResultSet_Call) Run(run func()) *RowsNextResultSet_HasNextResultSet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RowsNextResultSet_HasNextResultSet_Call) Return(_a0 bool) *RowsNextResultSet_HasNextResultSet_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RowsNextResultSet_HasNextResultSet_Call) RunAndReturn(run func() bool) *RowsNextResultSet_HasNextResultSet_Call {
	_c.Call.Return(run)
	return _c
}

// Next provides a mock function with given fields: dest
func (_m *RowsNextResultSet) Next(dest []driver.Value) error {
	ret := _m.Called(dest)

	var r0 error
	if rf, ok := ret.Get(0).(func([]driver.Value) error); ok {
		r0 = rf(dest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RowsNextResultSet_Next_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Next'
type RowsNextResultSet_Next_Call struct {
	*mock.Call
}

// Next is a helper method to define mock.On call
//   - dest []driver.Value
func (_e *RowsNextResultSet_Expecter) Next(dest interface{}) *RowsNextResultSet_Next_Call {
	return &RowsNextResultSet_Next_Call{Call: _e.mock.On("Next", dest)}
}

func (_c *RowsNextResultSet_Next_Call) Run(run func(dest []driver.Value)) *RowsNextResultSet_Next_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]driver.Value))
	})
	return _c
}

func (_c *RowsNextResultSet_Next_Call) Return(_a0 error) *RowsNextResultSet_Next_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RowsNextResultSet_Next_Call) RunAndReturn(run func([]driver.Value) error) *RowsNextResultSet_Next_Call {
	_c.Call.Return(run)
	return _c
}

// NextResultSet provides a mock function with given fields:
func (_m *RowsNextResultSet) NextResultSet() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RowsNextResultSet_NextResultSet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NextResultSet'
type RowsNextResultSet_NextResultSet_Call struct {
	*mock.Call
}

// NextResultSet is a helper method to define mock.On call
func (_e *RowsNextResultSet_Expecter) NextResultSet() *RowsNextResultSet_NextResultSet_Call {
	return &RowsNextResultSet_NextResultSet_Call{Call: _e.mock.On("NextResultSet")}
}

func (_c *RowsNextResultSet_NextResultSet_Call) Run(run func()) *RowsNextResultSet_NextResultSet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RowsNextResultSet_NextResultSet_Call) Return(_a0 error) *RowsNextResultSet_NextResultSet_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RowsNextResultSet_NextResultSet_Call) RunAndReturn(run func() error) *RowsNextResultSet_NextResultSet_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewRowsNextResultSet interface {
	mock.TestingT
	Cleanup(func())
}

// NewRowsNextResultSet creates a new instance of RowsNextResultSet. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRowsNextResultSet(t mockConstructorTestingTNewRowsNextResultSet) *RowsNextResultSet {
	mock := &RowsNextResultSet{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
