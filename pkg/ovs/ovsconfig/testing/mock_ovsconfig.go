// Copyright 2019 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vmware-tanzu/antrea/pkg/ovs/ovsconfig (interfaces: OVSBridgeClient)

// Package testing is a generated GoMock package.
package testing

import (
	gomock "github.com/golang/mock/gomock"
	ovsconfig "github.com/vmware-tanzu/antrea/pkg/ovs/ovsconfig"
	reflect "reflect"
	time "time"
)

// MockOVSBridgeClient is a mock of OVSBridgeClient interface
type MockOVSBridgeClient struct {
	ctrl     *gomock.Controller
	recorder *MockOVSBridgeClientMockRecorder
}

// MockOVSBridgeClientMockRecorder is the mock recorder for MockOVSBridgeClient
type MockOVSBridgeClientMockRecorder struct {
	mock *MockOVSBridgeClient
}

// NewMockOVSBridgeClient creates a new mock instance
func NewMockOVSBridgeClient(ctrl *gomock.Controller) *MockOVSBridgeClient {
	mock := &MockOVSBridgeClient{ctrl: ctrl}
	mock.recorder = &MockOVSBridgeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOVSBridgeClient) EXPECT() *MockOVSBridgeClientMockRecorder {
	return m.recorder
}

// CheckConnectionHealth mocks base method
func (m *MockOVSBridgeClient) CheckConnectionHealth(arg0 time.Duration) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckConnectionHealth", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckConnectionHealth indicates an expected call of CheckConnectionHealth
func (mr *MockOVSBridgeClientMockRecorder) CheckConnectionHealth(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckConnectionHealth", reflect.TypeOf((*MockOVSBridgeClient)(nil).CheckConnectionHealth), arg0)
}

// Create mocks base method
func (m *MockOVSBridgeClient) Create() ovsconfig.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create")
	ret0, _ := ret[0].(ovsconfig.Error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockOVSBridgeClientMockRecorder) Create() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOVSBridgeClient)(nil).Create))
}

// CreateInternalPort mocks base method
func (m *MockOVSBridgeClient) CreateInternalPort(arg0 string, arg1 int32, arg2 map[string]interface{}) (string, ovsconfig.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInternalPort", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(ovsconfig.Error)
	return ret0, ret1
}

// CreateInternalPort indicates an expected call of CreateInternalPort
func (mr *MockOVSBridgeClientMockRecorder) CreateInternalPort(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInternalPort", reflect.TypeOf((*MockOVSBridgeClient)(nil).CreateInternalPort), arg0, arg1, arg2)
}

// CreatePort mocks base method
func (m *MockOVSBridgeClient) CreatePort(arg0, arg1 string, arg2 map[string]interface{}) (string, ovsconfig.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePort", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(ovsconfig.Error)
	return ret0, ret1
}

// CreatePort indicates an expected call of CreatePort
func (mr *MockOVSBridgeClientMockRecorder) CreatePort(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePort", reflect.TypeOf((*MockOVSBridgeClient)(nil).CreatePort), arg0, arg1, arg2)
}

// CreateTunnelPort mocks base method
func (m *MockOVSBridgeClient) CreateTunnelPort(arg0 string, arg1 ovsconfig.TunnelType, arg2 int32) (string, ovsconfig.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTunnelPort", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(ovsconfig.Error)
	return ret0, ret1
}

// CreateTunnelPort indicates an expected call of CreateTunnelPort
func (mr *MockOVSBridgeClientMockRecorder) CreateTunnelPort(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTunnelPort", reflect.TypeOf((*MockOVSBridgeClient)(nil).CreateTunnelPort), arg0, arg1, arg2)
}

// CreateTunnelPortExt mocks base method
func (m *MockOVSBridgeClient) CreateTunnelPortExt(arg0 string, arg1 ovsconfig.TunnelType, arg2 int32, arg3, arg4 string, arg5 map[string]interface{}) (string, ovsconfig.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTunnelPortExt", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(ovsconfig.Error)
	return ret0, ret1
}

// CreateTunnelPortExt indicates an expected call of CreateTunnelPortExt
func (mr *MockOVSBridgeClientMockRecorder) CreateTunnelPortExt(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTunnelPortExt", reflect.TypeOf((*MockOVSBridgeClient)(nil).CreateTunnelPortExt), arg0, arg1, arg2, arg3, arg4, arg5)
}

// Delete mocks base method
func (m *MockOVSBridgeClient) Delete() ovsconfig.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete")
	ret0, _ := ret[0].(ovsconfig.Error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockOVSBridgeClientMockRecorder) Delete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOVSBridgeClient)(nil).Delete))
}

