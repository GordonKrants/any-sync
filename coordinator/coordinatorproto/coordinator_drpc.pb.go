// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: v0.0.34
// source: coordinator/coordinatorproto/protos/coordinator.proto

package coordinatorproto

import (
	bytes "bytes"
	context "context"
	errors "errors"
	jsonpb "github.com/anyproto/protobuf/jsonpb"
	proto "github.com/anyproto/protobuf/proto"
	drpc "storj.io/drpc"
	drpcerr "storj.io/drpc/drpcerr"
)

type drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto struct{}

func (drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return proto.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return proto.Unmarshal(buf, msg.(proto.Message))
}

func (drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto) JSONMarshal(msg drpc.Message) ([]byte, error) {
	var buf bytes.Buffer
	err := new(jsonpb.Marshaler).Marshal(&buf, msg.(proto.Message))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto) JSONUnmarshal(buf []byte, msg drpc.Message) error {
	return jsonpb.Unmarshal(bytes.NewReader(buf), msg.(proto.Message))
}

type DRPCCoordinatorClient interface {
	DRPCConn() drpc.Conn

	SpaceSign(ctx context.Context, in *SpaceSignRequest) (*SpaceSignResponse, error)
	SpaceStatusCheck(ctx context.Context, in *SpaceStatusCheckRequest) (*SpaceStatusCheckResponse, error)
	SpaceStatusCheckMany(ctx context.Context, in *SpaceStatusCheckManyRequest) (*SpaceStatusCheckManyResponse, error)
	SpaceStatusChange(ctx context.Context, in *SpaceStatusChangeRequest) (*SpaceStatusChangeResponse, error)
	SpaceMakeShareable(ctx context.Context, in *SpaceMakeShareableRequest) (*SpaceMakeShareableResponse, error)
	SpaceMakeUnshareable(ctx context.Context, in *SpaceMakeUnshareableRequest) (*SpaceMakeUnshareableResponse, error)
	NetworkConfiguration(ctx context.Context, in *NetworkConfigurationRequest) (*NetworkConfigurationResponse, error)
	DeletionLog(ctx context.Context, in *DeletionLogRequest) (*DeletionLogResponse, error)
	SpaceDelete(ctx context.Context, in *SpaceDeleteRequest) (*SpaceDeleteResponse, error)
	AccountDelete(ctx context.Context, in *AccountDeleteRequest) (*AccountDeleteResponse, error)
	AccountRevertDeletion(ctx context.Context, in *AccountRevertDeletionRequest) (*AccountRevertDeletionResponse, error)
	AclAddRecord(ctx context.Context, in *AclAddRecordRequest) (*AclAddRecordResponse, error)
	AclGetRecords(ctx context.Context, in *AclGetRecordsRequest) (*AclGetRecordsResponse, error)
	AccountLimitsSet(ctx context.Context, in *AccountLimitsSetRequest) (*AccountLimitsSetResponse, error)
}

type drpcCoordinatorClient struct {
	cc drpc.Conn
}

func NewDRPCCoordinatorClient(cc drpc.Conn) DRPCCoordinatorClient {
	return &drpcCoordinatorClient{cc}
}

