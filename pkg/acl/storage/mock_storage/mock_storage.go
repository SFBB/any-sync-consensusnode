// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anytypeio/go-anytype-infrastructure-experiments/pkg/acl/storage (interfaces: ListStorage,TreeStorage)

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	context "context"
	reflect "reflect"

	aclrecordproto "github.com/anytypeio/go-anytype-infrastructure-experiments/pkg/acl/aclrecordproto"
	treechangeproto "github.com/anytypeio/go-anytype-infrastructure-experiments/pkg/acl/treechangeproto"
	gomock "github.com/golang/mock/gomock"
)

// MockListStorage is a mock of ListStorage interface.
type MockListStorage struct {
	ctrl     *gomock.Controller
	recorder *MockListStorageMockRecorder
}

// MockListStorageMockRecorder is the mock recorder for MockListStorage.
type MockListStorageMockRecorder struct {
	mock *MockListStorage
}

// NewMockListStorage creates a new mock instance.
func NewMockListStorage(ctrl *gomock.Controller) *MockListStorage {
	mock := &MockListStorage{ctrl: ctrl}
	mock.recorder = &MockListStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockListStorage) EXPECT() *MockListStorageMockRecorder {
	return m.recorder
}

// AddRawRecord mocks base method.
func (m *MockListStorage) AddRawRecord(arg0 context.Context, arg1 *aclrecordproto.RawACLRecordWithId) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRawRecord", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRawRecord indicates an expected call of AddRawRecord.
func (mr *MockListStorageMockRecorder) AddRawRecord(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRawRecord", reflect.TypeOf((*MockListStorage)(nil).AddRawRecord), arg0, arg1)
}

// GetRawRecord mocks base method.
func (m *MockListStorage) GetRawRecord(arg0 context.Context, arg1 string) (*aclrecordproto.RawACLRecordWithId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRawRecord", arg0, arg1)
	ret0, _ := ret[0].(*aclrecordproto.RawACLRecordWithId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRawRecord indicates an expected call of GetRawRecord.
func (mr *MockListStorageMockRecorder) GetRawRecord(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawRecord", reflect.TypeOf((*MockListStorage)(nil).GetRawRecord), arg0, arg1)
}

// Head mocks base method.
func (m *MockListStorage) Head() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Head")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Head indicates an expected call of Head.
func (mr *MockListStorageMockRecorder) Head() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Head", reflect.TypeOf((*MockListStorage)(nil).Head))
}

// ID mocks base method.
func (m *MockListStorage) ID() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ID indicates an expected call of ID.
func (mr *MockListStorageMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockListStorage)(nil).ID))
}

// Root mocks base method.
func (m *MockListStorage) Root() (*aclrecordproto.RawACLRecordWithId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Root")
	ret0, _ := ret[0].(*aclrecordproto.RawACLRecordWithId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Root indicates an expected call of Root.
func (mr *MockListStorageMockRecorder) Root() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Root", reflect.TypeOf((*MockListStorage)(nil).Root))
}

// SetHead mocks base method.
func (m *MockListStorage) SetHead(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHead", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHead indicates an expected call of SetHead.
func (mr *MockListStorageMockRecorder) SetHead(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHead", reflect.TypeOf((*MockListStorage)(nil).SetHead), arg0)
}

// MockTreeStorage is a mock of TreeStorage interface.
type MockTreeStorage struct {
	ctrl     *gomock.Controller
	recorder *MockTreeStorageMockRecorder
}

// MockTreeStorageMockRecorder is the mock recorder for MockTreeStorage.
type MockTreeStorageMockRecorder struct {
	mock *MockTreeStorage
}

// NewMockTreeStorage creates a new mock instance.
func NewMockTreeStorage(ctrl *gomock.Controller) *MockTreeStorage {
	mock := &MockTreeStorage{ctrl: ctrl}
	mock.recorder = &MockTreeStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTreeStorage) EXPECT() *MockTreeStorageMockRecorder {
	return m.recorder
}

// AddRawChange mocks base method.
func (m *MockTreeStorage) AddRawChange(arg0 *treechangeproto.RawTreeChangeWithId) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRawChange", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRawChange indicates an expected call of AddRawChange.
func (mr *MockTreeStorageMockRecorder) AddRawChange(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRawChange", reflect.TypeOf((*MockTreeStorage)(nil).AddRawChange), arg0)
}

// GetRawChange mocks base method.
func (m *MockTreeStorage) GetRawChange(arg0 context.Context, arg1 string) (*treechangeproto.RawTreeChangeWithId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRawChange", arg0, arg1)
	ret0, _ := ret[0].(*treechangeproto.RawTreeChangeWithId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRawChange indicates an expected call of GetRawChange.
func (mr *MockTreeStorageMockRecorder) GetRawChange(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawChange", reflect.TypeOf((*MockTreeStorage)(nil).GetRawChange), arg0, arg1)
}

// HasChange mocks base method.
func (m *MockTreeStorage) HasChange(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasChange", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasChange indicates an expected call of HasChange.
func (mr *MockTreeStorageMockRecorder) HasChange(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasChange", reflect.TypeOf((*MockTreeStorage)(nil).HasChange), arg0, arg1)
}

// Heads mocks base method.
func (m *MockTreeStorage) Heads() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Heads")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Heads indicates an expected call of Heads.
func (mr *MockTreeStorageMockRecorder) Heads() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Heads", reflect.TypeOf((*MockTreeStorage)(nil).Heads))
}

// ID mocks base method.
func (m *MockTreeStorage) ID() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ID indicates an expected call of ID.
func (mr *MockTreeStorageMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockTreeStorage)(nil).ID))
}

// Root mocks base method.
func (m *MockTreeStorage) Root() (*treechangeproto.RawTreeChangeWithId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Root")
	ret0, _ := ret[0].(*treechangeproto.RawTreeChangeWithId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Root indicates an expected call of Root.
func (mr *MockTreeStorageMockRecorder) Root() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Root", reflect.TypeOf((*MockTreeStorage)(nil).Root))
}

// SetHeads mocks base method.
func (m *MockTreeStorage) SetHeads(arg0 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeads", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeads indicates an expected call of SetHeads.
func (mr *MockTreeStorageMockRecorder) SetHeads(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeads", reflect.TypeOf((*MockTreeStorage)(nil).SetHeads), arg0)
}
