// Code generated by MockGen. DO NOT EDIT.
// Source: ./checker/checker.go

// Package checker is a generated GoMock package.
package checker

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockChecker is a mock of Checker interface
type MockChecker struct {
	ctrl     *gomock.Controller
	recorder *MockCheckerMockRecorder
}

// MockCheckerMockRecorder is the mock recorder for MockChecker
type MockCheckerMockRecorder struct {
	mock *MockChecker
}

// NewMockChecker creates a new mock instance
func NewMockChecker(ctrl *gomock.Controller) *MockChecker {
	mock := &MockChecker{ctrl: ctrl}
	mock.recorder = &MockCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChecker) EXPECT() *MockCheckerMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockChecker) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start
func (mr *MockCheckerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockChecker)(nil).Start))
}

// Get mocks base method
func (m *MockChecker) Get(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockCheckerMockRecorder) Get(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockChecker)(nil).Get), name)
}

// GetAll mocks base method
func (m *MockChecker) GetAll() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAll indicates an expected call of GetAll
func (mr *MockCheckerMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockChecker)(nil).GetAll))
}
