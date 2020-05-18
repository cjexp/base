// +build debug

// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

package cache

import (
	ctx "github.com/cjtoolkit/ctx"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockCore is a mock of Core interface
type MockCore struct {
	ctrl     *gomock.Controller
	recorder *MockCoreMockRecorder
}

// MockCoreMockRecorder is the mock recorder for MockCore
type MockCoreMockRecorder struct {
	mock *MockCore
}

// NewMockCore creates a new mock instance
func NewMockCore(ctrl *gomock.Controller) *MockCore {
	mock := &MockCore{ctrl: ctrl}
	mock.recorder = &MockCoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCore) EXPECT() *MockCoreMockRecorder {
	return m.recorder
}

// GetBytes mocks base method
func (m *MockCore) GetBytes(key string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBytes", key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBytes indicates an expected call of GetBytes
func (mr *MockCoreMockRecorder) GetBytes(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBytes", reflect.TypeOf((*MockCore)(nil).GetBytes), key)
}

// MustGetBytes mocks base method
func (m *MockCore) MustGetBytes(key string) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MustGetBytes", key)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// MustGetBytes indicates an expected call of MustGetBytes
func (mr *MockCoreMockRecorder) MustGetBytes(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MustGetBytes", reflect.TypeOf((*MockCore)(nil).MustGetBytes), key)
}

// SetBytes mocks base method
func (m *MockCore) SetBytes(key string, value []byte, expiration time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetBytes", key, value, expiration)
}

// SetBytes indicates an expected call of SetBytes
func (mr *MockCoreMockRecorder) SetBytes(key, value, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBytes", reflect.TypeOf((*MockCore)(nil).SetBytes), key, value, expiration)
}

// Exist mocks base method
func (m *MockCore) Exist(key string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exist", key)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exist indicates an expected call of Exist
func (mr *MockCoreMockRecorder) Exist(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exist", reflect.TypeOf((*MockCore)(nil).Exist), key)
}

// Delete mocks base method
func (m *MockCore) Delete(keys ...string) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Delete", varargs...)
}

// Delete indicates an expected call of Delete
func (mr *MockCoreMockRecorder) Delete(keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCore)(nil).Delete), keys...)
}

// MockCoreGetCheck is a mock of CoreGetCheck interface
type MockCoreGetCheck struct {
	ctrl     *gomock.Controller
	recorder *MockCoreGetCheckMockRecorder
}

// MockCoreGetCheckMockRecorder is the mock recorder for MockCoreGetCheck
type MockCoreGetCheckMockRecorder struct {
	mock *MockCoreGetCheck
}

// NewMockCoreGetCheck creates a new mock instance
func NewMockCoreGetCheck(ctrl *gomock.Controller) *MockCoreGetCheck {
	mock := &MockCoreGetCheck{ctrl: ctrl}
	mock.recorder = &MockCoreGetCheckMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCoreGetCheck) EXPECT() *MockCoreGetCheckMockRecorder {
	return m.recorder
}

// GetBytes mocks base method
func (m *MockCoreGetCheck) GetBytes(key string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBytes", key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBytes indicates an expected call of GetBytes
func (mr *MockCoreGetCheckMockRecorder) GetBytes(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBytes", reflect.TypeOf((*MockCoreGetCheck)(nil).GetBytes), key)
}

// MustGetBytes mocks base method
func (m *MockCoreGetCheck) MustGetBytes(key string) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MustGetBytes", key)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// MustGetBytes indicates an expected call of MustGetBytes
func (mr *MockCoreGetCheckMockRecorder) MustGetBytes(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MustGetBytes", reflect.TypeOf((*MockCoreGetCheck)(nil).MustGetBytes), key)
}

// SetBytes mocks base method
func (m *MockCoreGetCheck) SetBytes(key string, value []byte, expiration time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetBytes", key, value, expiration)
}

// SetBytes indicates an expected call of SetBytes
func (mr *MockCoreGetCheckMockRecorder) SetBytes(key, value, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBytes", reflect.TypeOf((*MockCoreGetCheck)(nil).SetBytes), key, value, expiration)
}

// Exist mocks base method
func (m *MockCoreGetCheck) Exist(key string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exist", key)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exist indicates an expected call of Exist
func (mr *MockCoreGetCheckMockRecorder) Exist(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exist", reflect.TypeOf((*MockCoreGetCheck)(nil).Exist), key)
}

// Delete mocks base method
func (m *MockCoreGetCheck) Delete(keys ...string) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Delete", varargs...)
}

// Delete indicates an expected call of Delete
func (mr *MockCoreGetCheckMockRecorder) Delete(keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCoreGetCheck)(nil).Delete), keys...)
}

// GetBytesCheck mocks base method
func (m *MockCoreGetCheck) GetBytesCheck(key string, expiration time.Duration) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBytesCheck", key, expiration)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBytesCheck indicates an expected call of GetBytesCheck
func (mr *MockCoreGetCheckMockRecorder) GetBytesCheck(key, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBytesCheck", reflect.TypeOf((*MockCoreGetCheck)(nil).GetBytesCheck), key, expiration)
}

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Persist mocks base method
func (m *MockRepository) Persist(name string, expiration time.Duration, miss Miss, hit Hit) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Persist", name, expiration, miss, hit)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Persist indicates an expected call of Persist
func (mr *MockRepositoryMockRecorder) Persist(name, expiration, miss, hit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Persist", reflect.TypeOf((*MockRepository)(nil).Persist), name, expiration, miss, hit)
}

// MockModifiedRepository is a mock of ModifiedRepository interface
type MockModifiedRepository struct {
	ctrl     *gomock.Controller
	recorder *MockModifiedRepositoryMockRecorder
}

// MockModifiedRepositoryMockRecorder is the mock recorder for MockModifiedRepository
type MockModifiedRepositoryMockRecorder struct {
	mock *MockModifiedRepository
}

// NewMockModifiedRepository creates a new mock instance
func NewMockModifiedRepository(ctrl *gomock.Controller) *MockModifiedRepository {
	mock := &MockModifiedRepository{ctrl: ctrl}
	mock.recorder = &MockModifiedRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockModifiedRepository) EXPECT() *MockModifiedRepositoryMockRecorder {
	return m.recorder
}

// Persist mocks base method
func (m *MockModifiedRepository) Persist(context ctx.Context, name string, expiration time.Duration, miss Miss, hit Hit) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Persist", context, name, expiration, miss, hit)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Persist indicates an expected call of Persist
func (mr *MockModifiedRepositoryMockRecorder) Persist(context, name, expiration, miss, hit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Persist", reflect.TypeOf((*MockModifiedRepository)(nil).Persist), context, name, expiration, miss, hit)
}