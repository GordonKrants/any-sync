// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/any-sync/commonspace/acl/aclclient (interfaces: AclJoiningClient,AclSpaceClient)
//
// Generated by this command:
//
//	mockgen -destination mock_aclclient/mock_aclclient.go github.com/anyproto/any-sync/commonspace/acl/aclclient AclJoiningClient,AclSpaceClient
//

// Package mock_aclclient is a generated GoMock package.
package mock_aclclient

import (
	context "context"
	reflect "reflect"

	app "github.com/anyproto/any-sync/app"
	list "github.com/anyproto/any-sync/commonspace/object/acl/list"
	consensusproto "github.com/anyproto/any-sync/consensus/consensusproto"
	crypto "github.com/anyproto/any-sync/util/crypto"
	gomock "go.uber.org/mock/gomock"
)

// MockAclJoiningClient is a mock of AclJoiningClient interface.
type MockAclJoiningClient struct {
	ctrl     *gomock.Controller
	recorder *MockAclJoiningClientMockRecorder
	isgomock struct{}
}

// MockAclJoiningClientMockRecorder is the mock recorder for MockAclJoiningClient.
type MockAclJoiningClientMockRecorder struct {
	mock *MockAclJoiningClient
}

// NewMockAclJoiningClient creates a new mock instance.
func NewMockAclJoiningClient(ctrl *gomock.Controller) *MockAclJoiningClient {
	mock := &MockAclJoiningClient{ctrl: ctrl}
	mock.recorder = &MockAclJoiningClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAclJoiningClient) EXPECT() *MockAclJoiningClientMockRecorder {
	return m.recorder
}

// AclGetRecords mocks base method.
func (m *MockAclJoiningClient) AclGetRecords(ctx context.Context, spaceId, aclHead string) ([]*consensusproto.RawRecordWithId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AclGetRecords", ctx, spaceId, aclHead)
	ret0, _ := ret[0].([]*consensusproto.RawRecordWithId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AclGetRecords indicates an expected call of AclGetRecords.
func (mr *MockAclJoiningClientMockRecorder) AclGetRecords(ctx, spaceId, aclHead any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AclGetRecords", reflect.TypeOf((*MockAclJoiningClient)(nil).AclGetRecords), ctx, spaceId, aclHead)
}

// CancelJoin mocks base method.
func (m *MockAclJoiningClient) CancelJoin(ctx context.Context, spaceId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelJoin", ctx, spaceId)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelJoin indicates an expected call of CancelJoin.
func (mr *MockAclJoiningClientMockRecorder) CancelJoin(ctx, spaceId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelJoin", reflect.TypeOf((*MockAclJoiningClient)(nil).CancelJoin), ctx, spaceId)
}

// CancelRemoveSelf mocks base method.
func (m *MockAclJoiningClient) CancelRemoveSelf(ctx context.Context, spaceId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelRemoveSelf", ctx, spaceId)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelRemoveSelf indicates an expected call of CancelRemoveSelf.
func (mr *MockAclJoiningClientMockRecorder) CancelRemoveSelf(ctx, spaceId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelRemoveSelf", reflect.TypeOf((*MockAclJoiningClient)(nil).CancelRemoveSelf), ctx, spaceId)
}

// Init mocks base method.
func (m *MockAclJoiningClient) Init(a *app.App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", a)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockAclJoiningClientMockRecorder) Init(a any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockAclJoiningClient)(nil).Init), a)
}

// Name mocks base method.
func (m *MockAclJoiningClient) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockAclJoiningClientMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockAclJoiningClient)(nil).Name))
}

// RequestJoin mocks base method.
func (m *MockAclJoiningClient) RequestJoin(ctx context.Context, spaceId string, payload list.RequestJoinPayload) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestJoin", ctx, spaceId, payload)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestJoin indicates an expected call of RequestJoin.
func (mr *MockAclJoiningClientMockRecorder) RequestJoin(ctx, spaceId, payload any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestJoin", reflect.TypeOf((*MockAclJoiningClient)(nil).RequestJoin), ctx, spaceId, payload)
}

// MockAclSpaceClient is a mock of AclSpaceClient interface.
type MockAclSpaceClient struct {
	ctrl     *gomock.Controller
	recorder *MockAclSpaceClientMockRecorder
	isgomock struct{}
}

// MockAclSpaceClientMockRecorder is the mock recorder for MockAclSpaceClient.
type MockAclSpaceClientMockRecorder struct {
	mock *MockAclSpaceClient
}

// NewMockAclSpaceClient creates a new mock instance.
func NewMockAclSpaceClient(ctrl *gomock.Controller) *MockAclSpaceClient {
	mock := &MockAclSpaceClient{ctrl: ctrl}
	mock.recorder = &MockAclSpaceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAclSpaceClient) EXPECT() *MockAclSpaceClientMockRecorder {
	return m.recorder
}

// AcceptRequest mocks base method.
func (m *MockAclSpaceClient) AcceptRequest(ctx context.Context, payload list.RequestAcceptPayload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptRequest", ctx, payload)
	ret0, _ := ret[0].(error)
	return ret0
}

// AcceptRequest indicates an expected call of AcceptRequest.
func (mr *MockAclSpaceClientMockRecorder) AcceptRequest(ctx, payload any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptRequest", reflect.TypeOf((*MockAclSpaceClient)(nil).AcceptRequest), ctx, payload)
}

