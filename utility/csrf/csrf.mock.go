// +build debug

// Code generated by MockGen. DO NOT EDIT.
// Source: csrf.go

package csrf

import (
	reflect "reflect"

	ctx "github.com/cjtoolkit/ctx/v2"
	gomock "github.com/golang/mock/gomock"
)

// MockCsrfController is a mock of Controller interface
type MockCsrfController struct {
	ctrl     *gomock.Controller
	recorder *MockCsrfControllerMockRecorder
}

// MockCsrfControllerMockRecorder is the mock recorder for MockCsrfController
type MockCsrfControllerMockRecorder struct {
	mock *MockCsrfController
}

// NewMockCsrfController creates a new mock instance
func NewMockCsrfController(ctrl *gomock.Controller) *MockCsrfController {
	mock := &MockCsrfController{ctrl: ctrl}
	mock.recorder = &MockCsrfControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCsrfController) EXPECT() *MockCsrfControllerMockRecorder {
	return m.recorder
}

// GetCsrfData mocks base method
func (m *MockCsrfController) GetCsrfData(context ctx.Context) Data {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCsrfData", context)
	ret0, _ := ret[0].(Data)
	return ret0
}

// GetCsrfData indicates an expected call of GetCsrfData
func (mr *MockCsrfControllerMockRecorder) GetCsrfData(context interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCsrfData", reflect.TypeOf((*MockCsrfController)(nil).GetCsrfData), context)
}

// InitCsrf mocks base method
func (m *MockCsrfController) InitCsrf(context ctx.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InitCsrf", context)
}

// InitCsrf indicates an expected call of InitCsrf
func (mr *MockCsrfControllerMockRecorder) InitCsrf(context interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitCsrf", reflect.TypeOf((*MockCsrfController)(nil).InitCsrf), context)
}

// CheckCsrf mocks base method
func (m *MockCsrfController) CheckCsrf(context ctx.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CheckCsrf", context)
}

// CheckCsrf indicates an expected call of CheckCsrf
func (mr *MockCsrfControllerMockRecorder) CheckCsrf(context interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckCsrf", reflect.TypeOf((*MockCsrfController)(nil).CheckCsrf), context)
}
