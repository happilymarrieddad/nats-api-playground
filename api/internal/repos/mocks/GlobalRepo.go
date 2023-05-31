// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/happilymarrieddad/nats-api-playground/api/internal/repos (interfaces: GlobalRepo)

// Package mock_repos is a generated GoMock package.
package mock_repos

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	repos "github.com/happilymarrieddad/nats-api-playground/api/internal/repos"
)

// MockGlobalRepo is a mock of GlobalRepo interface.
type MockGlobalRepo struct {
	ctrl     *gomock.Controller
	recorder *MockGlobalRepoMockRecorder
}

// MockGlobalRepoMockRecorder is the mock recorder for MockGlobalRepo.
type MockGlobalRepoMockRecorder struct {
	mock *MockGlobalRepo
}

// NewMockGlobalRepo creates a new mock instance.
func NewMockGlobalRepo(ctrl *gomock.Controller) *MockGlobalRepo {
	mock := &MockGlobalRepo{ctrl: ctrl}
	mock.recorder = &MockGlobalRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGlobalRepo) EXPECT() *MockGlobalRepoMockRecorder {
	return m.recorder
}

// Users mocks base method.
func (m *MockGlobalRepo) Users() repos.Users {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Users")
	ret0, _ := ret[0].(repos.Users)
	return ret0
}

// Users indicates an expected call of Users.
func (mr *MockGlobalRepoMockRecorder) Users() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Users", reflect.TypeOf((*MockGlobalRepo)(nil).Users))
}