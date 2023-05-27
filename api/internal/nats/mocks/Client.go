// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/happilymarrieddad/nats-api-playground/api/internal/nats (interfaces: Client)

// Package mock_nats is a generated GoMock package.
package mock_nats

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	nats "github.com/nats-io/nats.go"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// HandleAuthRequest mocks base method.
func (m *MockClient) HandleAuthRequest(arg0, arg1 string, arg2 func(*nats.Msg)) (*nats.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleAuthRequest", arg0, arg1, arg2)
	ret0, _ := ret[0].(*nats.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleAuthRequest indicates an expected call of HandleAuthRequest.
func (mr *MockClientMockRecorder) HandleAuthRequest(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleAuthRequest", reflect.TypeOf((*MockClient)(nil).HandleAuthRequest), arg0, arg1, arg2)
}

// HandleRequest mocks base method.
func (m *MockClient) HandleRequest(arg0, arg1 string, arg2 func(*nats.Msg)) (*nats.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleRequest", arg0, arg1, arg2)
	ret0, _ := ret[0].(*nats.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleRequest indicates an expected call of HandleRequest.
func (mr *MockClientMockRecorder) HandleRequest(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleRequest", reflect.TypeOf((*MockClient)(nil).HandleRequest), arg0, arg1, arg2)
}

// Request mocks base method.
func (m *MockClient) Request(arg0 string, arg1 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Request", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Request indicates an expected call of Request.
func (mr *MockClientMockRecorder) Request(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockClient)(nil).Request), arg0, arg1)
}

// Respond mocks base method.
func (m *MockClient) Respond(arg0 string, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Respond", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Respond indicates an expected call of Respond.
func (mr *MockClientMockRecorder) Respond(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Respond", reflect.TypeOf((*MockClient)(nil).Respond), arg0, arg1)
}
