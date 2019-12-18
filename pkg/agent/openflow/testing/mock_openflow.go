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
// Source: github.com/vmware-tanzu/antrea/pkg/agent/openflow (interfaces: Client,FlowOperations)

// Package testing is a generated GoMock package.
package testing

import (
	gomock "github.com/golang/mock/gomock"
	types "github.com/vmware-tanzu/antrea/pkg/agent/types"
	openflow "github.com/vmware-tanzu/antrea/pkg/ovs/openflow"
	net "net"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// AddPolicyRuleAddress mocks base method
func (m *MockClient) AddPolicyRuleAddress(arg0 uint32, arg1 types.AddressType, arg2 []types.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPolicyRuleAddress", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPolicyRuleAddress indicates an expected call of AddPolicyRuleAddress
func (mr *MockClientMockRecorder) AddPolicyRuleAddress(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPolicyRuleAddress", reflect.TypeOf((*MockClient)(nil).AddPolicyRuleAddress), arg0, arg1, arg2)
}

// DeletePolicyRuleAddress mocks base method
func (m *MockClient) DeletePolicyRuleAddress(arg0 uint32, arg1 types.AddressType, arg2 []types.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePolicyRuleAddress", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePolicyRuleAddress indicates an expected call of DeletePolicyRuleAddress
func (mr *MockClientMockRecorder) DeletePolicyRuleAddress(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePolicyRuleAddress", reflect.TypeOf((*MockClient)(nil).DeletePolicyRuleAddress), arg0, arg1, arg2)
}

// DeleteStaleFlows mocks base method
func (m *MockClient) DeleteStaleFlows() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStaleFlows")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStaleFlows indicates an expected call of DeleteStaleFlows
func (mr *MockClientMockRecorder) DeleteStaleFlows() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStaleFlows", reflect.TypeOf((*MockClient)(nil).DeleteStaleFlows))
}

// Disconnect mocks base method
func (m *MockClient) Disconnect() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disconnect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Disconnect indicates an expected call of Disconnect
func (mr *MockClientMockRecorder) Disconnect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnect", reflect.TypeOf((*MockClient)(nil).Disconnect))
}

// GetFlowTableStatus mocks base method
func (m *MockClient) GetFlowTableStatus() []openflow.TableStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFlowTableStatus")
	ret0, _ := ret[0].([]openflow.TableStatus)
	return ret0
}

// GetFlowTableStatus indicates an expected call of GetFlowTableStatus
func (mr *MockClientMockRecorder) GetFlowTableStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlowTableStatus", reflect.TypeOf((*MockClient)(nil).GetFlowTableStatus))
}

// Initialize mocks base method
func (m *MockClient) Initialize(arg0 types.RoundInfo, _ *types.NodeConfig) (<-chan struct{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", arg0)
	ret0, _ := ret[0].(<-chan struct{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Initialize indicates an expected call of Initialize
func (mr *MockClientMockRecorder) Initialize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockClient)(nil).Initialize), arg0)
}

