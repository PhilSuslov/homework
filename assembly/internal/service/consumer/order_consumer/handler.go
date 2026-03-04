package orderconsumer

import (
	"context"

	"github.com/PhilSuslov/homework/assembly/internal/model"
	kafka "github.com/PhilSuslov/homework/platform/pkg/kafka/consumer"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	events_v1 "github.com/PhilSuslov/homework/shared/pkg/proto/events/v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

func (s *service) AssemblyHandler(ctx context.Context, msg kafka.Message) error {
	event, err := s.assemblyDecoder.Decode(msg.Value)
	if err != nil {
		logger.Error(ctx, "failed to decode AssemblyHandler", zap.Error(err))
		return err
	}

	logger.Info(ctx, "Processing message",
		zap.String("Event_uuid", event.Event_uuid),
		zap.String("Order_uuid", event.Order_uuid),
		zap.String("User_uuid", event.User_uuid),
		zap.Int64("Build_time_sec", event.Build_time_sec),
	)

	assembled := model.ShipAssembled{
		Event_uuid:     event.Event_uuid,
		Order_uuid:     event.Order_uuid,
		User_uuid:      event.User_uuid,
		Build_time_sec: event.Build_time_sec,
	}

	payload, err := proto.Marshal(&events_v1.ShipAssembled{
		EventUuid:    assembled.Event_uuid,
		OrderUuid:    assembled.Order_uuid,
		UserUuid:     assembled.User_uuid,
		BuildTimeSec: assembled.Build_time_sec,
	})

	if err != nil {
		return err
	}

	return s.assemblyProducer.Send(ctx, []byte(assembled.Event_uuid), payload)
}