// AddAccounts mocks base method.
func (m *MockAclSpaceClient) AddAccounts(ctx context.Context, add list.AccountsAddPayload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAccounts", ctx, add)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAccounts indicates an expected call of AddAccounts.
func (mr *MockAclSpaceClientMockRecorder) AddAccounts(ctx, add any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAccounts", reflect.TypeOf((*MockAclSpaceClient)(nil).AddAccounts), ctx, add)
}

// AddRecord mocks base method.
func (m *MockAclSpaceClient) AddRecord(ctx context.Context, consRec *consensusproto.RawRecord) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRecord", ctx, consRec)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRecord indicates an expected call of AddRecord.
func (mr *MockAclSpaceClientMockRecorder) AddRecord(ctx, consRec any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRecord", reflect.TypeOf((*MockAclSpaceClient)(nil).AddRecord), ctx, consRec)
}

// CancelRequest mocks base method.
func (m *MockAclSpaceClient) CancelRequest(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelRequest", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelRequest indicates an expected call of CancelRequest.
func (mr *MockAclSpaceClientMockRecorder) CancelRequest(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelRequest", reflect.TypeOf((*MockAclSpaceClient)(nil).CancelRequest), ctx)
}

// ChangePermissions mocks base method.
func (m *MockAclSpaceClient) ChangePermissions(ctx context.Context, permChange list.PermissionChangesPayload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePermissions", ctx, permChange)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePermissions indicates an expected call of ChangePermissions.
func (mr *MockAclSpaceClientMockRecorder) ChangePermissions(ctx, permChange any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePermissions", reflect.TypeOf((*MockAclSpaceClient)(nil).ChangePermissions), ctx, permChange)
}

// DeclineRequest mocks base method.
func (m *MockAclSpaceClient) DeclineRequest(ctx context.Context, identity crypto.PubKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeclineRequest", ctx, identity)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeclineRequest indicates an expected call of DeclineRequest.
func (mr *MockAclSpaceClientMockRecorder) DeclineRequest(ctx, identity any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeclineRequest", reflect.TypeOf((*MockAclSpaceClient)(nil).DeclineRequest), ctx, identity)
}

// GenerateInvite mocks base method.
func (m *MockAclSpaceClient) GenerateInvite() (list.InviteResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateInvite")
	ret0, _ := ret[0].(list.InviteResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateInvite indicates an expected call of GenerateInvite.
func (mr *MockAclSpaceClientMockRecorder) GenerateInvite() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateInvite", reflect.TypeOf((*MockAclSpaceClient)(nil).GenerateInvite))
}

// Init mocks base method.
func (m *MockAclSpaceClient) Init(a *app.App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", a)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockAclSpaceClientMockRecorder) Init(a any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockAclSpaceClient)(nil).Init), a)
}

// Name mocks base method.
func (m *MockAclSpaceClient) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockAclSpaceClientMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockAclSpaceClient)(nil).Name))
}

// RemoveAccounts mocks base method.
func (m *MockAclSpaceClient) RemoveAccounts(ctx context.Context, payload list.AccountRemovePayload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAccounts", ctx, payload)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAccounts indicates an expected call of RemoveAccounts.
func (mr *MockAclSpaceClientMockRecorder) RemoveAccounts(ctx, payload any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAccounts", reflect.TypeOf((*MockAclSpaceClient)(nil).RemoveAccounts), ctx, payload)
}

// RequestSelfRemove mocks base method.
func (m *MockAclSpaceClient) RequestSelfRemove(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestSelfRemove", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// RequestSelfRemove indicates an expected call of RequestSelfRemove.
func (mr *MockAclSpaceClientMockRecorder) RequestSelfRemove(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestSelfRemove", reflect.TypeOf((*MockAclSpaceClient)(nil).RequestSelfRemove), ctx)
}

// RevokeAllInvites mocks base method.
func (m *MockAclSpaceClient) RevokeAllInvites(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeAllInvites", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeAllInvites indicates an expected call of RevokeAllInvites.
func (mr *MockAclSpaceClientMockRecorder) RevokeAllInvites(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeAllInvites", reflect.TypeOf((*MockAclSpaceClient)(nil).RevokeAllInvites), ctx)
}

// RevokeInvite mocks base method.
func (m *MockAclSpaceClient) RevokeInvite(ctx context.Context, inviteRecordId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeInvite", ctx, inviteRecordId)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeInvite indicates an expected call of RevokeInvite.
func (mr *MockAclSpaceClientMockRecorder) RevokeInvite(ctx, inviteRecordId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeInvite", reflect.TypeOf((*MockAclSpaceClient)(nil).RevokeInvite), ctx, inviteRecordId)
}

// StopSharing mocks base method.
func (m *MockAclSpaceClient) StopSharing(ctx context.Context, readKeyChange list.ReadKeyChangePayload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopSharing", ctx, readKeyChange)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopSharing indicates an expected call of StopSharing.
func (mr *MockAclSpaceClientMockRecorder) StopSharing(ctx, readKeyChange any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopSharing", reflect.TypeOf((*MockAclSpaceClient)(nil).StopSharing), ctx, readKeyChange)
}
