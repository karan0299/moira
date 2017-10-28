// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/moira-alert/moira (interfaces: Sender)

// Package mock_moira_alert is a generated GoMock package.
package mock_moira_alert

import (
	gomock "github.com/golang/mock/gomock"
	moira "github.com/moira-alert/moira"
	reflect "reflect"
	time "time"
)

// MockSender is a mock of Sender interface
type MockSender struct {
	ctrl     *gomock.Controller
	recorder *MockSenderMockRecorder
}

// MockSenderMockRecorder is the mock recorder for MockSender
type MockSenderMockRecorder struct {
	mock *MockSender
}

// NewMockSender creates a new mock instance
func NewMockSender(ctrl *gomock.Controller) *MockSender {
	mock := &MockSender{ctrl: ctrl}
	mock.recorder = &MockSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSender) EXPECT() *MockSenderMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockSender) Init(arg0 map[string]string, arg1 moira.Logger, arg2 *time.Location) error {
	ret := m.ctrl.Call(m, "Init", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockSenderMockRecorder) Init(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockSender)(nil).Init), arg0, arg1, arg2)
}

// SendEvents mocks base method
func (m *MockSender) SendEvents(arg0 moira.NotificationEvents, arg1 moira.ContactData, arg2 moira.TriggerData, arg3 bool) error {
	ret := m.ctrl.Call(m, "SendEvents", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendEvents indicates an expected call of SendEvents
func (mr *MockSenderMockRecorder) SendEvents(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEvents", reflect.TypeOf((*MockSender)(nil).SendEvents), arg0, arg1, arg2, arg3)
}
