// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/localConfigProvider/localConfigProvider.go

// Package localConfigProvider is a generated GoMock package.
package localConfigProvider

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLocalConfigProvider is a mock of LocalConfigProvider interface.
type MockLocalConfigProvider struct {
	ctrl     *gomock.Controller
	recorder *MockLocalConfigProviderMockRecorder
}

// MockLocalConfigProviderMockRecorder is the mock recorder for MockLocalConfigProvider.
type MockLocalConfigProviderMockRecorder struct {
	mock *MockLocalConfigProvider
}

// NewMockLocalConfigProvider creates a new mock instance.
func NewMockLocalConfigProvider(ctrl *gomock.Controller) *MockLocalConfigProvider {
	mock := &MockLocalConfigProvider{ctrl: ctrl}
	mock.recorder = &MockLocalConfigProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocalConfigProvider) EXPECT() *MockLocalConfigProviderMockRecorder {
	return m.recorder
}

// Exists mocks base method.
func (m *MockLocalConfigProvider) Exists() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exists indicates an expected call of Exists.
func (mr *MockLocalConfigProviderMockRecorder) Exists() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockLocalConfigProvider)(nil).Exists))
}

// GetApplication mocks base method.
func (m *MockLocalConfigProvider) GetApplication() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplication")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetApplication indicates an expected call of GetApplication.
func (mr *MockLocalConfigProviderMockRecorder) GetApplication() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplication", reflect.TypeOf((*MockLocalConfigProvider)(nil).GetApplication))
}

// GetContainers mocks base method.
func (m *MockLocalConfigProvider) GetContainers() ([]LocalContainer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainers")
	ret0, _ := ret[0].([]LocalContainer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContainers indicates an expected call of GetContainers.
func (mr *MockLocalConfigProviderMockRecorder) GetContainers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainers", reflect.TypeOf((*MockLocalConfigProvider)(nil).GetContainers))
}

// GetDebugPort mocks base method.
func (m *MockLocalConfigProvider) GetDebugPort() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDebugPort")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetDebugPort indicates an expected call of GetDebugPort.
func (mr *MockLocalConfigProviderMockRecorder) GetDebugPort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDebugPort", reflect.TypeOf((*MockLocalConfigProvider)(nil).GetDebugPort))
}

// GetName mocks base method.
func (m *MockLocalConfigProvider) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockLocalConfigProviderMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockLocalConfigProvider)(nil).GetName))
}

// GetNamespace mocks base method.
func (m *MockLocalConfigProvider) GetNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetNamespace indicates an expected call of GetNamespace.
func (mr *MockLocalConfigProviderMockRecorder) GetNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockLocalConfigProvider)(nil).GetNamespace))
}

// ListStorage mocks base method.
func (m *MockLocalConfigProvider) ListStorage() ([]LocalStorage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStorage")
	ret0, _ := ret[0].([]LocalStorage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStorage indicates an expected call of ListStorage.
func (mr *MockLocalConfigProviderMockRecorder) ListStorage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStorage", reflect.TypeOf((*MockLocalConfigProvider)(nil).ListStorage))
}
