// Code generated by MockGen. DO NOT EDIT.
// Source: app/services/auth.go

package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	schema "github.com/quangdangfit/go-admin/app/schema"
	reflect "reflect"
)

// MockIAuthService is a mock of IAuthService interface
type MockIAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockIAuthServiceMockRecorder
}

// MockIAuthServiceMockRecorder is the mock recorder for MockIAuthService
type MockIAuthServiceMockRecorder struct {
	mock *MockIAuthService
}

// NewMockIAuthService creates a new mock instance
func NewMockIAuthService(ctrl *gomock.Controller) *MockIAuthService {
	mock := &MockIAuthService{ctrl: ctrl}
	mock.recorder = &MockIAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockIAuthService) EXPECT() *MockIAuthServiceMockRecorder {
	return _m.recorder
}

// Login mocks base method
func (_m *MockIAuthService) Login(ctx context.Context, bodyParam *schema.LoginBodyParam) (*schema.UserTokenInfo, error) {
	ret := _m.ctrl.Call(_m, "Login", ctx, bodyParam)
	ret0, _ := ret[0].(*schema.UserTokenInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login
func (_mr *MockIAuthServiceMockRecorder) Login(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Login", reflect.TypeOf((*MockIAuthService)(nil).Login), arg0, arg1)
}

// Register mocks base method
func (_m *MockIAuthService) Register(ctx context.Context, bodyParam *schema.RegisterBodyParam) (*schema.UserTokenInfo, error) {
	ret := _m.ctrl.Call(_m, "Register", ctx, bodyParam)
	ret0, _ := ret[0].(*schema.UserTokenInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register
func (_mr *MockIAuthServiceMockRecorder) Register(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Register", reflect.TypeOf((*MockIAuthService)(nil).Register), arg0, arg1)
}

// Refresh mocks base method
func (_m *MockIAuthService) Refresh(ctx context.Context, bodyParam *schema.RefreshBodyParam) (*schema.UserTokenInfo, error) {
	ret := _m.ctrl.Call(_m, "Refresh", ctx, bodyParam)
	ret0, _ := ret[0].(*schema.UserTokenInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Refresh indicates an expected call of Refresh
func (_mr *MockIAuthServiceMockRecorder) Refresh(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Refresh", reflect.TypeOf((*MockIAuthService)(nil).Refresh), arg0, arg1)
}

// Logout mocks base method
func (_m *MockIAuthService) Logout(ctx context.Context) error {
	ret := _m.ctrl.Call(_m, "Logout", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout
func (_mr *MockIAuthServiceMockRecorder) Logout(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Logout", reflect.TypeOf((*MockIAuthService)(nil).Logout), arg0)
}