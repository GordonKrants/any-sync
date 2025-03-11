// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/any-sync/net/peer (interfaces: Peer)
//
// Generated by this command:
//
//	mockgen -destination mock_peer/mock_peer.go github.com/anyproto/any-sync/net/peer Peer
//
// Package mock_peer is a generated GoMock package.
package mock_peer

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
	drpc "storj.io/drpc"
)

// MockPeer is a mock of Peer interface.
type MockPeer struct {
	ctrl     *gomock.Controller
	recorder *MockPeerMockRecorder
}

// MockPeerMockRecorder is the mock recorder for MockPeer.
type MockPeerMockRecorder struct {
	mock *MockPeer
}

// NewMockPeer creates a new mock instance.
func NewMockPeer(ctrl *gomock.Controller) *MockPeer {
	mock := &MockPeer{ctrl: ctrl}
	mock.recorder = &MockPeerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPeer) EXPECT() *MockPeerMockRecorder {
	return m.recorder
}

// AcquireDrpcConn mocks base method.
func (m *MockPeer) AcquireDrpcConn(arg0 context.Context) (drpc.Conn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcquireDrpcConn", arg0)
	ret0, _ := ret[0].(drpc.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcquireDrpcConn indicates an expected call of AcquireDrpcConn.
func (mr *MockPeerMockRecorder) AcquireDrpcConn(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcquireDrpcConn", reflect.TypeOf((*MockPeer)(nil).AcquireDrpcConn), arg0)
}

// Close mocks base method.
func (m *MockPeer) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockPeerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPeer)(nil).Close))
}

// CloseChan mocks base method.
func (m *MockPeer) CloseChan() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseChan")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// CloseChan indicates an expected call of CloseChan.
func (mr *MockPeerMockRecorder) CloseChan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseChan", reflect.TypeOf((*MockPeer)(nil).CloseChan))
}

// Context mocks base method.
func (m *MockPeer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockPeerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockPeer)(nil).Context))
}

// DoDrpc mocks base method.
func (m *MockPeer) DoDrpc(arg0 context.Context, arg1 func(drpc.Conn) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoDrpc", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoDrpc indicates an expected call of DoDrpc.
func (mr *MockPeerMockRecorder) DoDrpc(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoDrpc", reflect.TypeOf((*MockPeer)(nil).DoDrpc), arg0, arg1)
}

// Id mocks base method.
func (m *MockPeer) Id() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(string)
	return ret0
}

// Id indicates an expected call of Id.
func (mr *MockPeerMockRecorder) Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockPeer)(nil).Id))
}

// IsClosed mocks base method.
func (m *MockPeer) IsClosed() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsClosed")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsClosed indicates an expected call of IsClosed.
func (mr *MockPeerMockRecorder) IsClosed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsClosed", reflect.TypeOf((*MockPeer)(nil).IsClosed))
}

// ReleaseDrpcConn mocks base method.
func (m *MockPeer) ReleaseDrpcConn(arg0 drpc.Conn) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReleaseDrpcConn", arg0)
}

// ReleaseDrpcConn indicates an expected call of ReleaseDrpcConn.
func (mr *MockPeerMockRecorder) ReleaseDrpcConn(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseDrpcConn", reflect.TypeOf((*MockPeer)(nil).ReleaseDrpcConn), arg0)
}

// SetTTL mocks base method.
func (m *MockPeer) SetTTL(arg0 time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTTL", arg0)
}

// SetTTL indicates an expected call of SetTTL.
func (mr *MockPeerMockRecorder) SetTTL(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTTL", reflect.TypeOf((*MockPeer)(nil).SetTTL), arg0)
}

// TryClose mocks base method.
func (m *MockPeer) TryClose(arg0 time.Duration) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TryClose", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TryClose indicates an expected call of TryClose.
func (mr *MockPeerMockRecorder) TryClose(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TryClose", reflect.TypeOf((*MockPeer)(nil).TryClose), arg0)
}
