// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vmware-tanzu/antrea/pkg/monitor (interfaces: AgentQuerier)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	reflect "reflect"
)

// MockAgentQuerier is a mock of AgentQuerier interface
type MockAgentQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockAgentQuerierMockRecorder
}

// MockAgentQuerierMockRecorder is the mock recorder for MockAgentQuerier
type MockAgentQuerierMockRecorder struct {
	mock *MockAgentQuerier
}

// NewMockAgentQuerier creates a new mock instance
func NewMockAgentQuerier(ctrl *gomock.Controller) *MockAgentQuerier {
	mock := &MockAgentQuerier{ctrl: ctrl}
	mock.recorder = &MockAgentQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAgentQuerier) EXPECT() *MockAgentQuerierMockRecorder {
	return m.recorder
}

// GetLocalPodNum mocks base method
func (m *MockAgentQuerier) GetLocalPodNum() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocalPodNum")
	ret0, _ := ret[0].(int32)
	return ret0
}

// GetLocalPodNum indicates an expected call of GetLocalPodNum
func (mr *MockAgentQuerierMockRecorder) GetLocalPodNum() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocalPodNum", reflect.TypeOf((*MockAgentQuerier)(nil).GetLocalPodNum))
}

// GetNodeSubnet mocks base method
func (m *MockAgentQuerier) GetNodeSubnet() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodeSubnet")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetNodeSubnet indicates an expected call of GetNodeSubnet
func (mr *MockAgentQuerierMockRecorder) GetNodeSubnet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeSubnet", reflect.TypeOf((*MockAgentQuerier)(nil).GetNodeSubnet))
}

// GetOVSFlowTable mocks base method
func (m *MockAgentQuerier) GetOVSFlowTable() map[string]int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOVSFlowTable")
	ret0, _ := ret[0].(map[string]int32)
	return ret0
}

// GetOVSFlowTable indicates an expected call of GetOVSFlowTable
func (mr *MockAgentQuerierMockRecorder) GetOVSFlowTable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOVSFlowTable", reflect.TypeOf((*MockAgentQuerier)(nil).GetOVSFlowTable))
}

// GetSelfNode mocks base method
func (m *MockAgentQuerier) GetSelfNode() v1.ObjectReference {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSelfNode")
	ret0, _ := ret[0].(v1.ObjectReference)
	return ret0
}

// GetSelfNode indicates an expected call of GetSelfNode
func (mr *MockAgentQuerierMockRecorder) GetSelfNode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSelfNode", reflect.TypeOf((*MockAgentQuerier)(nil).GetSelfNode))
}

// GetSelfPod mocks base method
func (m *MockAgentQuerier) GetSelfPod() v1.ObjectReference {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSelfPod")
	ret0, _ := ret[0].(v1.ObjectReference)
	return ret0
}

// GetSelfPod indicates an expected call of GetSelfPod
func (mr *MockAgentQuerierMockRecorder) GetSelfPod() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSelfPod", reflect.TypeOf((*MockAgentQuerier)(nil).GetSelfPod))
}

// GetVersion mocks base method
func (m *MockAgentQuerier) GetVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetVersion indicates an expected call of GetVersion
func (mr *MockAgentQuerierMockRecorder) GetVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersion", reflect.TypeOf((*MockAgentQuerier)(nil).GetVersion))
}
