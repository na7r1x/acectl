// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/na7r1x/acectl/internal/core/ports (interfaces: BrokerService,ExecutorService)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	domain "github.com/na7r1x/acectl/internal/core/domain"
	reflect "reflect"
)

// MockBrokerService is a mock of BrokerService interface
type MockBrokerService struct {
	ctrl     *gomock.Controller
	recorder *MockBrokerServiceMockRecorder
}

// MockBrokerServiceMockRecorder is the mock recorder for MockBrokerService
type MockBrokerServiceMockRecorder struct {
	mock *MockBrokerService
}

// NewMockBrokerService creates a new mock instance
func NewMockBrokerService(ctrl *gomock.Controller) *MockBrokerService {
	mock := &MockBrokerService{ctrl: ctrl}
	mock.recorder = &MockBrokerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBrokerService) EXPECT() *MockBrokerServiceMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockBrokerService) Get(arg0 string) (domain.Broker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(domain.Broker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockBrokerServiceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBrokerService)(nil).Get), arg0)
}

// List mocks base method
func (m *MockBrokerService) List() ([]domain.Broker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]domain.Broker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockBrokerServiceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockBrokerService)(nil).List))
}

// Register mocks base method
func (m *MockBrokerService) Register(arg0 domain.Broker) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockBrokerServiceMockRecorder) Register(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockBrokerService)(nil).Register), arg0)
}

// Start mocks base method
func (m *MockBrokerService) Start(arg0 string, arg1 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Start indicates an expected call of Start
func (mr *MockBrokerServiceMockRecorder) Start(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockBrokerService)(nil).Start), arg0, arg1)
}

// Status mocks base method
func (m *MockBrokerService) Status(arg0 string, arg1 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status
func (mr *MockBrokerServiceMockRecorder) Status(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockBrokerService)(nil).Status), arg0, arg1)
}

// Stop mocks base method
func (m *MockBrokerService) Stop(arg0 string, arg1 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stop indicates an expected call of Stop
func (mr *MockBrokerServiceMockRecorder) Stop(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockBrokerService)(nil).Stop), arg0, arg1)
}

// Unregister mocks base method
func (m *MockBrokerService) Unregister(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unregister", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unregister indicates an expected call of Unregister
func (mr *MockBrokerServiceMockRecorder) Unregister(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unregister", reflect.TypeOf((*MockBrokerService)(nil).Unregister), arg0)
}

// MockExecutorService is a mock of ExecutorService interface
type MockExecutorService struct {
	ctrl     *gomock.Controller
	recorder *MockExecutorServiceMockRecorder
}

// MockExecutorServiceMockRecorder is the mock recorder for MockExecutorService
type MockExecutorServiceMockRecorder struct {
	mock *MockExecutorService
}

// NewMockExecutorService creates a new mock instance
func NewMockExecutorService(ctrl *gomock.Controller) *MockExecutorService {
	mock := &MockExecutorService{ctrl: ctrl}
	mock.recorder = &MockExecutorServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExecutorService) EXPECT() *MockExecutorServiceMockRecorder {
	return m.recorder
}

// Exec mocks base method
func (m *MockExecutorService) Exec(arg0 domain.Broker, arg1 []string, arg2 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exec", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec
func (mr *MockExecutorServiceMockRecorder) Exec(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockExecutorService)(nil).Exec), arg0, arg1, arg2)
}
