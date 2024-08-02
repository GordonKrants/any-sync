package objectmessages

import (
	"fmt"

	"github.com/anyproto/protobuf/proto"

	"github.com/anyproto/any-sync/commonspace/spacesyncproto"
)

type BroadcastOptions struct {
	EmptyPeers []string
}

type InnerHeadUpdate interface {
	Marshall(data ObjectMeta) ([]byte, error)
	Heads() []string
	MsgSize() uint64
}

type ObjectMeta struct {
	PeerId   string
	ObjectId string
	SpaceId  string
}

type HeadUpdate struct {
	Meta   ObjectMeta
	Bytes  []byte
	Update InnerHeadUpdate
}

func (h *HeadUpdate) MsgSize() uint64 {
	var byteSize uint64
	if h.Update != nil {
		byteSize += h.Update.MsgSize()
	} else {
		byteSize += uint64(len(h.Bytes))
	}
	return byteSize + uint64(len(h.Meta.PeerId)) + uint64(len(h.Meta.ObjectId)) + uint64(len(h.Meta.SpaceId))
}

func (h *HeadUpdate) SetPeerId(peerId string) {
	h.Meta.PeerId = peerId
}

func (h *HeadUpdate) SetProtoMessage(message proto.Message) error {
	var (
		msg *spacesyncproto.ObjectSyncMessage
		ok  bool
	)
	if msg, ok = message.(*spacesyncproto.ObjectSyncMessage); !ok {
		return fmt.Errorf("unexpected message type: %T", message)
	}
	h.Bytes = msg.GetPayload()
	h.Meta.SpaceId = msg.SpaceId
	h.Meta.ObjectId = msg.ObjectId
	return nil
}

func (h *HeadUpdate) ProtoMessage() (proto.Message, error) {
	if h.Update != nil {
		payload, err := h.Update.Marshall(h.Meta)
		if err != nil {
			return nil, err
		}
		return &spacesyncproto.ObjectSyncMessage{
			SpaceId:  h.Meta.SpaceId,
			Payload:  payload,
			ObjectId: h.Meta.ObjectId,
		}, nil
	}
	return &spacesyncproto.ObjectSyncMessage{}, nil
}

func (h *HeadUpdate) PeerId() string {
	return h.Meta.PeerId
}

func (h *HeadUpdate) ObjectId() string {
	return h.Meta.ObjectId
}
