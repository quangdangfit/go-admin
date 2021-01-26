// Code generated by MockGen. DO NOT EDIT.
// Source: app/repositories/user.go

package mocks

import (
	gomock "github.com/golang/mock/gomock"
	models "github.com/quangdangfit/go-admin/app/models"
	schema "github.com/quangdangfit/go-admin/app/schema"
	reflect "reflect"
)

// MockIUserRepository is a mock of IUserRepository interface
type MockIUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepositoryMockRecorder
}

// MockIUserRepositoryMockRecorder is the mock recorder for MockIUserRepository
type MockIUserRepositoryMockRecorder struct {
	mock *MockIUserRepository
}

// NewMockIUserRepository creates a new mock instance
func NewMockIUserRepository(ctrl *gomock.Controller) *MockIUserRepository {
	mock := &MockIUserRepository{ctrl: ctrl}
	mock.recorder = &MockIUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockIUserRepository) EXPECT() *MockIUserRepositoryMockRecorder {
	return _m.recorder
}

// Login mocks base method
func (_m *MockIUserRepository) Login(item *schema.LoginBodyParam) (*models.User, error) {
	ret := _m.ctrl.Call(_m, "Login", item)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login
func (_mr *MockIUserRepositoryMockRecorder) Login(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Login", reflect.TypeOf((*MockIUserRepository)(nil).Login), arg0)
}

// Register mocks base method
func (_m *MockIUserRepository) Register(item *schema.RegisterBodyParam) (*models.User, error) {
	ret := _m.ctrl.Call(_m, "Register", item)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register
func (_mr *MockIUserRepositoryMockRecorder) Register(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Register", reflect.TypeOf((*MockIUserRepository)(nil).Register), arg0)
}

// GetUserByID mocks base method
func (_m *MockIUserRepository) GetUserByID(id string) (*models.User, error) {
	ret := _m.ctrl.Call(_m, "GetUserByID", id)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID
func (_mr *MockIUserRepositoryMockRecorder) GetUserByID(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetUserByID", reflect.TypeOf((*MockIUserRepository)(nil).GetUserByID), arg0)
}

// GetUserByToken mocks base method
func (_m *MockIUserRepository) GetUserByToken(token string) (*models.User, error) {
	ret := _m.ctrl.Call(_m, "GetUserByToken", token)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByToken indicates an expected call of GetUserByToken
func (_mr *MockIUserRepositoryMockRecorder) GetUserByToken(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetUserByToken", reflect.TypeOf((*MockIUserRepository)(nil).GetUserByToken), arg0)
}

// GetUsers mocks base method
func (_m *MockIUserRepository) GetUsers(queryParam *schema.UserQueryParam) (*[]models.User, error) {
	ret := _m.ctrl.Call(_m, "GetUsers", queryParam)
	ret0, _ := ret[0].(*[]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers
func (_mr *MockIUserRepositoryMockRecorder) GetUsers(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetUsers", reflect.TypeOf((*MockIUserRepository)(nil).GetUsers), arg0)
}

// Update mocks base method
func (_m *MockIUserRepository) Update(userId string, bodyParam *schema.UserUpdateBodyParam) (*models.User, error) {
	ret := _m.ctrl.Call(_m, "Update", userId, bodyParam)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (_mr *MockIUserRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Update", reflect.TypeOf((*MockIUserRepository)(nil).Update), arg0, arg1)
}

// RemoveToken mocks base method
func (_m *MockIUserRepository) RemoveToken(userId string) (*models.User, error) {
	ret := _m.ctrl.Call(_m, "RemoveToken", userId)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveToken indicates an expected call of RemoveToken
func (_mr *MockIUserRepositoryMockRecorder) RemoveToken(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RemoveToken", reflect.TypeOf((*MockIUserRepository)(nil).RemoveToken), arg0)
}

// Create mocks base method
func (_m *MockIUserRepository) Create(user *models.User) error {
	ret := _m.ctrl.Call(_m, "Create", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (_mr *MockIUserRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Create", reflect.TypeOf((*MockIUserRepository)(nil).Create), arg0)
}

// GetByID mocks base method
func (_m *MockIUserRepository) GetByID(id string) (*models.User, error) {
	ret := _m.ctrl.Call(_m, "GetByID", id)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (_mr *MockIUserRepositoryMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetByID", reflect.TypeOf((*MockIUserRepository)(nil).GetByID), arg0)
}

// List mocks base method
func (_m *MockIUserRepository) List(queryParam *schema.UserQueryParam) (*[]models.User, error) {
	ret := _m.ctrl.Call(_m, "List", queryParam)
	ret0, _ := ret[0].(*[]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (_mr *MockIUserRepositoryMockRecorder) List(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "List", reflect.TypeOf((*MockIUserRepository)(nil).List), arg0)
}
