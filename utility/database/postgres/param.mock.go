// +build debug

// Code generated by MockGen. DO NOT EDIT.
// Source: param.go

package postgres

import (
	sql "database/sql"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockParamTemplate is a mock of ParamTemplate interface
type MockParamTemplate struct {
	ctrl     *gomock.Controller
	recorder *MockParamTemplateMockRecorder
}

// MockParamTemplateMockRecorder is the mock recorder for MockParamTemplate
type MockParamTemplateMockRecorder struct {
	mock *MockParamTemplate
}

// NewMockParamTemplate creates a new mock instance
func NewMockParamTemplate(ctrl *gomock.Controller) *MockParamTemplate {
	mock := &MockParamTemplate{ctrl: ctrl}
	mock.recorder = &MockParamTemplateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockParamTemplate) EXPECT() *MockParamTemplateMockRecorder {
	return m.recorder
}

// ToSlice mocks base method
func (m *MockParamTemplate) ToSlice(param Param) []interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToSlice", param)
	ret0, _ := ret[0].([]interface{})
	return ret0
}

// ToSlice indicates an expected call of ToSlice
func (mr *MockParamTemplateMockRecorder) ToSlice(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToSlice", reflect.TypeOf((*MockParamTemplate)(nil).ToSlice), param)
}

// MockParamBuilder is a mock of ParamBuilder interface
type MockParamBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockParamBuilderMockRecorder
}

// MockParamBuilderMockRecorder is the mock recorder for MockParamBuilder
type MockParamBuilderMockRecorder struct {
	mock *MockParamBuilder
}

// NewMockParamBuilder creates a new mock instance
func NewMockParamBuilder(ctrl *gomock.Controller) *MockParamBuilder {
	mock := &MockParamBuilder{ctrl: ctrl}
	mock.recorder = &MockParamBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockParamBuilder) EXPECT() *MockParamBuilderMockRecorder {
	return m.recorder
}

// BuildParamTemplate mocks base method
func (m *MockParamBuilder) BuildParamTemplate(query string) (string, ParamTemplate) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildParamTemplate", query)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(ParamTemplate)
	return ret0, ret1
}

// BuildParamTemplate indicates an expected call of BuildParamTemplate
func (mr *MockParamBuilderMockRecorder) BuildParamTemplate(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildParamTemplate", reflect.TypeOf((*MockParamBuilder)(nil).BuildParamTemplate), query)
}

// BuildParamTemplateAndPrepare mocks base method
func (m *MockParamBuilder) BuildParamTemplateAndPrepare(dbConn *sql.DB, query string) (*sql.Stmt, ParamTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildParamTemplateAndPrepare", dbConn, query)
	ret0, _ := ret[0].(*sql.Stmt)
	ret1, _ := ret[1].(ParamTemplate)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// BuildParamTemplateAndPrepare indicates an expected call of BuildParamTemplateAndPrepare
func (mr *MockParamBuilderMockRecorder) BuildParamTemplateAndPrepare(dbConn, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildParamTemplateAndPrepare", reflect.TypeOf((*MockParamBuilder)(nil).BuildParamTemplateAndPrepare), dbConn, query)
}
