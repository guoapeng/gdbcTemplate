// Code generated by MockGen. DO NOT EDIT.
// Source: mapper/rowsmapper.go

// Package mocks is a generated GoMock package.
package mocks

import (
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mapper "github.com/guoapeng/gdbcTemplate/mapper"
)

// MockBeanPropertyRowMapper is a mock of BeanPropertyRowMapper interface.
type MockBeanPropertyRowMapper struct {
	ctrl     *gomock.Controller
	recorder *MockBeanPropertyRowMapperMockRecorder
}

// MockBeanPropertyRowMapperMockRecorder is the mock recorder for MockBeanPropertyRowMapper.
type MockBeanPropertyRowMapperMockRecorder struct {
	mock *MockBeanPropertyRowMapper
}

// NewMockBeanPropertyRowMapper creates a new mock instance.
func NewMockBeanPropertyRowMapper(ctrl *gomock.Controller) *MockBeanPropertyRowMapper {
	mock := &MockBeanPropertyRowMapper{ctrl: ctrl}
	mock.recorder = &MockBeanPropertyRowMapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBeanPropertyRowMapper) EXPECT() *MockBeanPropertyRowMapperMockRecorder {
	return m.recorder
}

// RowMapper mocks base method.
func (m *MockBeanPropertyRowMapper) RowMapper(row *sql.Row) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RowMapper", row)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// RowMapper indicates an expected call of RowMapper.
func (mr *MockBeanPropertyRowMapperMockRecorder) RowMapper(row interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RowMapper", reflect.TypeOf((*MockBeanPropertyRowMapper)(nil).RowMapper), row)
}

// RowsMapper mocks base method.
func (m *MockBeanPropertyRowMapper) RowsMapper(rows *sql.Rows) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RowsMapper", rows)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// RowsMapper indicates an expected call of RowsMapper.
func (mr *MockBeanPropertyRowMapperMockRecorder) RowsMapper(rows interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RowsMapper", reflect.TypeOf((*MockBeanPropertyRowMapper)(nil).RowsMapper), rows)
}

// MockRowConvertor is a mock of RowConvertor interface.
type MockRowConvertor struct {
	ctrl     *gomock.Controller
	recorder *MockRowConvertorMockRecorder
}

// MockRowConvertorMockRecorder is the mock recorder for MockRowConvertor.
type MockRowConvertorMockRecorder struct {
	mock *MockRowConvertor
}

// NewMockRowConvertor creates a new mock instance.
func NewMockRowConvertor(ctrl *gomock.Controller) *MockRowConvertor {
	mock := &MockRowConvertor{ctrl: ctrl}
	mock.recorder = &MockRowConvertorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRowConvertor) EXPECT() *MockRowConvertorMockRecorder {
	return m.recorder
}

// Map mocks base method.
func (m *MockRowConvertor) Map(rowmapper mapper.RowMapper) mapper.RowConvertor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map", rowmapper)
	ret0, _ := ret[0].(mapper.RowConvertor)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockRowConvertorMockRecorder) Map(mapper interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockRowConvertor)(nil).Map), mapper)
}

// MapTo mocks base method.
func (m *MockRowConvertor) MapTo(example interface{}) mapper.RowConvertor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MapTo", example)
	ret0, _ := ret[0].(mapper.RowConvertor)
	return ret0
}

// MapTo indicates an expected call of MapTo.
func (mr *MockRowConvertorMockRecorder) MapTo(example interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MapTo", reflect.TypeOf((*MockRowConvertor)(nil).MapTo), example)
}

// ToObject mocks base method.
func (m *MockRowConvertor) ToObject() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToObject")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// ToObject indicates an expected call of ToObject.
func (mr *MockRowConvertorMockRecorder) ToObject() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToObject", reflect.TypeOf((*MockRowConvertor)(nil).ToObject))
}

// MockRowsConvertor is a mock of RowsConvertor interface.
type MockRowsConvertor struct {
	ctrl     *gomock.Controller
	recorder *MockRowsConvertorMockRecorder
}

// MockRowsConvertorMockRecorder is the mock recorder for MockRowsConvertor.
type MockRowsConvertorMockRecorder struct {
	mock *MockRowsConvertor
}

// NewMockRowsConvertor creates a new mock instance.
func NewMockRowsConvertor(ctrl *gomock.Controller) *MockRowsConvertor {
	mock := &MockRowsConvertor{ctrl: ctrl}
	mock.recorder = &MockRowsConvertorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRowsConvertor) EXPECT() *MockRowsConvertorMockRecorder {
	return m.recorder
}

// Map mocks base method.
func (m *MockRowsConvertor) Map(rowmapper mapper.RowsMapper) mapper.RowsConvertor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map", rowmapper)
	ret0, _ := ret[0].(mapper.RowsConvertor)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockRowsConvertorMockRecorder) Map(mapper interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockRowsConvertor)(nil).Map), mapper)
}

// MapTo mocks base method.
func (m *MockRowsConvertor) MapTo(example interface{}) mapper.RowsConvertor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MapTo", example)
	ret0, _ := ret[0].(mapper.RowsConvertor)
	return ret0
}

// MapTo indicates an expected call of MapTo.
func (mr *MockRowsConvertorMockRecorder) MapTo(example interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MapTo", reflect.TypeOf((*MockRowsConvertor)(nil).MapTo), example)
}

// ToArray mocks base method.
func (m *MockRowsConvertor) ToArray() []interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToArray")
	ret0, _ := ret[0].([]interface{})
	return ret0
}

// ToArray indicates an expected call of ToArray.
func (mr *MockRowsConvertorMockRecorder) ToArray() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToArray", reflect.TypeOf((*MockRowsConvertor)(nil).ToArray))
}
