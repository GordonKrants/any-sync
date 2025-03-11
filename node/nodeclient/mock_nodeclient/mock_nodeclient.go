// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/any-sync/node/nodeclient (interfaces: NodeClient)
//
// Generated by this command:
//
//	mockgen -destination mock_nodeclient/mock_nodeclient.go github.com/anyproto/any-sync/node/nodeclient NodeClient
//
// Package mock_nodeclient is a generated GoMock package.
package mock_nodeclient

import (
	context "context"
	reflect "reflect"

	app "github.com/anyproto/any-sync/app"
	consensusproto "github.com/anyproto/any-sync/consensus/consensusproto"
	gomock "go.uber.org/mock/gomock"
)

// MockNodeClient is a mock of NodeClient interface.
type MockNodeClient struct {
	ctrl     *gomock.Controller
	recorder *MockNodeClientMockRecorder
}

// MockNodeClientMockRecorder is the mock recorder for MockNodeClient.
type MockNodeClientMockRecorder struct {
	mock *MockNodeClient
}

// NewMockNodeClient creates a new mock instance.
func NewMockNodeClient(ctrl *gomock.Controller) *MockNodeClient {
	mock := &MockNodeClient{ctrl: ctrl}
	mock.recorder = &MockNodeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeClient) EXPECT() *MockNodeClientMockRecorder {
	return m.recorder
}

// AclAddRecord mocks base method.
func (m *MockNodeClient) AclAddRecord(arg0 context.Context, arg1 string, arg2 *consensusproto.RawRecord) (*consensusproto.RawRecordWithId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AclAddRecord", arg0, arg1, arg2)
	ret0, _ := ret[0].(*consensusproto.RawRecordWithId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AclAddRecord indicates an expected call of AclAddRecord.
func (mr *MockNodeClientMockRecorder) AclAddRecord(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AclAddRecord", reflect.TypeOf((*MockNodeClient)(nil).AclAddRecord), arg0, arg1, arg2)
}

// AclGetRecords mocks base method.
func (m *MockNodeClient) AclGetRecords(arg0 context.Context, arg1, arg2 string) ([]*consensusproto.RawRecordWithId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AclGetRecords", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*consensusproto.RawRecordWithId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AclGetRecords indicates an expected call of AclGetRecords.
func (mr *MockNodeClientMockRecorder) AclGetRecords(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AclGetRecords", reflect.TypeOf((*MockNodeClient)(nil).AclGetRecords), arg0, arg1, arg2)
}

// Init mocks base method.
func (m *MockNodeClient) Init(arg0 *app.App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockNodeClientMockRecorder) Init(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockNodeClient)(nil).Init), arg0)
}

// Name mocks base method.
func (m *MockNodeClient) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockNodeClientMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockNodeClient)(nil).Name))
}
