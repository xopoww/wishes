// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/xopoww/wishes/internal/service (interfaces: ListTokenProvider)

// Package service_test is a generated GoMock package.
package service_test

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	service "github.com/xopoww/wishes/internal/service"
)

// MockListTokenProvider is a mock of ListTokenProvider interface.
type MockListTokenProvider struct {
	ctrl     *gomock.Controller
	recorder *MockListTokenProviderMockRecorder
}

// MockListTokenProviderMockRecorder is the mock recorder for MockListTokenProvider.
type MockListTokenProviderMockRecorder struct {
	mock *MockListTokenProvider
}

// NewMockListTokenProvider creates a new mock instance.
func NewMockListTokenProvider(ctrl *gomock.Controller) *MockListTokenProvider {
	mock := &MockListTokenProvider{ctrl: ctrl}
	mock.recorder = &MockListTokenProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockListTokenProvider) EXPECT() *MockListTokenProviderMockRecorder {
	return m.recorder
}

// GenerateToken mocks base method.
func (m *MockListTokenProvider) GenerateToken(arg0 service.ListClaims) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockListTokenProviderMockRecorder) GenerateToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockListTokenProvider)(nil).GenerateToken), arg0)
}

// ValidateToken mocks base method.
func (m *MockListTokenProvider) ValidateToken(arg0 string) (service.ListClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateToken", arg0)
	ret0, _ := ret[0].(service.ListClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateToken indicates an expected call of ValidateToken.
func (mr *MockListTokenProviderMockRecorder) ValidateToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateToken", reflect.TypeOf((*MockListTokenProvider)(nil).ValidateToken), arg0)
}