// Code generated by MockGen. DO NOT EDIT.
// Source: app/repositories/role.go

package mocks

import (
	gomock "github.com/golang/mock/gomock"
	models "github.com/quangdangfit/go-admin/app/models"
	schema "github.com/quangdangfit/go-admin/app/schema"
	reflect "reflect"
)

// MockIRoleRepository is a mock of IRoleRepository interface
type MockIRoleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRoleRepositoryMockRecorder
}

// MockIRoleRepositoryMockRecorder is the mock recorder for MockIRoleRepository
type MockIRoleRepositoryMockRecorder struct {
	mock *MockIRoleRepository
}

// NewMockIRoleRepository creates a new mock instance
func NewMockIRoleRepository(ctrl *gomock.Controller) *MockIRoleRepository {
	mock := &MockIRoleRepository{ctrl: ctrl}
	mock.recorder = &MockIRoleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockIRoleRepository) EXPECT() *MockIRoleRepositoryMockRecorder {
	return _m.recorder
}

// CreateRole mocks base method
func (_m *MockIRoleRepository) CreateRole(req *schema.RoleBodyParam) (*models.Role, error) {
	ret := _m.ctrl.Call(_m, "CreateRole", req)
	ret0, _ := ret[0].(*models.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRole indicates an expected call of CreateRole
func (_mr *MockIRoleRepositoryMockRecorder) CreateRole(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CreateRole", reflect.TypeOf((*MockIRoleRepository)(nil).CreateRole), arg0)
}

// GetRoleByName mocks base method
func (_m *MockIRoleRepository) GetRoleByName(name string) (*models.Role, error) {
	ret := _m.ctrl.Call(_m, "GetRoleByName", name)
	ret0, _ := ret[0].(*models.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoleByName indicates an expected call of GetRoleByName
func (_mr *MockIRoleRepositoryMockRecorder) GetRoleByName(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetRoleByName", reflect.TypeOf((*MockIRoleRepository)(nil).GetRoleByName), arg0)
}