// InstallClusterServiceCIDRFlows mocks base method
func (m *MockClient) InstallClusterServiceCIDRFlows(arg0 *net.IPNet, arg1 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallClusterServiceCIDRFlows", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallClusterServiceCIDRFlows indicates an expected call of InstallClusterServiceCIDRFlows
func (mr *MockClientMockRecorder) InstallClusterServiceCIDRFlows(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallClusterServiceCIDRFlows", reflect.TypeOf((*MockClient)(nil).InstallClusterServiceCIDRFlows), arg0, arg1)
}

// InstallDefaultTunnelFlows mocks base method
func (m *MockClient) InstallDefaultTunnelFlows(arg0 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallDefaultTunnelFlows", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallDefaultTunnelFlows indicates an expected call of InstallDefaultTunnelFlows
func (mr *MockClientMockRecorder) InstallDefaultTunnelFlows(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallDefaultTunnelFlows", reflect.TypeOf((*MockClient)(nil).InstallDefaultTunnelFlows), arg0)
}

// InstallGatewayFlows mocks base method
func (m *MockClient) InstallGatewayFlows(arg0 net.IP, arg1 net.HardwareAddr, arg2 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallGatewayFlows", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallGatewayFlows indicates an expected call of InstallGatewayFlows
func (mr *MockClientMockRecorder) InstallGatewayFlows(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallGatewayFlows", reflect.TypeOf((*MockClient)(nil).InstallGatewayFlows), arg0, arg1, arg2)
}

// InstallNodeFlows mocks base method
func (m *MockClient) InstallNodeFlows(arg0 string, arg1 net.HardwareAddr, arg2 net.IP, arg3 net.IPNet, arg4 net.IP, arg5 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallNodeFlows", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallNodeFlows indicates an expected call of InstallNodeFlows
func (mr *MockClientMockRecorder) InstallNodeFlows(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallNodeFlows", reflect.TypeOf((*MockClient)(nil).InstallNodeFlows), arg0, arg1, arg2, arg3, arg4, arg5)
}

// InstallPodFlows mocks base method
func (m *MockClient) InstallPodFlows(arg0 string, arg1 net.IP, arg2, arg3 net.HardwareAddr, arg4 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallPodFlows", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallPodFlows indicates an expected call of InstallPodFlows
func (mr *MockClientMockRecorder) InstallPodFlows(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallPodFlows", reflect.TypeOf((*MockClient)(nil).InstallPodFlows), arg0, arg1, arg2, arg3, arg4)
}

// InstallPolicyRuleFlows mocks base method
func (m *MockClient) InstallPolicyRuleFlows(arg0 *types.PolicyRule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallPolicyRuleFlows", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallPolicyRuleFlows indicates an expected call of InstallPolicyRuleFlows
func (mr *MockClientMockRecorder) InstallPolicyRuleFlows(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallPolicyRuleFlows", reflect.TypeOf((*MockClient)(nil).InstallPolicyRuleFlows), arg0)
}

// IsConnected mocks base method
func (m *MockClient) IsConnected() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsConnected")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsConnected indicates an expected call of IsConnected
func (mr *MockClientMockRecorder) IsConnected() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsConnected", reflect.TypeOf((*MockClient)(nil).IsConnected))
}

// ReplayFlows mocks base method
func (m *MockClient) ReplayFlows() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReplayFlows")
}

// ReplayFlows indicates an expected call of ReplayFlows
func (mr *MockClientMockRecorder) ReplayFlows() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplayFlows", reflect.TypeOf((*MockClient)(nil).ReplayFlows))
}

// UninstallNodeFlows mocks base method
func (m *MockClient) UninstallNodeFlows(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UninstallNodeFlows", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UninstallNodeFlows indicates an expected call of UninstallNodeFlows
func (mr *MockClientMockRecorder) UninstallNodeFlows(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UninstallNodeFlows", reflect.TypeOf((*MockClient)(nil).UninstallNodeFlows), arg0)
}

// UninstallPodFlows mocks base method
func (m *MockClient) UninstallPodFlows(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UninstallPodFlows", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UninstallPodFlows indicates an expected call of UninstallPodFlows
func (mr *MockClientMockRecorder) UninstallPodFlows(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UninstallPodFlows", reflect.TypeOf((*MockClient)(nil).UninstallPodFlows), arg0)
}

// UninstallPolicyRuleFlows mocks base method
func (m *MockClient) UninstallPolicyRuleFlows(arg0 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UninstallPolicyRuleFlows", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UninstallPolicyRuleFlows indicates an expected call of UninstallPolicyRuleFlows
func (mr *MockClientMockRecorder) UninstallPolicyRuleFlows(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UninstallPolicyRuleFlows", reflect.TypeOf((*MockClient)(nil).UninstallPolicyRuleFlows), arg0)
}

// MockFlowOperations is a mock of FlowOperations interface
type MockFlowOperations struct {
	ctrl     *gomock.Controller
	recorder *MockFlowOperationsMockRecorder
}

// MockFlowOperationsMockRecorder is the mock recorder for MockFlowOperations
type MockFlowOperationsMockRecorder struct {
	mock *MockFlowOperations
}

// NewMockFlowOperations creates a new mock instance
func NewMockFlowOperations(ctrl *gomock.Controller) *MockFlowOperations {
	mock := &MockFlowOperations{ctrl: ctrl}
	mock.recorder = &MockFlowOperationsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFlowOperations) EXPECT() *MockFlowOperationsMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockFlowOperations) Add(arg0 openflow.Flow) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockFlowOperationsMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockFlowOperations)(nil).Add), arg0)
}

// AddAll mocks base method
func (m *MockFlowOperations) AddAll(arg0 []openflow.Flow) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAll", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAll indicates an expected call of AddAll
func (mr *MockFlowOperationsMockRecorder) AddAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAll", reflect.TypeOf((*MockFlowOperations)(nil).AddAll), arg0)
}

// Delete mocks base method
func (m *MockFlowOperations) Delete(arg0 openflow.Flow) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockFlowOperationsMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFlowOperations)(nil).Delete), arg0)
}

// DeleteAll mocks base method
func (m *MockFlowOperations) DeleteAll(arg0 []openflow.Flow) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAll indicates an expected call of DeleteAll
func (mr *MockFlowOperationsMockRecorder) DeleteAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockFlowOperations)(nil).DeleteAll), arg0)
}

// Modify mocks base method
func (m *MockFlowOperations) Modify(arg0 openflow.Flow) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Modify", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Modify indicates an expected call of Modify
func (mr *MockFlowOperationsMockRecorder) Modify(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Modify", reflect.TypeOf((*MockFlowOperations)(nil).Modify), arg0)
}
