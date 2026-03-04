package decode

import (
	"fmt"

	"github.com/PhilSuslov/homework/notification/internal/model"
	events_v1 "github.com/PhilSuslov/homework/shared/pkg/proto/events/v1"
	"google.golang.org/protobuf/proto"
)

type decodeAssembled struct{}

func NewOrderAssembledDecoder() *decodeAssembled { return &decodeAssembled{} }

func (d *decodeAssembled) Decode(data []byte) (model.ShipAssembled, error) {
	var pb events_v1.ShipAssembled

	if err := proto.Unmarshal(data, &pb); err != nil {
		return model.ShipAssembled{}, fmt.Errorf("failed to unmarshal data: %v\n pb:%v\n", data, pb)
	}

	return model.ShipAssembled{
		Event_uuid:     pb.EventUuid,
		Order_uuid:     pb.OrderUuid,
		User_uuid:      pb.UserUuid,
		Build_time_sec: pb.BuildTimeSec,
	}, nil
}
