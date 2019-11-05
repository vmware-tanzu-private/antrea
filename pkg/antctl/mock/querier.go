// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vmware-tanzu/antrea/pkg/monitor (interfaces: Querier)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	reflect "reflect"
)

// MockQuerier is a mock of Querier interface
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// GetSelfNode mocks base method
func (m *MockQuerier) GetSelfNode() v1.ObjectReference {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSelfNode")
	ret0, _ := ret[0].(v1.ObjectReference)
	return ret0
}

// GetSelfNode indicates an expected call of GetSelfNode
func (mr *MockQuerierMockRecorder) GetSelfNode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSelfNode", reflect.TypeOf((*MockQuerier)(nil).GetSelfNode))
}

// GetSelfPod mocks base method
func (m *MockQuerier) GetSelfPod() v1.ObjectReference {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSelfPod")
	ret0, _ := ret[0].(v1.ObjectReference)
	return ret0
}

// GetSelfPod indicates an expected call of GetSelfPod
func (mr *MockQuerierMockRecorder) GetSelfPod() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSelfPod", reflect.TypeOf((*MockQuerier)(nil).GetSelfPod))
}

// GetVersion mocks base method
func (m *MockQuerier) GetVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetVersion indicates an expected call of GetVersion
func (mr *MockQuerierMockRecorder) GetVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersion", reflect.TypeOf((*MockQuerier)(nil).GetVersion))
}