func (c *drpcCoordinatorClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcCoordinatorClient) SpaceSign(ctx context.Context, in *SpaceSignRequest) (*SpaceSignResponse, error) {
	out := new(SpaceSignResponse)
	err := c.cc.Invoke(ctx, "/coordinator.Coordinator/SpaceSign", drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcCoordinatorClient) SpaceStatusCheck(ctx context.Context, in *SpaceStatusCheckRequest) (*SpaceStatusCheckResponse, error) {
	out := new(SpaceStatusCheckResponse)
	err := c.cc.Invoke(ctx, "/coordinator.Coordinator/SpaceStatusCheck", drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcCoordinatorClient) SpaceStatusCheckMany(ctx context.Context, in *SpaceStatusCheckManyRequest) (*SpaceStatusCheckManyResponse, error) {
	out := new(SpaceStatusCheckManyResponse)
	err := c.cc.Invoke(ctx, "/coordinator.Coordinator/SpaceStatusCheckMany", drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcCoordinatorClient) SpaceStatusChange(ctx context.Context, in *SpaceStatusChangeRequest) (*SpaceStatusChangeResponse, error) {
	out := new(SpaceStatusChangeResponse)
	err := c.cc.Invoke(ctx, "/coordinator.Coordinator/SpaceStatusChange", drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcCoordinatorClient) SpaceMakeShareable(ctx context.Context, in *SpaceMakeShareableRequest) (*SpaceMakeShareableResponse, error) {
	out := new(SpaceMakeShareableResponse)
	err := c.cc.Invoke(ctx, "/coordinator.Coordinator/SpaceMakeShareable", drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcCoordinatorClient) SpaceMakeUnshareable(ctx context.Context, in *SpaceMakeUnshareableRequest) (*SpaceMakeUnshareableResponse, error) {
	out := new(SpaceMakeUnshareableResponse)
	err := c.cc.Invoke(ctx, "/coordinator.Coordinator/SpaceMakeUnshareable", drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcCoordinatorClient) NetworkConfiguration(ctx context.Context, in *NetworkConfigurationRequest) (*NetworkConfigurationResponse, error) {
	out := new(NetworkConfigurationResponse)
	err := c.cc.Invoke(ctx, "/coordinator.Coordinator/NetworkConfiguration", drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcCoordinatorClient) DeletionLog(ctx context.Context, in *DeletionLogRequest) (*DeletionLogResponse, error) {
	out := new(DeletionLogResponse)
	err := c.cc.Invoke(ctx, "/coordinator.Coordinator/DeletionLog", drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcCoordinatorClient) SpaceDelete(ctx context.Context, in *SpaceDeleteRequest) (*SpaceDeleteResponse, error) {
	out := new(SpaceDeleteResponse)
	err := c.cc.Invoke(ctx, "/coordinator.Coordinator/SpaceDelete", drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcCoordinatorClient) AccountDelete(ctx context.Context, in *AccountDeleteRequest) (*AccountDeleteResponse, error) {
	out := new(AccountDeleteResponse)
	err := c.cc.Invoke(ctx, "/coordinator.Coordinator/AccountDelete", drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcCoordinatorClient) AccountRevertDeletion(ctx context.Context, in *AccountRevertDeletionRequest) (*AccountRevertDeletionResponse, error) {
	out := new(AccountRevertDeletionResponse)
	err := c.cc.Invoke(ctx, "/coordinator.Coordinator/AccountRevertDeletion", drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcCoordinatorClient) AclAddRecord(ctx context.Context, in *AclAddRecordRequest) (*AclAddRecordResponse, error) {
	out := new(AclAddRecordResponse)
	err := c.cc.Invoke(ctx, "/coordinator.Coordinator/AclAddRecord", drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcCoordinatorClient) AclGetRecords(ctx context.Context, in *AclGetRecordsRequest) (*AclGetRecordsResponse, error) {
	out := new(AclGetRecordsResponse)
	err := c.cc.Invoke(ctx, "/coordinator.Coordinator/AclGetRecords", drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcCoordinatorClient) AccountLimitsSet(ctx context.Context, in *AccountLimitsSetRequest) (*AccountLimitsSetResponse, error) {
	out := new(AccountLimitsSetResponse)
	err := c.cc.Invoke(ctx, "/coordinator.Coordinator/AccountLimitsSet", drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCCoordinatorServer interface {
	SpaceSign(context.Context, *SpaceSignRequest) (*SpaceSignResponse, error)
	SpaceStatusCheck(context.Context, *SpaceStatusCheckRequest) (*SpaceStatusCheckResponse, error)
	SpaceStatusCheckMany(context.Context, *SpaceStatusCheckManyRequest) (*SpaceStatusCheckManyResponse, error)
	SpaceStatusChange(context.Context, *SpaceStatusChangeRequest) (*SpaceStatusChangeResponse, error)
	SpaceMakeShareable(context.Context, *SpaceMakeShareableRequest) (*SpaceMakeShareableResponse, error)
	SpaceMakeUnshareable(context.Context, *SpaceMakeUnshareableRequest) (*SpaceMakeUnshareableResponse, error)
	NetworkConfiguration(context.Context, *NetworkConfigurationRequest) (*NetworkConfigurationResponse, error)
	DeletionLog(context.Context, *DeletionLogRequest) (*DeletionLogResponse, error)
	SpaceDelete(context.Context, *SpaceDeleteRequest) (*SpaceDeleteResponse, error)
	AccountDelete(context.Context, *AccountDeleteRequest) (*AccountDeleteResponse, error)
	AccountRevertDeletion(context.Context, *AccountRevertDeletionRequest) (*AccountRevertDeletionResponse, error)
	AclAddRecord(context.Context, *AclAddRecordRequest) (*AclAddRecordResponse, error)
	AclGetRecords(context.Context, *AclGetRecordsRequest) (*AclGetRecordsResponse, error)
	AccountLimitsSet(context.Context, *AccountLimitsSetRequest) (*AccountLimitsSetResponse, error)
}

type DRPCCoordinatorUnimplementedServer struct{}

func (s *DRPCCoordinatorUnimplementedServer) SpaceSign(context.Context, *SpaceSignRequest) (*SpaceSignResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCCoordinatorUnimplementedServer) SpaceStatusCheck(context.Context, *SpaceStatusCheckRequest) (*SpaceStatusCheckResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCCoordinatorUnimplementedServer) SpaceStatusCheckMany(context.Context, *SpaceStatusCheckManyRequest) (*SpaceStatusCheckManyResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCCoordinatorUnimplementedServer) SpaceStatusChange(context.Context, *SpaceStatusChangeRequest) (*SpaceStatusChangeResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCCoordinatorUnimplementedServer) SpaceMakeShareable(context.Context, *SpaceMakeShareableRequest) (*SpaceMakeShareableResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCCoordinatorUnimplementedServer) SpaceMakeUnshareable(context.Context, *SpaceMakeUnshareableRequest) (*SpaceMakeUnshareableResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCCoordinatorUnimplementedServer) NetworkConfiguration(context.Context, *NetworkConfigurationRequest) (*NetworkConfigurationResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCCoordinatorUnimplementedServer) DeletionLog(context.Context, *DeletionLogRequest) (*DeletionLogResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCCoordinatorUnimplementedServer) SpaceDelete(context.Context, *SpaceDeleteRequest) (*SpaceDeleteResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCCoordinatorUnimplementedServer) AccountDelete(context.Context, *AccountDeleteRequest) (*AccountDeleteResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCCoordinatorUnimplementedServer) AccountRevertDeletion(context.Context, *AccountRevertDeletionRequest) (*AccountRevertDeletionResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCCoordinatorUnimplementedServer) AclAddRecord(context.Context, *AclAddRecordRequest) (*AclAddRecordResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCCoordinatorUnimplementedServer) AclGetRecords(context.Context, *AclGetRecordsRequest) (*AclGetRecordsResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCCoordinatorUnimplementedServer) AccountLimitsSet(context.Context, *AccountLimitsSetRequest) (*AccountLimitsSetResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

type DRPCCoordinatorDescription struct{}

func (DRPCCoordinatorDescription) NumMethods() int { return 14 }

func (DRPCCoordinatorDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/coordinator.Coordinator/SpaceSign", drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCCoordinatorServer).
					SpaceSign(
						ctx,
						in1.(*SpaceSignRequest),
					)
			}, DRPCCoordinatorServer.SpaceSign, true
	case 1:
		return "/coordinator.Coordinator/SpaceStatusCheck", drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCCoordinatorServer).
					SpaceStatusCheck(
						ctx,
						in1.(*SpaceStatusCheckRequest),
					)
			}, DRPCCoordinatorServer.SpaceStatusCheck, true
	case 2:
		return "/coordinator.Coordinator/SpaceStatusCheckMany", drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCCoordinatorServer).
					SpaceStatusCheckMany(
						ctx,
						in1.(*SpaceStatusCheckManyRequest),
					)
			}, DRPCCoordinatorServer.SpaceStatusCheckMany, true
	case 3:
		return "/coordinator.Coordinator/SpaceStatusChange", drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCCoordinatorServer).
					SpaceStatusChange(
						ctx,
						in1.(*SpaceStatusChangeRequest),
					)
			}, DRPCCoordinatorServer.SpaceStatusChange, true
	case 4:
		return "/coordinator.Coordinator/SpaceMakeShareable", drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCCoordinatorServer).
					SpaceMakeShareable(
						ctx,
						in1.(*SpaceMakeShareableRequest),
					)
			}, DRPCCoordinatorServer.SpaceMakeShareable, true
	case 5:
		return "/coordinator.Coordinator/SpaceMakeUnshareable", drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCCoordinatorServer).
					SpaceMakeUnshareable(
						ctx,
						in1.(*SpaceMakeUnshareableRequest),
					)
			}, DRPCCoordinatorServer.SpaceMakeUnshareable, true
	case 6:
		return "/coordinator.Coordinator/NetworkConfiguration", drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCCoordinatorServer).
					NetworkConfiguration(
						ctx,
						in1.(*NetworkConfigurationRequest),
					)
			}, DRPCCoordinatorServer.NetworkConfiguration, true
	case 7:
		return "/coordinator.Coordinator/DeletionLog", drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCCoordinatorServer).
					DeletionLog(
						ctx,
						in1.(*DeletionLogRequest),
					)
			}, DRPCCoordinatorServer.DeletionLog, true
	case 8:
		return "/coordinator.Coordinator/SpaceDelete", drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCCoordinatorServer).
					SpaceDelete(
						ctx,
						in1.(*SpaceDeleteRequest),
					)
			}, DRPCCoordinatorServer.SpaceDelete, true
	case 9:
		return "/coordinator.Coordinator/AccountDelete", drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCCoordinatorServer).
					AccountDelete(
						ctx,
						in1.(*AccountDeleteRequest),
					)
			}, DRPCCoordinatorServer.AccountDelete, true
	case 10:
		return "/coordinator.Coordinator/AccountRevertDeletion", drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCCoordinatorServer).
					AccountRevertDeletion(
						ctx,
						in1.(*AccountRevertDeletionRequest),
					)
			}, DRPCCoordinatorServer.AccountRevertDeletion, true
	case 11:
		return "/coordinator.Coordinator/AclAddRecord", drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCCoordinatorServer).
					AclAddRecord(
						ctx,
						in1.(*AclAddRecordRequest),
					)
			}, DRPCCoordinatorServer.AclAddRecord, true
	case 12:
		return "/coordinator.Coordinator/AclGetRecords", drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCCoordinatorServer).
					AclGetRecords(
						ctx,
						in1.(*AclGetRecordsRequest),
					)
			}, DRPCCoordinatorServer.AclGetRecords, true
	case 13:
		return "/coordinator.Coordinator/AccountLimitsSet", drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCCoordinatorServer).
					AccountLimitsSet(
						ctx,
						in1.(*AccountLimitsSetRequest),
					)
			}, DRPCCoordinatorServer.AccountLimitsSet, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterCoordinator(mux drpc.Mux, impl DRPCCoordinatorServer) error {
	return mux.Register(impl, DRPCCoordinatorDescription{})
}

