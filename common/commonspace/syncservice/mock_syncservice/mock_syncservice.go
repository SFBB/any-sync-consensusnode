// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anytypeio/go-anytype-infrastructure-experiments/common/commonspace/syncservice (interfaces: SyncClient)

// Package mock_syncservice is a generated GoMock package.
package mock_syncservice

import (
	"github.com/anytypeio/go-anytype-infrastructure-experiments/common/pkg/acl/tree"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/common/pkg/acl/treechangeproto"
	reflect "reflect"

	spacesyncproto "github.com/anytypeio/go-anytype-infrastructure-experiments/common/commonspace/spacesyncproto"
	gomock "github.com/golang/mock/gomock"
)

// MockSyncClient is a mock of SyncClient interface.
type MockSyncClient struct {
	ctrl     *gomock.Controller
	recorder *MockSyncClientMockRecorder
}

// MockSyncClientMockRecorder is the mock recorder for MockSyncClient.
type MockSyncClientMockRecorder struct {
	mock *MockSyncClient
}

// NewMockSyncClient creates a new mock instance.
func NewMockSyncClient(ctrl *gomock.Controller) *MockSyncClient {
	mock := &MockSyncClient{ctrl: ctrl}
	mock.recorder = &MockSyncClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSyncClient) EXPECT() *MockSyncClientMockRecorder {
	return m.recorder
}

// AddAndReadStreamAsync mocks base method.
func (m *MockSyncClient) AddAndReadStreamAsync(arg0 spacesyncproto.DRPCSpace_StreamStream) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddAndReadStreamAsync", arg0)
}

// AddAndReadStreamAsync indicates an expected call of AddAndReadStreamAsync.
func (mr *MockSyncClientMockRecorder) AddAndReadStreamAsync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAndReadStreamAsync", reflect.TypeOf((*MockSyncClient)(nil).AddAndReadStreamAsync), arg0)
}

// AddAndReadStreamSync mocks base method.
func (m *MockSyncClient) AddAndReadStreamSync(arg0 spacesyncproto.DRPCSpace_StreamStream) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAndReadStreamSync", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAndReadStreamSync indicates an expected call of AddAndReadStreamSync.
func (mr *MockSyncClientMockRecorder) AddAndReadStreamSync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAndReadStreamSync", reflect.TypeOf((*MockSyncClient)(nil).AddAndReadStreamSync), arg0)
}

// BroadcastAsync mocks base method.
func (m *MockSyncClient) BroadcastAsync(arg0 *spacesyncproto.ObjectSyncMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BroadcastAsync", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// BroadcastAsync indicates an expected call of BroadcastAsync.
func (mr *MockSyncClientMockRecorder) BroadcastAsync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastAsync", reflect.TypeOf((*MockSyncClient)(nil).BroadcastAsync), arg0)
}

// BroadcastAsyncOrSendResponsible mocks base method.
func (m *MockSyncClient) BroadcastAsyncOrSendResponsible(arg0 *spacesyncproto.ObjectSyncMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BroadcastAsyncOrSendResponsible", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// BroadcastAsyncOrSendResponsible indicates an expected call of BroadcastAsyncOrSendResponsible.
func (mr *MockSyncClientMockRecorder) BroadcastAsyncOrSendResponsible(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastAsyncOrSendResponsible", reflect.TypeOf((*MockSyncClient)(nil).BroadcastAsyncOrSendResponsible), arg0)
}

// Close mocks base method.
func (m *MockSyncClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockSyncClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSyncClient)(nil).Close))
}

// CreateFullSyncRequest mocks base method.
func (m *MockSyncClient) CreateFullSyncRequest(arg0 tree.ObjectTree, arg1, arg2 []string, arg3 string) (*spacesyncproto.ObjectSyncMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFullSyncRequest", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*spacesyncproto.ObjectSyncMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFullSyncRequest indicates an expected call of CreateFullSyncRequest.
func (mr *MockSyncClientMockRecorder) CreateFullSyncRequest(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFullSyncRequest", reflect.TypeOf((*MockSyncClient)(nil).CreateFullSyncRequest), arg0, arg1, arg2, arg3)
}

// CreateFullSyncResponse mocks base method.
func (m *MockSyncClient) CreateFullSyncResponse(arg0 tree.ObjectTree, arg1, arg2 []string, arg3 string) (*spacesyncproto.ObjectSyncMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFullSyncResponse", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*spacesyncproto.ObjectSyncMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFullSyncResponse indicates an expected call of CreateFullSyncResponse.
func (mr *MockSyncClientMockRecorder) CreateFullSyncResponse(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFullSyncResponse", reflect.TypeOf((*MockSyncClient)(nil).CreateFullSyncResponse), arg0, arg1, arg2, arg3)
}

// CreateHeadUpdate mocks base method.
func (m *MockSyncClient) CreateHeadUpdate(arg0 tree.ObjectTree, arg1 []*treechangeproto.RawTreeChangeWithId) *spacesyncproto.ObjectSyncMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHeadUpdate", arg0, arg1)
	ret0, _ := ret[0].(*spacesyncproto.ObjectSyncMessage)
	return ret0
}

// CreateHeadUpdate indicates an expected call of CreateHeadUpdate.
func (mr *MockSyncClientMockRecorder) CreateHeadUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHeadUpdate", reflect.TypeOf((*MockSyncClient)(nil).CreateHeadUpdate), arg0, arg1)
}

// CreateNewTreeRequest mocks base method.
func (m *MockSyncClient) CreateNewTreeRequest(arg0 string) *spacesyncproto.ObjectSyncMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewTreeRequest", arg0)
	ret0, _ := ret[0].(*spacesyncproto.ObjectSyncMessage)
	return ret0
}

// CreateNewTreeRequest indicates an expected call of CreateNewTreeRequest.
func (mr *MockSyncClientMockRecorder) CreateNewTreeRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewTreeRequest", reflect.TypeOf((*MockSyncClient)(nil).CreateNewTreeRequest), arg0)
}

// HasActiveStream mocks base method.
func (m *MockSyncClient) HasActiveStream(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasActiveStream", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasActiveStream indicates an expected call of HasActiveStream.
func (mr *MockSyncClientMockRecorder) HasActiveStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasActiveStream", reflect.TypeOf((*MockSyncClient)(nil).HasActiveStream), arg0)
}

// SendAsync mocks base method.
func (m *MockSyncClient) SendAsync(arg0 []string, arg1 *spacesyncproto.ObjectSyncMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAsync", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendAsync indicates an expected call of SendAsync.
func (mr *MockSyncClientMockRecorder) SendAsync(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAsync", reflect.TypeOf((*MockSyncClient)(nil).SendAsync), arg0, arg1)
}

// SendSync mocks base method.
func (m *MockSyncClient) SendSync(arg0 string, arg1 *spacesyncproto.ObjectSyncMessage) (*spacesyncproto.ObjectSyncMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendSync", arg0, arg1)
	ret0, _ := ret[0].(*spacesyncproto.ObjectSyncMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendSync indicates an expected call of SendSync.
func (mr *MockSyncClientMockRecorder) SendSync(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendSync", reflect.TypeOf((*MockSyncClient)(nil).SendSync), arg0, arg1)
}