// DeletePort mocks base method
func (m *MockOVSBridgeClient) DeletePort(arg0 string) ovsconfig.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePort", arg0)
	ret0, _ := ret[0].(ovsconfig.Error)
	return ret0
}

// DeletePort indicates an expected call of DeletePort
func (mr *MockOVSBridgeClientMockRecorder) DeletePort(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePort", reflect.TypeOf((*MockOVSBridgeClient)(nil).DeletePort), arg0)
}

// DeletePorts mocks base method
func (m *MockOVSBridgeClient) DeletePorts(arg0 []string) ovsconfig.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePorts", arg0)
	ret0, _ := ret[0].(ovsconfig.Error)
	return ret0
}

// DeletePorts indicates an expected call of DeletePorts
func (mr *MockOVSBridgeClientMockRecorder) DeletePorts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePorts", reflect.TypeOf((*MockOVSBridgeClient)(nil).DeletePorts), arg0)
}

// GetExternalIDs mocks base method
func (m *MockOVSBridgeClient) GetExternalIDs() (map[string]string, ovsconfig.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExternalIDs")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(ovsconfig.Error)
	return ret0, ret1
}

// GetExternalIDs indicates an expected call of GetExternalIDs
func (mr *MockOVSBridgeClientMockRecorder) GetExternalIDs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExternalIDs", reflect.TypeOf((*MockOVSBridgeClient)(nil).GetExternalIDs))
}

// GetOFPort mocks base method
func (m *MockOVSBridgeClient) GetOFPort(arg0 string) (int32, ovsconfig.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOFPort", arg0)
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(ovsconfig.Error)
	return ret0, ret1
}

// GetOFPort indicates an expected call of GetOFPort
func (mr *MockOVSBridgeClientMockRecorder) GetOFPort(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOFPort", reflect.TypeOf((*MockOVSBridgeClient)(nil).GetOFPort), arg0)
}

// GetOVSVersion mocks base method
func (m *MockOVSBridgeClient) GetOVSVersion() (string, ovsconfig.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOVSVersion")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(ovsconfig.Error)
	return ret0, ret1
}

// GetOVSVersion indicates an expected call of GetOVSVersion
func (mr *MockOVSBridgeClientMockRecorder) GetOVSVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOVSVersion", reflect.TypeOf((*MockOVSBridgeClient)(nil).GetOVSVersion))
}

// GetPortData mocks base method
func (m *MockOVSBridgeClient) GetPortData(arg0, arg1 string) (*ovsconfig.OVSPortData, ovsconfig.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPortData", arg0, arg1)
	ret0, _ := ret[0].(*ovsconfig.OVSPortData)
	ret1, _ := ret[1].(ovsconfig.Error)
	return ret0, ret1
}

// GetPortData indicates an expected call of GetPortData
func (mr *MockOVSBridgeClientMockRecorder) GetPortData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPortData", reflect.TypeOf((*MockOVSBridgeClient)(nil).GetPortData), arg0, arg1)
}

// GetPortList mocks base method
func (m *MockOVSBridgeClient) GetPortList() ([]ovsconfig.OVSPortData, ovsconfig.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPortList")
	ret0, _ := ret[0].([]ovsconfig.OVSPortData)
	ret1, _ := ret[1].(ovsconfig.Error)
	return ret0, ret1
}

// GetPortList indicates an expected call of GetPortList
func (mr *MockOVSBridgeClientMockRecorder) GetPortList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPortList", reflect.TypeOf((*MockOVSBridgeClient)(nil).GetPortList))
}

// SetExternalIDs mocks base method
func (m *MockOVSBridgeClient) SetExternalIDs(arg0 map[string]interface{}) ovsconfig.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetExternalIDs", arg0)
	ret0, _ := ret[0].(ovsconfig.Error)
	return ret0
}

// SetExternalIDs indicates an expected call of SetExternalIDs
func (mr *MockOVSBridgeClientMockRecorder) SetExternalIDs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetExternalIDs", reflect.TypeOf((*MockOVSBridgeClient)(nil).SetExternalIDs), arg0)
}

// SetInterfaceMTU mocks base method
func (m *MockOVSBridgeClient) SetInterfaceMTU(arg0 string, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetInterfaceMTU", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetInterfaceMTU indicates an expected call of SetInterfaceMTU
func (mr *MockOVSBridgeClientMockRecorder) SetInterfaceMTU(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInterfaceMTU", reflect.TypeOf((*MockOVSBridgeClient)(nil).SetInterfaceMTU), arg0, arg1)
}