type DRPCCoordinator_SpaceSignStream interface {
	drpc.Stream
	SendAndClose(*SpaceSignResponse) error
}

type drpcCoordinator_SpaceSignStream struct {
	drpc.Stream
}

func (x *drpcCoordinator_SpaceSignStream) SendAndClose(m *SpaceSignResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCCoordinator_SpaceStatusCheckStream interface {
	drpc.Stream
	SendAndClose(*SpaceStatusCheckResponse) error
}

type drpcCoordinator_SpaceStatusCheckStream struct {
	drpc.Stream
}

func (x *drpcCoordinator_SpaceStatusCheckStream) SendAndClose(m *SpaceStatusCheckResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCCoordinator_SpaceStatusCheckManyStream interface {
	drpc.Stream
	SendAndClose(*SpaceStatusCheckManyResponse) error
}

type drpcCoordinator_SpaceStatusCheckManyStream struct {
	drpc.Stream
}

func (x *drpcCoordinator_SpaceStatusCheckManyStream) SendAndClose(m *SpaceStatusCheckManyResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCCoordinator_SpaceStatusChangeStream interface {
	drpc.Stream
	SendAndClose(*SpaceStatusChangeResponse) error
}

type drpcCoordinator_SpaceStatusChangeStream struct {
	drpc.Stream
}

func (x *drpcCoordinator_SpaceStatusChangeStream) SendAndClose(m *SpaceStatusChangeResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCCoordinator_SpaceMakeShareableStream interface {
	drpc.Stream
	SendAndClose(*SpaceMakeShareableResponse) error
}

type drpcCoordinator_SpaceMakeShareableStream struct {
	drpc.Stream
}

func (x *drpcCoordinator_SpaceMakeShareableStream) SendAndClose(m *SpaceMakeShareableResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCCoordinator_SpaceMakeUnshareableStream interface {
	drpc.Stream
	SendAndClose(*SpaceMakeUnshareableResponse) error
}

type drpcCoordinator_SpaceMakeUnshareableStream struct {
	drpc.Stream
}

func (x *drpcCoordinator_SpaceMakeUnshareableStream) SendAndClose(m *SpaceMakeUnshareableResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCCoordinator_NetworkConfigurationStream interface {
	drpc.Stream
	SendAndClose(*NetworkConfigurationResponse) error
}

type drpcCoordinator_NetworkConfigurationStream struct {
	drpc.Stream
}

func (x *drpcCoordinator_NetworkConfigurationStream) SendAndClose(m *NetworkConfigurationResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCCoordinator_DeletionLogStream interface {
	drpc.Stream
	SendAndClose(*DeletionLogResponse) error
}

type drpcCoordinator_DeletionLogStream struct {
	drpc.Stream
}

func (x *drpcCoordinator_DeletionLogStream) SendAndClose(m *DeletionLogResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCCoordinator_SpaceDeleteStream interface {
	drpc.Stream
	SendAndClose(*SpaceDeleteResponse) error
}

type drpcCoordinator_SpaceDeleteStream struct {
	drpc.Stream
}

func (x *drpcCoordinator_SpaceDeleteStream) SendAndClose(m *SpaceDeleteResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCCoordinator_AccountDeleteStream interface {
	drpc.Stream
	SendAndClose(*AccountDeleteResponse) error
}

type drpcCoordinator_AccountDeleteStream struct {
	drpc.Stream
}

func (x *drpcCoordinator_AccountDeleteStream) SendAndClose(m *AccountDeleteResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCCoordinator_AccountRevertDeletionStream interface {
	drpc.Stream
	SendAndClose(*AccountRevertDeletionResponse) error
}

type drpcCoordinator_AccountRevertDeletionStream struct {
	drpc.Stream
}

func (x *drpcCoordinator_AccountRevertDeletionStream) SendAndClose(m *AccountRevertDeletionResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCCoordinator_AclAddRecordStream interface {
	drpc.Stream
	SendAndClose(*AclAddRecordResponse) error
}

type drpcCoordinator_AclAddRecordStream struct {
	drpc.Stream
}

func (x *drpcCoordinator_AclAddRecordStream) SendAndClose(m *AclAddRecordResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCCoordinator_AclGetRecordsStream interface {
	drpc.Stream
	SendAndClose(*AclGetRecordsResponse) error
}

type drpcCoordinator_AclGetRecordsStream struct {
	drpc.Stream
}

func (x *drpcCoordinator_AclGetRecordsStream) SendAndClose(m *AclGetRecordsResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCCoordinator_AccountLimitsSetStream interface {
	drpc.Stream
	SendAndClose(*AccountLimitsSetResponse) error
}

type drpcCoordinator_AccountLimitsSetStream struct {
	drpc.Stream
}

func (x *drpcCoordinator_AccountLimitsSetStream) SendAndClose(m *AccountLimitsSetResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_coordinator_coordinatorproto_protos_coordinator_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}
