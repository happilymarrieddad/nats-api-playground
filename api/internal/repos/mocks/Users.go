// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/happilymarrieddad/nats-api-playground/api/internal/repos (interfaces: Users)

// Package mock_repos is a generated GoMock package.
package mock_repos

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	repos "github.com/happilymarrieddad/nats-api-playground/api/internal/repos"
	types "github.com/happilymarrieddad/nats-api-playground/api/types"
	xorm "xorm.io/xorm"
)

// MockUsers is a mock of Users interface.
type MockUsers struct {
	ctrl     *gomock.Controller
	recorder *MockUsersMockRecorder
}

// MockUsersMockRecorder is the mock recorder for MockUsers.
type MockUsersMockRecorder struct {
	mock *MockUsers
}

// NewMockUsers creates a new mock instance.
func NewMockUsers(ctrl *gomock.Controller) *MockUsers {
	mock := &MockUsers{ctrl: ctrl}
	mock.recorder = &MockUsersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsers) EXPECT() *MockUsersMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUsers) Create(arg0 context.Context, arg1 *types.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUsersMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUsers)(nil).Create), arg0, arg1)
}

// CreateTx mocks base method.
func (m *MockUsers) CreateTx(arg0 context.Context, arg1 *xorm.Session, arg2 *types.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTx", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTx indicates an expected call of CreateTx.
func (mr *MockUsersMockRecorder) CreateTx(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTx", reflect.TypeOf((*MockUsers)(nil).CreateTx), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockUsers) Delete(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUsersMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUsers)(nil).Delete), arg0, arg1)
}

// DeleteTx mocks base method.
func (m *MockUsers) DeleteTx(arg0 context.Context, arg1 *xorm.Session, arg2 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTx", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTx indicates an expected call of DeleteTx.
func (mr *MockUsersMockRecorder) DeleteTx(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTx", reflect.TypeOf((*MockUsers)(nil).DeleteTx), arg0, arg1, arg2)
}

// Find mocks base method.
func (m *MockUsers) Find(arg0 context.Context, arg1 *repos.UserFindOpts) ([]*types.User, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].([]*types.User)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Find indicates an expected call of Find.
func (mr *MockUsersMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockUsers)(nil).Find), arg0, arg1)
}

// Get mocks base method.
func (m *MockUsers) Get(arg0 context.Context, arg1 int64) (*types.User, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockUsersMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUsers)(nil).Get), arg0, arg1)
}

// GetByEmail mocks base method.
func (m *MockUsers) GetByEmail(arg0 context.Context, arg1 string) (*types.User, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", arg0, arg1)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByEmail indicates an expected call of GetByEmail.
func (mr *MockUsersMockRecorder) GetByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockUsers)(nil).GetByEmail), arg0, arg1)
}

// Update mocks base method.
func (m *MockUsers) Update(arg0 context.Context, arg1 types.UserUpdate) (*types.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUsersMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUsers)(nil).Update), arg0, arg1)
}

// UpdateTx mocks base method.
func (m *MockUsers) UpdateTx(arg0 context.Context, arg1 *xorm.Session, arg2 types.UserUpdate) (*types.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTx", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTx indicates an expected call of UpdateTx.
func (mr *MockUsersMockRecorder) UpdateTx(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTx", reflect.TypeOf((*MockUsers)(nil).UpdateTx), arg0, arg1, arg2)
}
