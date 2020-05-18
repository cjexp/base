// +build debug

// Code generated by MockGen. DO NOT EDIT.
// Source: cookie.go

package cookie

import (
	ctx "github.com/cjtoolkit/ctx"
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockCookieHelper is a mock of CookieHelper interface
type MockCookieHelper struct {
	ctrl     *gomock.Controller
	recorder *MockCookieHelperMockRecorder
}

// MockCookieHelperMockRecorder is the mock recorder for MockCookieHelper
type MockCookieHelperMockRecorder struct {
	mock *MockCookieHelper
}

// NewMockCookieHelper creates a new mock instance
func NewMockCookieHelper(ctrl *gomock.Controller) *MockCookieHelper {
	mock := &MockCookieHelper{ctrl: ctrl}
	mock.recorder = &MockCookieHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCookieHelper) EXPECT() *MockCookieHelperMockRecorder {
	return m.recorder
}

// Set mocks base method
func (m *MockCookieHelper) Set(context ctx.Context, cookie *http.Cookie) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", context, cookie)
}

// Set indicates an expected call of Set
func (mr *MockCookieHelperMockRecorder) Set(context, cookie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCookieHelper)(nil).Set), context, cookie)
}

// Get mocks base method
func (m *MockCookieHelper) Get(context ctx.Context, name string) *http.Cookie {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", context, name)
	ret0, _ := ret[0].(*http.Cookie)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockCookieHelperMockRecorder) Get(context, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCookieHelper)(nil).Get), context, name)
}

// Delete mocks base method
func (m *MockCookieHelper) Delete(context ctx.Context, name string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", context, name)
}

// Delete indicates an expected call of Delete
func (mr *MockCookieHelperMockRecorder) Delete(context, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCookieHelper)(nil).Delete), context, name)
}

// GetValue mocks base method
func (m *MockCookieHelper) GetValue(context ctx.Context, name string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValue", context, name)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetValue indicates an expected call of GetValue
func (mr *MockCookieHelperMockRecorder) GetValue(context, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValue", reflect.TypeOf((*MockCookieHelper)(nil).GetValue), context, name)
}