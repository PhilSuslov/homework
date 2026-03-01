package decode

import (
	"fmt"

	"github.com/PhilSuslov/homework/order/internal/model"
	events_v1 "github.com/PhilSuslov/homework/shared/pkg/proto/events/v1"
	"google.golang.org/protobuf/proto"
)

type decoder struct {}

func NewOrderRecordedDecoder() *decoder {return &decoder{}}

func (d *decoder) Decode(data[]byte) (model.OrderRecordedEvent, error) {
	var pb events_v1.OrderRecorded

	if err := proto.Unmarshal(data, &pb); err != nil{
		return model.OrderRecordedEvent{}, fmt.Errorf("failed to unmarshal data: %v\n pb:%v\n", data, pb)
	}

	return model.OrderRecordedEvent{
		Event_uuid: pb.EventUuid,
		Order_uuid: pb.OrderUuid,
		User_uuid: pb.UserUuid,
		Payment_method: pb.PaymentMethod,
		Transaction_uuid: pb.TransactionUuid,
	}, nil
}