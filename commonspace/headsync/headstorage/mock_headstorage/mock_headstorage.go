// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/any-sync/commonspace/headsync/headstorage (interfaces: HeadStorage)
//
// Generated by this command:
//
//	mockgen -destination mock_headstorage/mock_headstorage.go github.com/anyproto/any-sync/commonspace/headsync/headstorage HeadStorage
//

// Package mock_headstorage is a generated GoMock package.
package mock_headstorage

import (
	context "context"
	reflect "reflect"

	headstorage "github.com/anyproto/any-sync/commonspace/headsync/headstorage"
	gomock "go.uber.org/mock/gomock"
)

// MockHeadStorage is a mock of HeadStorage interface.
type MockHeadStorage struct {
	ctrl     *gomock.Controller
	recorder *MockHeadStorageMockRecorder
}

// MockHeadStorageMockRecorder is the mock recorder for MockHeadStorage.
type MockHeadStorageMockRecorder struct {
	mock *MockHeadStorage
}

// NewMockHeadStorage creates a new mock instance.
func NewMockHeadStorage(ctrl *gomock.Controller) *MockHeadStorage {
	mock := &MockHeadStorage{ctrl: ctrl}
	mock.recorder = &MockHeadStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHeadStorage) EXPECT() *MockHeadStorageMockRecorder {
	return m.recorder
}

// AddObserver mocks base method.
func (m *MockHeadStorage) AddObserver(arg0 headstorage.Observer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddObserver", arg0)
}

// AddObserver indicates an expected call of AddObserver.
func (mr *MockHeadStorageMockRecorder) AddObserver(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddObserver", reflect.TypeOf((*MockHeadStorage)(nil).AddObserver), arg0)
}

// DeleteEntryTx mocks base method.
func (m *MockHeadStorage) DeleteEntryTx(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEntryTx", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEntryTx indicates an expected call of DeleteEntryTx.
func (mr *MockHeadStorageMockRecorder) DeleteEntryTx(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEntryTx", reflect.TypeOf((*MockHeadStorage)(nil).DeleteEntryTx), arg0, arg1)
}

// GetEntry mocks base method.
func (m *MockHeadStorage) GetEntry(arg0 context.Context, arg1 string) (headstorage.HeadsEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntry", arg0, arg1)
	ret0, _ := ret[0].(headstorage.HeadsEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntry indicates an expected call of GetEntry.
func (mr *MockHeadStorageMockRecorder) GetEntry(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntry", reflect.TypeOf((*MockHeadStorage)(nil).GetEntry), arg0, arg1)
}

// IterateEntries mocks base method.
func (m *MockHeadStorage) IterateEntries(arg0 context.Context, arg1 headstorage.IterOpts, arg2 headstorage.EntryIterator) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IterateEntries", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// IterateEntries indicates an expected call of IterateEntries.
func (mr *MockHeadStorageMockRecorder) IterateEntries(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateEntries", reflect.TypeOf((*MockHeadStorage)(nil).IterateEntries), arg0, arg1, arg2)
}

// UpdateEntry mocks base method.
func (m *MockHeadStorage) UpdateEntry(arg0 context.Context, arg1 headstorage.HeadsUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEntry", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEntry indicates an expected call of UpdateEntry.
func (mr *MockHeadStorageMockRecorder) UpdateEntry(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEntry", reflect.TypeOf((*MockHeadStorage)(nil).UpdateEntry), arg0, arg1)
}

// UpdateEntryTx mocks base method.
func (m *MockHeadStorage) UpdateEntryTx(arg0 context.Context, arg1 headstorage.HeadsUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEntryTx", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEntryTx indicates an expected call of UpdateEntryTx.
func (mr *MockHeadStorageMockRecorder) UpdateEntryTx(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEntryTx", reflect.TypeOf((*MockHeadStorage)(nil).UpdateEntryTx), arg0, arg1)
}
