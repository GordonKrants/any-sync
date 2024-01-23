// Code generated by MockGen. DO NOT EDIT.
// Source: paymentservice/paymentserviceclient/paymentserviceclient.go
//
// Generated by this command:
//
//	mockgen -source paymentservice/paymentserviceclient/paymentserviceclient.go
//
// Package mock_paymentserviceclient is a generated GoMock package.
package mock_paymentserviceclient

import (
	context "context"
	reflect "reflect"

	app "github.com/anyproto/any-sync/app"
	paymentserviceproto "github.com/anyproto/any-sync/paymentservice/paymentserviceproto"
	gomock "go.uber.org/mock/gomock"
)

// MockAnyPpClientService is a mock of AnyPpClientService interface.
type MockAnyPpClientService struct {
	ctrl     *gomock.Controller
	recorder *MockAnyPpClientServiceMockRecorder
}

// MockAnyPpClientServiceMockRecorder is the mock recorder for MockAnyPpClientService.
type MockAnyPpClientServiceMockRecorder struct {
	mock *MockAnyPpClientService
}

// NewMockAnyPpClientService creates a new mock instance.
func NewMockAnyPpClientService(ctrl *gomock.Controller) *MockAnyPpClientService {
	mock := &MockAnyPpClientService{ctrl: ctrl}
	mock.recorder = &MockAnyPpClientServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAnyPpClientService) EXPECT() *MockAnyPpClientServiceMockRecorder {
	return m.recorder
}

// BuySubscription mocks base method.
func (m *MockAnyPpClientService) BuySubscription(ctx context.Context, in *paymentserviceproto.BuySubscriptionRequestSigned) (*paymentserviceproto.BuySubscriptionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuySubscription", ctx, in)
	ret0, _ := ret[0].(*paymentserviceproto.BuySubscriptionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuySubscription indicates an expected call of BuySubscription.
func (mr *MockAnyPpClientServiceMockRecorder) BuySubscription(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuySubscription", reflect.TypeOf((*MockAnyPpClientService)(nil).BuySubscription), ctx, in)
}

// GetSubscriptionStatus mocks base method.
func (m *MockAnyPpClientService) GetSubscriptionStatus(ctx context.Context, in *paymentserviceproto.GetSubscriptionRequestSigned) (*paymentserviceproto.GetSubscriptionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscriptionStatus", ctx, in)
	ret0, _ := ret[0].(*paymentserviceproto.GetSubscriptionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscriptionStatus indicates an expected call of GetSubscriptionStatus.
func (mr *MockAnyPpClientServiceMockRecorder) GetSubscriptionStatus(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscriptionStatus", reflect.TypeOf((*MockAnyPpClientService)(nil).GetSubscriptionStatus), ctx, in)
}

// Init mocks base method.
func (m *MockAnyPpClientService) Init(a *app.App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", a)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockAnyPpClientServiceMockRecorder) Init(a any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockAnyPpClientService)(nil).Init), a)
}

// Name mocks base method.
func (m *MockAnyPpClientService) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockAnyPpClientServiceMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockAnyPpClientService)(nil).Name))
}
